package etcd

import (
	"fmt"
	"github.com/coreos/go-etcd/etcd"
	"testing"
)

var cli Interface

func TestNewClient(t *testing.T) {
	cli = NewClient("http://127.0.0.1:2379", "", "")
}

func TestStorage_WatchKey(t *testing.T) {
	go func() {
		ch := make(chan *etcd.Response, 10)
		stop := make(chan bool, 1)

		go receiver(ch, stop)
		_, err := cli.WatchKey("/test/abc", 0, true, ch, stop)
		if err != etcd.ErrWatchStoppedByUser {
			t.Fatalf("Watch returned a non-user stop error")
		}

	}()
}

func TestStorage_SetString(t *testing.T) {
	if err := cli.SetString("/test/abc", "123"); err != nil {
		t.Fatal(err)
	}
}

func TestStorage_GetString(t *testing.T) {
	if s, err := cli.GetString("/test/abc"); err != nil && s != "123" {
		t.Fatal(err)
	}
}

func TestStorage_Delete(t *testing.T) {
	if err := cli.DeleteKey("/test/abc", false); err != nil {
		t.Fatal(err)
	}
}

func TestStorage_CreateObject(t *testing.T) {
	if err := cli.CreateObject("/test/abc", map[string]string{"1": "2"}); err != nil {
		t.Fatal(err)
	}

	if err := cli.CreateObject("/test/abc", map[string]string{"1": "2"}); !AlreadyExistErr(err) {
		t.Fatal(err)
	}
}

func TestStorage_GetDir(t *testing.T) {
	if rsp, err := cli.GetDir("/test"); err != nil {
		t.Fatal(err)
	}
}
