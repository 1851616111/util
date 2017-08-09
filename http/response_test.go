package http

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestConvert(t *testing.T) {
	type student struct {
		Name string
		Age  int
	}

	var stu student

	m := map[string]interface{}{
		"code":    0,
		"message": "ok",
		"data": student{
			Name: "michael",
			Age:  20,
		},
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(m); err != nil {
		t.Fatal(err)
	}

	if err := ReadToTarget(&buf, &stu, nil); err != nil {
		t.Fatal(err)
	}

	if stu.Name != "michael" || stu.Age != 20 {
		t.Fatal(stu)
	}
}
