package mobile

import "testing"

func TestValidate(t *testing.T) {
	validMobiles := []string{"13800000000", "17744524444", "15668688888"}
	invalidMobiles := []string{"", "138000000000", "1380000000a", "aaaaaaaaaaa"}

	for _, mobile := range validMobiles {
		if err := Validate(mobile); err != nil {
			t.Fatal(err)
		}
	}

	for _, mobile := range invalidMobiles {
		if err := Validate(mobile); err == nil {
			t.Fatal(err)
		}
	}
}
