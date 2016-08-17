package db

import (
	"testing"
)

func Test_NewDBFromJsonFile(t *testing.T) {
	_, err := NewDBFromJsonFile("../conf.json")
	if err != nil {
		t.Errorf("new db err %v\n", err)
	}

}
