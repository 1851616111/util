package name

import "testing"

func TestValidate(t *testing.T) {
	targets := []string{"郭靖", "毛泽东", "欧阳美声"}
	for _, v := range targets {
		if err := Validate(v); err != nil {
			t.Fatal(err)
		}
	}

	targets = []string{"郭", "毛", "欧阳美声订单", "sssssssssssssss", "1111", ""}

	for _, v := range targets {
		if err := Validate(v); err == nil {
			t.Fatal()
		}
	}
}
