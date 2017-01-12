package strings

import (
	"testing"
	"reflect"
)

func Test_SubString(t *testing.T) {
	str := "Abc"
	if SubString(str, 0, 1) != "A" {
		t.Fatal(str)
	}
}


func Test_InterceptNumber(t *testing.T) {
	str := "123a456d7a"
	if ints := InterceptNumber(str); !reflect.DeepEqual(ints, []int{123,456,7})  {
		t.Fatal(ints)
	}
}