package exel

import (
	"github.com/1851616111/util/validator/mobile"
	"testing"
)

func Test_ReadColumn(t *testing.T) {
	_, _, err := ReadColumn("22-30岁北京.xlsx", 1, mobile.Validate)
	if err != nil {
		t.Fatal(err)
	}
}
