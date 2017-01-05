package etcd

import "testing"

func TestNewClient(t *testing.T) {
	 NewClient("http://127.0.0.1:2379","", "")
}

func TestStorage_SetString(t *testing.T) {
	cli := NewClient("http://127.0.0.1:2379","", "")
	if err := cli.SetString("abc", "123"); err != nil {
		t.Fatal(err)
	}
}

func TestStorage_GetString(t *testing.T) {
	cli := NewClient("http://127.0.0.1:2379","", "")
	if s, err := cli.GetString("abc"); err != nil && s != "123" {
		t.Fatal(err)
	}
}

