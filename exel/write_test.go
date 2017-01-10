package exel

import (
	"testing"
)

func TestMarshalToFile(t *testing.T) {
	obj := []test{test{A: "a", C: child{"child"}}, test{A: "aa", B: 15}, test{A: "aaa", B: 7}}
	if err := MarshalToFile(obj); err != nil {
		t.Fatal(err)
	}

}
