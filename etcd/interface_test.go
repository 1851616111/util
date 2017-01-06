package etcd

import (
	"testing"
)

var cli Interface

func TestNewClient(t *testing.T) {
	cli = NewClient("http://127.0.0.1:2379", "", "")
}

func TestStorage_SetString(t *testing.T) {
	if err := cli.SetString("abc", "123"); err != nil {
		t.Fatal(err)
	}
}

func TestStorage_GetString(t *testing.T) {
	if s, err := cli.GetString("abc"); err != nil && s != "123" {
		t.Fatal(err)
	}
}

func TestStorage_Delete(t *testing.T) {
	if err := cli.DeleteKey("abc", false); err != nil {
		t.Fatal(err)
	}
}

func TestStorage_CreateObject(t *testing.T) {
	if err := cli.CreateObject("abc", map[string]string{"1": "2"}); err != nil {
		t.Fatal(err)
	}

	if err := cli.CreateObject("abc", map[string]string{"1": "2"}); !AlreadyExistErr(err) {
		t.Fatal(err)
	}
}
