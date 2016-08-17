package router

import (
	"encoding/json"
	"fmt"
)

type route struct {
	addrRouter  Router
	dbRouter    Router
	TableRouter Router
}

func (s *route) Route(key string) (string, error) {
	var addr, db, table string
	var err error

	addr, err = s.addrRouter.Route(key)
	if err != nil {
		return "", fmt.Errorf("routing id %s, route addr err %v", key, err.Error())
	}

	db, err = s.dbRouter.Route(key)
	if err != nil {
		return "", fmt.Errorf("routing id %s, route database err %v", key, err.Error())
	}

	table, err = s.TableRouter.Route(key)
	if err != nil {
		return "", fmt.Errorf("routing id %s, route table err %v", key, err.Error())
	}

	rsp := &RouteResponse{
		Addr:  addr,
		DB:    db,
		Table: table,
	}

	b, _ := json.Marshal(rsp)

	return string(b), nil
}

type RouteResponse struct {
	Addr  string `json:"addr"`
	DB    string `json:"db"`
	Table string `json:"table"`
}

type Config struct {
	Nodes      []string
	Size       uint64
	Begin      uint64
	Compatible bool
}

func NewRouter(addrConfig, dbConfig, tableConfig *Config) (Router, error) {
	var err error
	var addr, dbRouter, tableRouter Router

	if addr, err = NewBlockRouter(addrConfig.Size, addrConfig.Begin, addrConfig.Nodes, addrConfig.Compatible); err != nil {
		return nil, err
	}

	if dbRouter, err = NewHashRouter(dbConfig.Nodes); err != nil {
		return nil, err
	}

	if tableRouter, err = NewHashRouter(tableConfig.Nodes); err != nil {
		return nil, err
	}

	return &route{
		addrRouter:  addr,
		dbRouter:    dbRouter,
		TableRouter: tableRouter,
	}, nil
}
