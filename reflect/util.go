package reflect

import (
	"fmt"
	"github.com/lytics/logrus"
	"reflect"
)

func ValueToString(rv reflect.Value) string {
	// Display each value based on its Kind.
	switch rv.Type().Kind() {

	case reflect.String:

		return fmt.Sprintf("%s", rv.String())

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:

		return fmt.Sprintf("%v", rv.Int())

	case reflect.Float32, reflect.Float64:

		return fmt.Sprintf("%v", rv.Float())

	case reflect.Bool:

		return fmt.Sprintf("%v", rv.Bool())

	case reflect.Ptr:
		if !rv.Elem().CanAddr() {
			return ""
		}
		return ValueToString(rv.Elem())

	case reflect.Slice:

		return SliceToString(rv)

	case reflect.Struct:

		return StructToString(rv)

	case reflect.Map:

		return MapToString(rv)

	default:
		logrus.Errorf("util.reflect.ValueToString: unknow value type %s\n", rv.Type().Kind().String())
		return ""
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
		s = append(s, []byte(ValueToString(rv.Index(i)))...)
		if i < rv.Len()-1 {
			s = append(s, []byte(",")...)
		}
	}

	s = append(s, []byte("]")...)
	return string(s)
}

func StructToString(rv reflect.Value) string {
	data := []byte("{")

	for i := 0; i < rv.NumField(); i++ {
		data = append(data, []byte(fmt.Sprintf("%s:%s", rv.Type().Field(i).Name, ValueToString(rv.Field(i))))...)
		if i < rv.NumField()-1 {
			data = append(data, []byte(",")...)
		}
	}

	data = append(data, []byte("}")...)
	return string(data)
}
