package user

import "testing"

func Test_unionUsersStr(t *testing.T) {
	s := unionUsersStr([]string{"table-1", "table_2"})
	if s != "SELECT id, name FROM table-1 UNION SELECT id, name FROM table_2;" {
		t.Errorf("test user union err out(%s) != expect(%s)", s, "SELECT id, name FROM table-1 UNION SELECT id, name FROM table_2;")
	}
}
