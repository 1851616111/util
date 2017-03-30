package api_token

import (
	"encoding/json"
	"errors"
	httput "github.com/1851616111/util/http"
	"net/http"
)

func (c *controller) updateToken() error {
	rsp, err := httput.Send(c.getRequestSpec())
	if err != nil {
		return err
	} else {
		if token, err := ParseToken(rsp); err != nil {
			return err
		} else {
			c.set(token)
			return nil
		}
	}
}

func (c *controller) getRequestSpec() *httput.HttpSpec {
	c.l.RLock()
	defer c.l.RUnlock()

	spec := httput.HttpSpec{
		URL:       TokenUrl,
		Method:    "GET",
		URLParams: c.params,
	}

	return &spec
}

func (c *controller) set(tk *token) {
	c.l.Lock()
	defer c.l.Unlock()

	c.token = tk.Token
	c.expireSec = tk.Expire
}

func (c *controller) getExpire() uint16 {
	c.l.RLock()
	defer c.l.RUnlock()

	ex := c.expireSec - 60
	if ex < 0 {
		ex = 1
	}

	return ex
}

func (c *controller) GetToken() string {
	c.l.RLock()
	defer c.l.RUnlock()
	return c.token
}

func (c *controller) setExpire(sec uint16) {
	c.l.Lock()
	defer c.l.Unlock()

	c.expireSec = sec
	return
}

func ParseToken(rsp *http.Response) (*token, error) {
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
