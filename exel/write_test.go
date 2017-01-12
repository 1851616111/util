package exel

import (
	"testing"
)

func TestMarshalToFile(t *testing.T) {
	obj := []Tag{Tag{
		Map :map[string]string{"123:456":"789"},

	},
	}
	if err := MarshalToFile(obj); err != nil {
		t.Fatal(err)
	}

}
