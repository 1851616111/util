package exel

import (
	"testing"
	"github.com/1851616111/util/validator/mobile"
	"fmt"
)

func Test_ReadColumn(t *testing.T) {
	ret, filters, err := ReadColumn("22-30岁北京.xlsx", 1, mobile.Validate)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ret)
	fmt.Println(filters)
}

