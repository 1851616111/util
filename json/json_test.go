package json

import "testing"

func TestWriteJSONToFile(t *testing.T) {
	m := map[string]string{
		"a": "123",
		"b": "456",
	}

	if err := WriteJSONToFile("test.json", m); err != nil {
		t.Fatal(err)
	}
}
