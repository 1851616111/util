package main

import (
	"fmt"
	"github.com/1851616111/tantan-test/db"
	"github.com/gorilla/mux"
	"github.com/pivotal-golang/lager"
	"log"
	"net/http"
	"os"
	"path/filepath"

	relationshipapi "github.com/1851616111/tantan-test/relationship"
	userapi "github.com/1851616111/tantan-test/user"
)

var (
	IDControllerCfg    *IDController
	ServiceCfg         *Service
	userClient         userapi.UserInterface
	relationShipClient relationshipapi.RelationShipInterface
)

func init() {

	db, err := db.NewDBFromJsonFile("simple-conf.json")
	if err != nil {
		log.Fatalf("init db err %v", err)
	}
	log.Println("init database router success.")

	userClient = userapi.NewUserClient(db)
	relationShipClient = relationshipapi.NewRelationShipClient(db)

	config := loadConfig("simple-conf.json")
	IDControllerCfg = &config.IDController
	ServiceCfg = &config.Service

	log.Printf("idcontroller config %v\n", IDControllerCfg)
	log.Printf("service config %v\n", ServiceCfg)
}

func main() {
	serviceName := filepath.Base(os.Args[0])

	logger := lager.NewLogger(serviceName)
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.INFO))

	Api := New(logger)
	http.Handle("/", Api)

	logger.Info("listening on 8080.")
	fmt.Println(http.ListenAndServe(":8080", nil))

}

func New(logger lager.Logger) http.Handler {
	router := mux.NewRouter()
	AttachRoutes(router, logger)
	return router
}

func AttachRoutes(router *mux.Router, logger lager.Logger) {
	router.HandleFunc("/users", middler(logger, listUsersHandler)).Methods("GET")
	router.HandleFunc("/users", middler(logger, createUserHandler)).Methods("POST")

	router.HandleFunc("/users/{user_id}/relationships/{other_user_id}", middler(logger, createRelationShipHandler)).Methods("PUT")
	router.HandleFunc("/users/{user_id}/relationships", middler(logger, listRelationShipsHandler)).Methods("GET")
}
