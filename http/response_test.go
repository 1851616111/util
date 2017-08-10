package http

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestConvert(t *testing.T) {

	type student struct {
		Name string `json:"Name"`
		Age  int    `json:"Age"`
	}

	type Data struct {
		Data student `json:"data"`
	}

	m := map[string]interface{}{
		"code":    200,
		"message": "ok",
		"data": student{
			Name: "michael",
			Age:  20,
		},
	}

	var data Data

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(m); err != nil {
		t.Fatal(err)
	}

	if err := ReadToTarget(&buf, &data, nil); err != nil {
		t.Fatal(err)
	}
}
