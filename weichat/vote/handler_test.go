package vote

import (
	"testing"
)

func TestNewRouter(t *testing.T) {
	var err error
	dbI, err = NewDBInterface("192.168.199.216", "5432", "postgres", "postgresql2016", "pinto")
	if err != nil {
		t.Fatal(err)
	}

	NewRouter()
}
