package exel

import (
	"reflect"
	"testing"
)

func Test_ValueToSlice(t *testing.T) {
	obj := Tag{
		A: "abc",
		B: 10,
	}

	ValueToSlice(reflect.ValueOf(obj))

}
