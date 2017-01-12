package exel

import (
	//reflectutil "github.com/1851616111/util/reflect"
	strutil "github.com/1851616111/util/strings"
	"reflect"
	"strings"
	//"fmt"
)

const _TAG_KEY = "xlsx"
const _TAG_FORBIDDEN = "-"
const _TAG_NOT_FOUND = ""

const UpperChar = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func columnNames(tp reflect.Type) []string {
	switch tp.Kind() {

	case reflect.Ptr:

		return columnNames(tp.Elem())

	case reflect.Map, reflect.Slice:

		return directColumnNames(tp)

	case reflect.Struct:

		return structColumnNames(tp)
	}
	return nil
}

func structColumnNames(tp reflect.Type) []string {

	names := []string{}
	for fldI := 0; fldI < tp.NumField(); fldI++ {

		//先取出名字， 若失败则抛弃（首字母小写，tag:"-"）
		ok, name := fieldExport(tp.Field(fldI))
		if !ok {
			continue
		}

		//不论是简单还是复杂架构，都要导出整体子段
		names = append(names, name)
		names = append(names, columnNames(tp.Field(fldI).Type)...)

		////interface{} is simple field for now, because when running,
		////we don't know the turely struct inner interface{}
		//if reflectutil.IsSimpleType(tp.Field(fldI).Type) || tp.Field(fldI).Type.Kind() == reflect.Interface {
		//}
	}

	return names
}

func directColumnNames(tp reflect.Type) []string {
	return []string{tp.Name()}
}

func fieldExport(field reflect.StructField) (bool, string) {
	//先判断第一个字母是否为大写，若小写则退出
	if !strings.Contains(UpperChar, strutil.SubString(field.Name, 0, 1)) {
		return false, ""
	}

	switch tag := field.Tag.Get(_TAG_KEY); tag {
	case _TAG_FORBIDDEN:
		return false, ""

	case _TAG_NOT_FOUND:
		return true, field.Name

	default:
		return true, tag
	}
}
