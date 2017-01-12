package relationship

import (
	"encoding/json"
	"fmt"
	"github.com/1851616111/tantan-test/router"
	"gopkg.in/pg.v4"
)

type DislikesInterface interface {
	Dislike() DislikeInterface
}

type DislikeInterface interface {
	Create(from_id, to_id string) error
	Delete(from_id, to_id string) error
	List(from_id string) (*RelationShipList, error)
}

type disLikeImpl struct {
	*impl
}

func (i *disLikeImpl) Create(from_id, to_id string) error {
	i.DB.RLock()

	//acquire route addr database table info
	routeInfoStr, err := i.DB.TableKindToRouterMappings[key_dislike].Route(from_id)
	if err != nil {
		return err
	}

	var routeInfo router.RouteResponse
	if err := json.Unmarshal([]byte(routeInfoStr), &routeInfo); err != nil {
		return err
	}
	i.DB.RUnlock()

	createUserFn := func(db *pg.DB) error {
		if _, err := db.Exec(createSQL(routeInfo.Table, from_id, to_id)); err != nil {
			return err
		}

		return nil
	}

	return i.DB.UsedDatabaseFunc(key_dislike, routeInfo.Addr, routeInfo.DB, createUserFn)
}

func (i *disLikeImpl) Delete(from_id, to_id string) error {
	i.DB.RLock()

	//acquire route addr database table info
	routeInfoStr, err := i.DB.TableKindToRouterMappings[key_dislike].Route(from_id)
	if err != nil {
		return err
	}

	var routeInfo router.RouteResponse
	if err := json.Unmarshal([]byte(routeInfoStr), &routeInfo); err != nil {
		return err
	}
	i.DB.RUnlock()

	deleteDislikeFn := func(db *pg.DB) error {
		_, err := db.Exec(deleteSQL(routeInfo.Table, from_id, to_id))
		if err != nil {
			return err
		}

		return nil
	}

	return i.DB.UsedDatabaseFunc(key_dislike, routeInfo.Addr, routeInfo.DB, deleteDislikeFn)
}

func (i *disLikeImpl) List(from_id string) (*RelationShipList, error) {
	i.DB.RLock()
	addrToDBToTablesMap := i.DB.TableKindToAddrToDBToTableMappings[key_dislike]
	i.DB.RUnlock()

	//todo need a static statis for every database, so i will known how much user in every table
	// and i will not query all the users for every table
	// for now, i have to make a big enough chan for not block
	// because i don'test know how many users, maybe 1 maybe 10000

	resCh := make(chan []*relation, 300)
	for addr, dbToTablesMap := range addrToDBToTablesMap {
		for db, tables := range dbToTablesMap {

			listUnionQuery := unionRelationShipStr(from_id, tables)
			if len(listUnionQuery) == 0 {
				continue
			}

			listFn := func(db *pg.DB) error {
				ids := []*relation{}
				res, err := db.Query(&ids, listUnionQuery)
				if err != nil {
					return err
				}

				if res.Affected() > 0 {
					resCh <- ids
				}

				return nil
			}

			i.DB.UsedDatabaseFunc(key_dislike, addr, db, listFn)
		}
	}

	l := []*RelationShip{}
	for i := 1; i <= len(resCh); i++ {
		for _, id := range <-resCh {
			l = append(l, &RelationShip{
				State:  "disliked",
				UserID: id.To_id,
			})
		}

	}
	fmt.Println(len(l))

	rsl := RelationShipList(l)
	return &rsl, nil
}
