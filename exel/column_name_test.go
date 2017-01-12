package exel

import (
	"reflect"
	"testing"
)

type chield struct {
	Name string `xlsx:"姓名2"`
	//Class  string
	//Friend int               `xlsx:"-"`
	//Map    map[string]string `xlsx:"mmm"`
	//Array  []string
}

type Tag struct {
	A     string      `xlsx:"姓名"`
	B     int         `xlsx:"年龄"`
	C     interface{} `xlsx:"性别"`
	D     chield      `xlsx:"孩子"`
	Map    map[string]string `xlsx:"mmm"`
	Array []string
}

func Test_ColumnNames(t *testing.T) {
	tag := Tag{
		A: "abc",
		B: 10,
	}

	columnNames(reflect.TypeOf(tag))
}
