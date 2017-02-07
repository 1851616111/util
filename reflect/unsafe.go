package reflect

import (
	"reflect"
	"unsafe"
)


func ToString(b []byte) string {
	h := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	return *(*string)(unsafe.Pointer(&reflect.StringHeader{h.Data, h.Len}))
}

func ToBytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{sh.Data, sh.Len, sh.Len}))
}
