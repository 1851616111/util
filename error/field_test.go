package error

import (
	"testing"
	"reflect"
)
func TestFieldEmptyError(t *testing.T) {

	var err error

	s := struct {
		a string
	}{a: "20"}

	if err = FieldEmptyError(s.a); err != nil {
		t.Fatal(err)
	}

	s = struct {
		a string
	}{}

	if err = FieldEmptyError(s.a); err == nil {
		t.Fatal(err)
	}

	if err != ErrFieldEmpty(reflect.TypeOf(s.a).Field()) {
		t.Fatal(err)
	}

}