package user

import (
	"encoding/json"
	"github.com/1851616111/tantan-test/db"
	"github.com/1851616111/tantan-test/router"
	"gopkg.in/pg.v4"
)

const key = "user"

type UserInterface interface {
	List(*ListOption) (*UserList, error)
	Create(*User) error
}

func NewUserClient(db *db.DB) UserInterface {
	return &impl{
		db,
	}
}

type ListOption struct {
	Per_Page int
}

type impl struct {
	*db.DB
}

// todo this list is not very efficiency
// it needs go func()
func (i *impl) List(o *ListOption) (*UserList, error) {
	i.DB.RLock()
	addrToDBToTablesMap := i.DB.TableKindToAddrToDBToTableMappings[key]
	i.DB.RUnlock()

	//todo need a static statis for every database, so i will known how much user in every table
	// and i will not query all the users for every table
	// for now, i have to make a big enough chan for not block
	// because i don't know how many users, maybe 1 maybe 10000

	resCh := make(chan []*User, 300)
	for addr, dbToTablesMap := range addrToDBToTablesMap {
		for db, tables := range dbToTablesMap {

			unionQuery := unionUsersStr(tables)
			if len(unionQuery) == 0 {
				continue
			}

			listFn := func(db *pg.DB) error {
				users := []*User{}
				res, err := db.Query(&users, unionQuery)
				if err != nil {
					return err
				}

				if res.Affected() > 0 {
					resCh <- users
				}

				return nil
			}

			i.DB.UsedDatabaseFunc(key, addr, db, listFn)
		}
	}

	l := []*User{}
	for i := 1; i <= len(resCh); i++ {
		l = append(l, (<-resCh)...)
	}

	ul := UserList(l)
	return &ul, nil
}

func (i *impl) Create(u *User) error {
	i.DB.RLock()

	//acquire route addr database table info
	routeInfoStr, err := i.DB.TableKindToRouterMappings[key].Route(u.Id)
	if err != nil {
		return err
	}

	var routeInfo router.RouteResponse
	if err := json.Unmarshal([]byte(routeInfoStr), &routeInfo); err != nil {
		return err
	}
	i.DB.RUnlock()

	createUserFn := func(db *pg.DB) error {
		if _, err := db.Exec(u.create(routeInfo.Table)); err != nil {
			return err
		}

		return nil
	}

	return i.DB.UsedDatabaseFunc(key, routeInfo.Addr, routeInfo.DB, createUserFn)

}
