package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Conversion interface {
	Convert() interface{}
}

type responseError struct {
	error
}

func (p *responseError) Convert() interface{} {
	return &struct {
		Message string `json:"message"`
		Type    string `json:"type"`
	}{
		p.error.Error(),
		"error",
	}
}

func respond(w http.ResponseWriter, status int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	encoder := json.NewEncoder(w)

	if isResponseErr(response) {
		c := &responseError{
			response.(error),
		}

		err := encoder.Encode(c.Convert())
		if err != nil {
			fmt.Printf("response being attempted %d %#v\n", status, response)
			fmt.Println(err)
		}

		return
	}

	if isResponseConversion(response) {
		err := encoder.Encode(response.(Conversion).Convert())
		if err != nil {
			fmt.Printf("response being attempted %d %#v\n", status, response)
			fmt.Println(err)
		}

		return
	}

	err := encoder.Encode(response)
	if err != nil {
		fmt.Printf("response being attempted %d %#v\n", status, response)
		fmt.Println(err)
	}

	return
}

func isResponseErr(response interface{}) bool {
	_, ok := response.(error)
	return ok
}
func isResponseConversion(response interface{}) bool {
	_, ok := response.(Conversion)
	return ok
}
