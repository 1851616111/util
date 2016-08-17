package user

import (
	"fmt"
	"github.com/1851616111/tantan-test/db"
	"testing"
)

func Test_Ceate(t *testing.T) {
	db, err := db.NewDBFromJsonFile("../conf.json")
	if err != nil {
		t.Errorf("new db err %v\n", err)
	}

	userCli := NewUserClient(db)

	if err := userCli.Create(&User{
		Id:   "10000000005",
		Name: "Michael",
	}); err != nil {
		t.Errorf("create user err %v\n", err)
	}

}

func Test_List(t *testing.T) {
	db, err := db.NewDBFromJsonFile("../conf.json")
	if err != nil {
		t.Errorf("new db err %v\n", err)
	}

	userCli := NewUserClient(db)

	_, err = userCli.List(nil)
	if err != nil {
		t.Errorf("list users err %v\n", err)
	}

}
