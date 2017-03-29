package api

import (
	"github.com/1851616111/util/weichat/menu"
	"github.com/1851616111/util/http"
)

const NewMenuURL = "https://api.weixin.qq.com/cgi-bin/menu/create"

func NewMenuReqSpec(bt * menu.Button, access_token string)*http.HttpSpec{
	return &http.HttpSpec{
		URL:  NewMenuURL,
		Method: "POST",
		ContentType: http.ContentType_JSON,
		URLParams: http.NewParams().Add("access_token", access_token),
		BodyParams:  http.NewBody().Add("button", []*menu.Button{bt}),
	}
}