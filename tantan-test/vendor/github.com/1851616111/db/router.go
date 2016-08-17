package db

import (
	"errors"
	"fmt"
	"github.com/1851616111/tantan-test/router"
	"log"
)

var (
	userHostRouterConfig = &router.Config{
		Begin: 10000000000, // user id start with 1000000000
		Size:  50000000,    // every host contains at most of 50 million user id
	}
	relationShipHostRouterConfig = &router.Config{
		Begin: 10000000000, // relationship id start with 1000000000
		Size:  50000000,    // every host contains at most of 50 million relationship id
	}
)

func newRouter(label string, dbs []dataBase) router.Router {
	var hostCfg, instanceCfg, tableCfg *router.Config
	var err error

	if hostCfg, err = parseHostConfig(dbs, label); err != nil {
		log.Fatalf("config %s databses host err %v", label, err)
		return nil
	}
	if instanceCfg, err = parseDBInstanceConfig(dbs, label); err != nil {
		log.Fatalf("config %s databses instances err %v", label, err)
		return nil
	}
	if tableCfg, err = parseTableConfig(dbs, label); err != nil {
		log.Fatalf("config %s databses table err %v", label, err)
		return nil
	}

	var r router.Router
	if r, err = router.NewRouter(hostCfg, instanceCfg, tableCfg); err != nil {
		log.Fatalf("new %s database router %v", label, err)
		return nil
	}

	return r
}

func parseHostConfig(dbs []dataBase, label string) (*router.Config, error) {
	reduceMap := map[string]struct{}{}

	rangeDataBaseFunc(dbs, func(db dataBase) {
		if db.Label != label {
			return
		}

		if db.Host == "" || db.Port == 0 {
			return
		}

		reduceMap[fmt.Sprintf("%s:%d", db.Host, db.Port)] = struct{}{}
	})

	if len(reduceMap) == 0 {
		return nil, errors.New("parse databases host config nil.")
	}

	c := &router.Config{
		Compatible: true,
		Nodes:      keys(reduceMap),
	}

	switch label {
	case "user":
		c.Begin = userHostRouterConfig.Begin
		c.Size = userHostRouterConfig.Begin
	default:
		c.Begin = relationShipHostRouterConfig.Begin
		c.Size = relationShipHostRouterConfig.Size
	}

	return c, nil
}

func parseDBInstanceConfig(dbs []dataBase, label string) (*router.Config, error) {
	reduceMap := map[string]struct{}{}

	rangeDataBaseFunc(dbs, func(db dataBase) {
		if db.Label != label {
			return
		}

		for _, instance := range db.Instances {
			reduceMap[instance.Name] = struct{}{}
		}
	})

	if len(reduceMap) == 0 {
		return nil, errors.New("parse databases instances config nil.")
	}

	return &router.Config{
		Nodes: keys(reduceMap),
	}, nil
}

func parseTableConfig(dbs []dataBase, label string) (*router.Config, error) {
	reduceMap := map[string]struct{}{}

	rangeDataBaseFunc(dbs, func(db dataBase) {
		if db.Label != label {
			return
		}

		for _, instance := range db.Instances {
			for _, table := range instance.Tables {
				reduceMap[table] = struct{}{}
			}
		}
	})

	if len(reduceMap) == 0 {
		return nil, errors.New("parse databases tabel config nil.")
	}

	return &router.Config{
		Nodes: keys(reduceMap),
	}, nil
}

func rangeDataBaseFunc(dbs []dataBase, fn func(dataBase)) {
	if len(dbs) == 0 {
		return
	}

	for _, db := range dbs {
		fn(db)
	}

	return
}

func keys(m map[string]struct{}) []string {
	l := []string{}

	for k := range m {
		l = append(l, k)
	}

	return l
}

//
//func routeUserId(id string)(*router.RouteResponse, error){
//	str, err := UserIDRouter.Route(id)
//	if err != nil {
//		return nil, err
//	}
//
//	var rsp router.RouteResponse
//	if err := json.Unmarshal([]byte(str), &rsp); err != nil {
//		return nil, err
//	}
//
//	return &rsp,nil
//}
//
//func routeRelationShipId(id string)(*router.RouteResponse, error){
//	str, err := RelationShipIDRouter.Route(id)
//	if err != nil {
//		return nil, err
//	}
//
//	var rsp router.RouteResponse
//	if err := json.Unmarshal([]byte(str), &rsp); err != nil {
//		return nil, err
//	}
//
//	return &rsp,nil
//}
