package main

import (
	//"errors"
	"github.com/gorilla/mux"
	"github.com/pivotal-golang/lager"
	"net/http"
)

const (
	StatusNotAcceptContentType = 207
)

func middler(log lager.Logger, handler func(http.ResponseWriter, *http.Request, lager.Logger, map[string]string)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//if !validateContentType(w, r) {
		//	respond(w, StatusNotAcceptContentType, errors.New("only Content-Type=application/json is accepted"))
		//	return
		//}
		handler(w, r, log, mux.Vars(r))
	}
}

func validateContentType(w http.ResponseWriter, r *http.Request) bool {
	return r.Header.Get("Content-Type") == "application/json"
}
