package http

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func Response(w http.ResponseWriter, header int, body interface{}) {
	w.WriteHeader(header)
	fmt.Fprintf(w, "%v", body)
	return
}

func ResponseJson(w http.ResponseWriter, header int, body interface{}) error {
	w.WriteHeader(header)
	return json.NewEncoder(w).Encode(body)
}