package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pivotal-golang/lager"
	"net/http"
	"os"
	"path/filepath"
)

var (
	userIDPool IDPool
)

func init() {
	userID := &safeID{idStarting: 100000000000} //10 billion

	userIDPool = NewPool(userID, nil)
}

func Run(port uint32) {
	serviceName := filepath.Base(os.Args[0])

	logger := lager.NewLogger(serviceName)
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.INFO))

	api := newHandlers(logger)
	http.Handle("/", api)
	fmt.Println(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))

}

func newHandlers(logger lager.Logger) http.Handler {
	router := mux.NewRouter()
	attachRoutes(router, logger)
	return router
}

func attachRoutes(router *mux.Router, logger lager.Logger) {
	router.HandleFunc("/internal/idcontroller/user/id", userIDHandler).Methods("GET")
}

//curl localhost:8081/internal/idcontroller/user/id
func userIDHandler(w http.ResponseWriter, r *http.Request) {
	ret := responseID{
		ID: userIDPool.GenID(),
	}

	respond(w, http.StatusOK, ret)
}

type responseID struct {
	ID string `json:"id"`
}
