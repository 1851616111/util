package validator

import "testing"

func Test_Validate(t *testing.T) {
	validateIPs := []string{"1.1.1.1", "192.168.0.1", "255.255.255.255"}
	invalidateIPs := []string{"a.1.1.1", "256.256.1.1", "1.1.1"}

	for _, ip := range validateIPs {
		if err := Validate(ip); err != nil {
			t.Fatal(err)
		}
	}

	for _, ip := range invalidateIPs {
		if err := Validate(ip); err == nil {
			t.Fatal(err)
		}
	}
}
