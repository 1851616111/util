package validator

import "testing"

func Test_Validate(t *testing.T) {
	validatePorts := []string{"22", "80", "1080", "65535"}
	invalidatePorts := []string{"", "-1", "65536", "0", "1a"}

	for _, port := range validatePorts {
		if err := Validate(port); err != nil {
			t.Fatal(err)
		}
	}

	for _, port := range invalidatePorts {
		if err := Validate(port); err == nil {
			t.Fatal(port, err)
		}
	}
}
