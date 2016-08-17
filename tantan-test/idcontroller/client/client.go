package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GenUserID(host string, port uint32) (string, error) {
	if host == "localhost" {
		host = "127.0.0.1"
	}

	rsp, err := http.Get(fmt.Sprintf("http://%s:%d/internal/idcontroller/user/id", host, port))
	if err != nil {
		return "", err
	}

	defer rsp.Body.Close()

	res := map[string]string{}
	if err := json.NewDecoder(rsp.Body).Decode(&res); err != nil {
		return "", err
	}

	return res["id"], nil
}

func GenRelationShipID(host string, port uint32) (string, error) {
	if host == "localhost" {
		host = "127.0.0.1"
	}

	rsp, err := http.Get(fmt.Sprintf("http://%s:%d/internal/idcontroller/relationship/id", host, port))
	if err != nil {
		return "", err
	}

	defer rsp.Body.Close()

	res := map[string]string{}
	if err := json.NewDecoder(rsp.Body).Decode(&res); err != nil {
		return "", err
	}

	return res["id"], nil
}
