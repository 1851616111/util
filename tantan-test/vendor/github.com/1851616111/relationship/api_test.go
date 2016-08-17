package relationship

import (
	"github.com/1851616111/tantan-test/db"
	"testing"
)

func Test_Like(t *testing.T) {
	db, err := db.NewDBFromJsonFile("../conf.json")
	if err != nil {
		t.Errorf("new db err %v\n", err)
	}

	if err := NewRelationShipClient(db).Like().Create("10000000000", "10000000001"); err != nil {
		if err != nil {
			t.Errorf("create like from %s to %s err %v\n", "10000000000", "10000000001", err)
		}
	}
	if bool, err := NewRelationShipClient(db).Like().IfLike("10000000000", "10000000001"); err != nil {
		if err != nil {
			t.Errorf("get like from %s to %s err %v\n", "10000000000", "10000000001", err)
		}
	} else if !bool {
		t.Errorf("get like from %s to %s err %v\n", "10000000000", "10000000001", err)
	}

	if err := NewRelationShipClient(db).Like().Delete("10000000000", "10000000001"); err != nil {
		if err != nil {
			t.Errorf("create like from %s to %s err %v\n", "10000000000", "10000000001", err)
		}
	}
}

func Test_Match(t *testing.T) {
	db, err := db.NewDBFromJsonFile("../conf.json")
	if err != nil {
		t.Errorf("new db err %v\n", err)
	}

	if err := NewRelationShipClient(db).Match().Create("10000000000", "10000000001"); err != nil {
		if err != nil {
			t.Errorf("create match from %s to %s err %v\n", "10000000000", "10000000001", err)
		}
	}
	if bool, err := NewRelationShipClient(db).Match().IfMatch("10000000000", "10000000001"); err != nil {
		if err != nil {
			t.Errorf("get match from %s to %s err %v\n", "10000000000", "10000000001", err)
		}
	} else if !bool {
		t.Errorf("get match from %s to %s err %v\n", "10000000000", "10000000001", err)
	}

	if err := NewRelationShipClient(db).Match().Delete("10000000000", "10000000001"); err != nil {
		if err != nil {
			t.Errorf("create match from %s to %s err %v\n", "10000000000", "10000000001", err)
		}
	}
}

func Test_Dislike(t *testing.T) {
	db, err := db.NewDBFromJsonFile("../conf.json")
	if err != nil {
		t.Errorf("new db err %v\n", err)
	}

	if err := NewRelationShipClient(db).Dislike().Create("10000000000", "10000000001"); err != nil {
		if err != nil {
			t.Errorf("create dislike from %s to %s err %v\n", "10000000000", "10000000001", err)
		}
	}

	if err := NewRelationShipClient(db).Dislike().Delete("10000000000", "10000000001"); err != nil {
		if err != nil {
			t.Errorf("create dislike from %s to %s err %v\n", "10000000000", "10000000001", err)
		}
	}
}
