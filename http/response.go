package http

import (
	"fmt"
	"net/http"
)

func Response(w http.ResponseWriter, header int, body interface{}) {
	w.WriteHeader(header)
	fmt.Fprintf(w, "%v", body)
	return
}
