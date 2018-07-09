package handler

import (
	"crypto/sha1"
	"fmt"
	"io"
	"net/http"
	"sort"
)

//validate qualification of weichat developer
func ImportWxHandler(token string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		s := sign(r.FormValue("nonce"), r.FormValue("timestamp"), token)

		if s != r.FormValue("signature") {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(200)
			w.Write([]byte(r.FormValue("echostr")))
		}
	}
}

// sort by value
func sign(nonce, timestamp, token string) string {
	ps := []string{nonce, timestamp, token}
	sort.Strings(ps)

	return sha(ps[0] + ps[1] + ps[2])
}

func sha(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}
