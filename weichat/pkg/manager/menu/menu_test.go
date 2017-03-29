package menu

import (
	"testing"
	"github.com/1851616111/util/http"
	"io/ioutil"
)
func TestNewMenuReqSpec(t *testing.T) {
	button := NewTopButton("体检向导").AddSub(NewViewButton("预约体检", "https://open.weixin.qq.com/connect/oauth2/authorize?appid=wxd09c7682905819e6&redirect_uri=http%3a%2f%2fwww.elepick.com%2fapi&response_type=code&scope=snsapi_base&state=123#wechat_redirect")).AddSub(NewViewButton("报告查询", "http://www.elepick.com/api"))
	rsp, err := http.Send(NewMenuReqSpec(button, "1ycE5slXSsVKxubaftm1gflLYF0Mrk21-fpsU0G-igMCqOj-Nk9OBW8tsx0bYXcVkGfSioyzkr1cH_7ja0Sh-irCghRLJmxeKDYzfByv3ctmF2fI3r_cmTmME-7Lt9l0YWFfACANLU"))
	if err != nil {
		t.Fatal(err)
	}

	_, err = ioutil.ReadAll(rsp.Body)
	if err != nil {
		t.Fatal(err)
	}
}
