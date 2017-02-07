package reflect

import (
	"testing"
	"reflect"
)


func TestToString(t *testing.T) {
	b := []byte{0,0,0}
	if str := ToString(b); str != "\x00\x00\x00" {
		t.Fatal()
	}
}

func TestToBytes(t *testing.T) {
	if !reflect.DeepEqual(ToBytes("123"), []byte{byte('1'), byte('2'), byte('3')}) {
		t.Fatal()
	}
}