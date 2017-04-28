package error

import (
	"reflect"
	"fmt"
)

func ErrFieldEmpty(field string) error {
	return fmt.Errorf("object field %s not found", field)
}

func FieldEmptyError(field interface{}) error {
	rv := reflect.ValueOf(field)
	switch rv.Kind() {
	case reflect.String:
		if rv.Len() == 0 {
			return ErrFieldEmpty(rv.Type().Name())
		}
	}
	
	//func ValueToSlice(rv reflect.Value) []string {
	//// Display each value based on its Kind.
	//switch rv.Type().Kind() {
	//
	//case reflect.String:
	//
	//	return []string{rv.String()}
	//
	//case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
	//	reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
	//
	//	return []string{fmt.Sprintf("%v", rv.Int())}
	//
	//case reflect.Float32, reflect.Float64:
	//
	//	return []string{fmt.Sprintf("%v", rv.Float())}
	//
	//case reflect.Bool:
	//
	//	return []string{fmt.Sprintf("%v", rv.Bool())}
	//
	//case reflect.Ptr:
	//	if !rv.Elem().CanAddr() {
	//		return []string{}
	//	}
	//	return ValueToSlice(rv.Elem())
	//
	//case reflect.Slice:
	//
	//	return []string{SliceToString(rv)}
	//
	//case reflect.Struct:
	//
	//	return append([]string{reflectutil.StructToString(rv)}, StructToSlice(rv)...)
	//
	//case reflect.Map:
	//
	//	return []string{MapToString(rv)}
	//
	//case reflect.Interface:
	//	if rv.IsNil() {
	//		return []string{""}
	//	} else {
	//		return []string{reflectutil.ValueToString(rv.Elem())}
	//	}
	//
	//default:
	//	logrus.Errorf("util.reflect.ValueToString: unknow value type %s\n", rv.Type().Kind().String())
	//	return []string{}
	//}

	return nil
}