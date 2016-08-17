package main

import (
	"encoding/json"
	"os"
)

func loadConfig(path string) *config {
	c := new(config)
	readJsonFile(path, c)
	return c
}

type config struct {
	IDController IDController `json:"idcontroller"`
	Service      Service      `json:"service"`
}

type IDController struct {
	Host string `json:"host"`
	Port uint32 `json:"port"`
}

type Service struct {
	Port uint32 `json:"port"`
}

func readJsonFile(path string, c interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(c)

}
