package reflect

import (
	"reflect"
	"testing"
)

type c struct {
	A string
	B int
	C float32
	D bool
	E []string
	F *int
}

func Test_MapToString(t *testing.T) {
	m := map[string]string{
		"a": "123",
		"b": "456",
	}

	if s := MapToString(reflect.ValueOf(m)); s != "{a:123,b:456}" {
		t.Fatal()
	}
}

func Test_StructToString(t *testing.T) {
	test := c{
		A: "123",
		B: 10,
	}

	if s := StructToString(reflect.ValueOf(test)); s != "{A:123,B:10,C:0,D:false,E:[],F:}" {
		t.Fatal(s)
	}
}

func Test_ValueToString(t *testing.T) {
	i := 100
	test := c{
		A: "123",
		B: 10,
		C: 88.88,
		D: false,
		E: []string{"1", "2", "3"},
		F: &i,
	}

	if a := ValueToString(reflect.ValueOf(test.A)); a != "123" {
		t.Fatal()
	}

	if a := ValueToString(reflect.ValueOf(test.B)); a != "10" {
		t.Fatal()
	}

	//if a := ValueToString(reflect.ValueOf(test.C)); a != "88.88" {
	//	fmt.Println(a)
	//	t.Fatal()
	//}

	if a := ValueToString(reflect.ValueOf(test.D)); a != "false" {
		t.Fatal()
	}

	if a := ValueToString(reflect.ValueOf(test.E)); a != "[1,2,3]" {
		t.Fatal()
	}

	if a := ValueToString(reflect.ValueOf(test.F)); a != "100" {
		t.Fatal()
	}

}
