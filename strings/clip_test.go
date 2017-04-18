package strings

import (
	"reflect"
	"testing"
	"fmt"
)

func TestClip(t *testing.T) {
	target := `"{2,1,3}"`
	if ret := Clip(&target, "\"{", ",", "}\""); !reflect.DeepEqual(ret, []string{"2", "1", "3"}) {
		t.Fatal(ret)
	}
	target = `2,1,3`
	if ret := Clip(&target, ``, `,`, ``); !reflect.DeepEqual(ret, []string{"2", "1", "3"}) {
		t.Fatal(ret)
	}

	target = `,,`
	if ret := Clip(&target, ``, `,`, ``); len(ret) != 3 {
		t.Fatal(ret)
	}

	if "" != `` {
		t.Fatal()
	}
}

func TestClip2(t *testing.T) {
	target := ""
	sl := ClipDBArray(&target)
	if len(sl) != 0 {
		t.Fatal()
	}
}