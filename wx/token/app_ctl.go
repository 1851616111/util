package token

import (
	"sync"
	"time"

	"encoding/json"
	"errors"
	httput "github.com/1851616111/util/http"
	"github.com/sirupsen/logrus"
	"net/http"
)

var (
	DefaultGrantType string = "client_credential"
	WXServerAddr     string = "https://api.weixin.qq.com/cgi-bin/token"
	log                     = logrus.New().WithFields(logrus.Fields{"pkg": "wx.token"})
)

type Controller struct {
	l sync.RWMutex //protect token and config

	expireSec uint16
	params    *httput.Params //appid,secret and grant_type

	token string
	err   error //unknown error
}

type token struct {
	Token  string `json:"access_token,omitempty"`
	Expire uint16 `json:"expires_in,omitempty"`
	Code   int    `json:"errcode,omitempty"`
	Msg    string `json:"errmsg,omitempty"`
}

//curl -G "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=wxd09c7682905819e6&secret=b9938ddfec045280eba89fab597a0c41"
func NewController(appID string, secret string) *Controller {
	return &Controller{
		l:      sync.RWMutex{},
		params: httput.NewParams().Add("grant_type", "client_credential").Add("appid", appID).Add("secret", secret),
	}
}

func (c *Controller) Run() error {
	if err := c.updateToken(); err != nil {
		return err
	}

	log.Infof("first update token(%s) success", c.token)
	go func() {
		for {
			//提前60秒进行更新
			time.Sleep(time.Second * time.Duration(c.getExpire()))
			//time.Sleep(time.Second*time.Duration(20 + 60))

			if err := c.updateToken(); err != nil {
				log.Errorf("update token err %v\n", err)

				c.setExpire(20 + 60) //actual 20 second
				c.err = err
			} else {
				log.Infof("update token(%s) success", c.token)
			}
		}
	}()

	return nil
}

func (c *Controller) updateToken() error {
	rsp, err := httput.Send(c.getRequestSpec())
	if err != nil {
		return err
	} else {
		if token, err := parseToken(rsp); err != nil {
			return err
		} else {
			c.set(token)
			return nil
		}
	}
}

func (c *Controller) getRequestSpec() *httput.HttpSpec {
	c.l.RLock()
	defer c.l.RUnlock()

	spec := httput.HttpSpec{
		URL:       WXServerAddr,
		Method:    "GET",
		URLParams: c.params,
	}

	return &spec
}

func (c *Controller) set(tk *token) {
	c.l.Lock()
	defer c.l.Unlock()

	c.token = tk.Token
	c.expireSec = tk.Expire
}

func (c *Controller) getExpire() uint16 {
	c.l.RLock()
	defer c.l.RUnlock()

	ex := c.expireSec - 60
	if ex < 0 {
		ex = 1
	}

	return ex
}

func (c *Controller) GetToken() string {
	c.l.RLock()
	defer c.l.RUnlock()
	return c.token
}

func (c *Controller) setExpire(sec uint16) {
	c.l.Lock()
	defer c.l.Unlock()

	c.expireSec = sec
	return
}

func parseToken(rsp *http.Response) (*token, error) {
	tk := &token{}
	if err := json.NewDecoder(rsp.Body).Decode(tk); err != nil {
		return nil, err
	}

	if err, ok := CodeErrMapping[tk.Code]; ok {
		return tk, err
	} else {
		return nil, errors.New(tk.Msg)
	}
}
