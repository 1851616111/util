package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func respond(w http.ResponseWriter, status int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	encoder := json.NewEncoder(w)

	err := encoder.Encode(response)
	if err != nil {
		fmt.Printf("response being attempted %d %#v\n", status, response)
		fmt.Println(err)
	}

	return
}
