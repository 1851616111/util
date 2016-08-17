package db

import (
	"encoding/json"
	"os"
)

func loadConfig(path string) (*config, error) {
	c := new(config)
	if err := readJsonFile(path, c); err != nil {
		return nil, err
	}

	return c, nil
}

type config struct {
	Databases []dataBase `json:"databases"`
}

type dataBase struct {
	Label       string `json:"label"`
	Host        string `json:"host"`
	Port        uint32 `json:"port"`
	MaxPool     int    `json:"maxPoolSize"`
	Credentials struct {
		User     string `json:"user"`
		Password string `json:"password"`
	} `json:credentials`
	Instances []struct {
		Name   string   `json:"name"`
		Tables []string `json:"tables"`
	} `json:"instances"`
}

func readJsonFile(path string, c interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(c)
}
