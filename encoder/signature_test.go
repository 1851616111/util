package encoder

import "testing"

func TestEncoder_Encode(t *testing.T) {

	en := NewEncoder("mix", func(text, key string) string {
		return Base64(text)
	})

	if en.Method() != "mix" {
		t.Fatal(en)
	}

	if ret := en.Encode("123456", ""); ret != Base64("123456") {
		t.Fatal(ret)
	}
}
