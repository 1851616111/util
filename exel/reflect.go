package exel

import (
	"fmt"
	reflectutil "github.com/1851616111/util/reflect"
	"github.com/lytics/logrus"
	"reflect"
)

func ValueToSlice(rv reflect.Value) []string {
	// Display each value based on its Kind.
	switch rv.Type().Kind() {

	case reflect.String:

		return []string{rv.String()}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:

		return []string{fmt.Sprintf("%v", rv.Int())}

	case reflect.Float32, reflect.Float64:

		return []string{fmt.Sprintf("%v", rv.Float())}

	case reflect.Bool:

		return []string{fmt.Sprintf("%v", rv.Bool())}

	case reflect.Ptr:
		if !rv.Elem().CanAddr() {
			return []string{}
		}
		return ValueToSlice(rv.Elem())

	case reflect.Slice:

		return []string{SliceToString(rv)}

	case reflect.Struct:

		return append([]string{reflectutil.StructToString(rv)}, StructToSlice(rv)...)

	case reflect.Map:

		return []string{MapToString(rv)}

	case reflect.Interface:
		if rv.IsNil() {
			return []string{""}
		} else {
			return []string{reflectutil.ValueToString(rv.Elem())}
		}

	default:
		logrus.Errorf("util.reflect.ValueToString: unknow value type %s\n", rv.Type().Kind().String())
		return []string{}
	}
}

func MapToString(rv reflect.Value) string {
	s := "{"

	for i, key := range rv.MapKeys() {
		s += fmt.Sprintf("%v:%v", key, rv.MapIndex(key))
		if i < rv.Len()-1 {
			s += ","
		}
	}

	s += "}"
	return s
}

func SliceToString(rv reflect.Value) string {
	s := []byte("[")

	for i := 0; i < rv.Len(); i++ {
		s = append(s, []byte(reflectutil.ValueToString(rv.Index(i)))...)
		if i < rv.Len()-1 {
			s = append(s, []byte(",")...)
		}
	}

	s = append(s, []byte("]")...)
	return string(s)
}

func StructToSlice(rv reflect.Value) []string {
	slice := []string{}
	for i := 0; i < rv.NumField(); i++ {
		export, _ := fieldExport(rv.Type().Field(i))
		if !export {
			continue
		}

		slice = append(slice, ValueToSlice(rv.Field(i))...)
	}

	return slice
}

func IsSimpleType(rt reflect.Type) bool {
	switch rt.Kind() {
	case reflect.Bool, reflect.String,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Uintptr, reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		return true

	default:
		return false
	}
}
