package handler

import (
	"net/http"

	"github.com/1851616111/util/weichat/util/sign"
	token "github.com/1851616111/util/weichat/util/user-token"
	"github.com/julienschmidt/httprouter"
)

var APP_ID string
var Token *token.Config

//validate qualification of weichat developer
func DeveloperValidater(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	s := sign.Sign(r.FormValue("nonce"), r.FormValue("timestamp"), APP_ID)

	if s != r.FormValue("signature") {
		w.WriteHeader(400)
	} else {
		w.WriteHeader(200)
		w.Write([]byte(r.FormValue("echostr")))
	}

	return
}

func SourceValidater(tokenCallBack func(*token.Token), handler func(http.ResponseWriter, *http.Request, httprouter.Params)) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		r.ParseForm()
		code := r.FormValue("code")
		if code == "" {
			forbiddenHandler(w, r, ps)
			return
		}

		tk, err := Token.Exchange(code)
		if err != nil {
			forbiddenHandler(w, r, ps)
			return
		}

		newPs := ([]httprouter.Param)(ps)
		newPs = append(newPs, httprouter.Param{"openid", tk.Open_ID})
		newPs = append(newPs, httprouter.Param{"access_token", tk.Access_Token})

		tokenCallBack(tk)

		handler(w, r, httprouter.Params(newPs))
		return
	}
}

func forbiddenHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte(`<html><head><title>迪安</title></head><body>请在微信客户端打开链接</body></html>`))
	return
}
