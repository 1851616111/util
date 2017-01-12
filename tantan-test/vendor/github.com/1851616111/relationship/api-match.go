package relationship

import (
	"encoding/json"
	"github.com/1851616111/tantan-test/router"
	"gopkg.in/pg.v4"
)

type MatchesInterface interface {
	Match() MatchInterface
}

type MatchInterface interface {
	Create(from_id, to_id string) error
	IfMatch(from_id, to_id string) (bool, error)
	Delete(from_id, to_id string) error
	List(from_id string) (*RelationShipList, error)
}

type matchImpl struct {
	*impl
}

func (i *matchImpl) Create(from_id, to_id string) error {
	i.DB.RLock()

	//acquire route addr database table info
	routeInfoStr, err := i.DB.TableKindToRouterMappings[key_match].Route(from_id)
	if err != nil {
		return err
	}

	var routeInfo router.RouteResponse
	if err := json.Unmarshal([]byte(routeInfoStr), &routeInfo); err != nil {
		return err
	}
	i.DB.RUnlock()

	createMatchFn := func(db *pg.DB) error {
		if _, err := db.Exec(createSQL(routeInfo.Table, from_id, to_id)); err != nil {
			return err
		}

		return nil
	}

	return i.DB.UsedDatabaseFunc(key_match, routeInfo.Addr, routeInfo.DB, createMatchFn)
}

func (i *matchImpl) IfMatch(from_id, to_id string) (bool, error) {
	i.DB.RLock()

	//acquire route addr database table info
	routeInfoStr, err := i.DB.TableKindToRouterMappings[key_match].Route(from_id)
	if err != nil {
		return false, err
	}

	var routeInfo router.RouteResponse
	if err := json.Unmarshal([]byte(routeInfoStr), &routeInfo); err != nil {
		return false, err
	}
	i.DB.RUnlock()

	var ifLike bool
	getMatchFn := func(db *pg.DB) error {
		res, err := db.Exec(selectSQL(routeInfo.Table, from_id, to_id))
		if err != nil {
			return err
		}

		if res.Affected() > 0 {
			ifLike = true
		}

		return nil
	}

	err = i.DB.UsedDatabaseFunc(key_match, routeInfo.Addr, routeInfo.DB, getMatchFn)
	return ifLike, err
}

func (i *matchImpl) Delete(from_id, to_id string) error {
	i.DB.RLock()

	//acquire route addr database table info
	routeInfoStr, err := i.DB.TableKindToRouterMappings[key_match].Route(from_id)
	if err != nil {
		return err
	}

	var routeInfo router.RouteResponse
	if err := json.Unmarshal([]byte(routeInfoStr), &routeInfo); err != nil {
		return err
	}
	i.DB.RUnlock()

	deleteLikeFn := func(db *pg.DB) error {
		_, err := db.Exec(deleteSQL(routeInfo.Table, from_id, to_id))
		if err != nil {
			return err
		}

		return nil
	}

	return i.DB.UsedDatabaseFunc(key_match, routeInfo.Addr, routeInfo.DB, deleteLikeFn)
}

func (i *matchImpl) List(from_id string) (*RelationShipList, error) {
	i.DB.RLock()
	addrToDBToTablesMap := i.DB.TableKindToAddrToDBToTableMappings[key_match]
	i.DB.RUnlock()

	//todo need a static statis for every database, so i will known how much user in every table
	// and i will not query all the users for every table
	// for now, i have to make a big enough chan for not block
	// because i don'test know how many users, maybe 1 maybe 10000

	resCh := make(chan []*relation, 300)
	for addr, dbToTablesMap := range addrToDBToTablesMap {
		for db, tables := range dbToTablesMap {

			unionQuery := unionRelationShipStr(from_id, tables)
			if len(unionQuery) == 0 {
				continue
			}

			listFn := func(db *pg.DB) error {
				ids := []*relation{}
				res, err := db.Query(&ids, unionQuery)
				if err != nil {
					return err
				}

				if res.Affected() > 0 {
					resCh <- ids
				}

				return nil
			}

			i.DB.UsedDatabaseFunc(key_match, addr, db, listFn)
		}
	}

	l := []*RelationShip{}
	for i := 1; i <= len(resCh); i++ {
		for _, id := range <-resCh {
			l = append(l, &RelationShip{
				State:  "matched",
				UserID: id.To_id,
			})
		}

	}

	rsl := RelationShipList(l)
	return &rsl, nil
}
