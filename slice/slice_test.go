package util

import (
	"testing"
)

func TestIndexAddString(t *testing.T) {
	s := []string{}
	IndexAddString(&s, 0, "123")
	if s[0] != "123" {
		t.Error("index add string fail \n")
	}
}

func TestIndexRemoveString(t *testing.T) {
	s := []string{"123"}
	IndexRemoveString(&s, 0)
	if len(s) != 0 {
		t.Error("remove element err")
	}
}
