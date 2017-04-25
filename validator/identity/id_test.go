package mobile

import "testing"

func TestValidate(t *testing.T) {
	validIDs := []string{"210221199906065211", ""}
	invalidIDs := []string{"", "138000000000", "1380000000a", "aaaaaaaaaaa"}

	for _, id := range validIDs {
		if err := Validate(id); err != nil {
			t.Fatal(err)
		}
	}

	for _, mobile := range invalidIDs {
		if err := Validate(mobile); err == nil {
			t.Fatal(err)
		}
	}
}
