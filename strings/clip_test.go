package strings

import (
	"testing"
	"reflect"
)

func TestClip(t *testing.T) {
	target := `"{2,1,3}"`
	if ret := Clip(target, ",", "\"{", "}\""); !reflect.DeepEqual(ret, []string {"2","1","3"}) {
		t.Error(ret)
	}
}
