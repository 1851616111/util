package router

import "testing"

func Test_Router_Block(t *testing.T) {
	nodes := []string{
		"192.168.0.1",
		"192.168.0.2",
		"192.168.0.3",
	}
	router, err := NewBlockRouter(20000000, 10000000000, nodes, true)
	if err != nil {
		t.Errorf("test Router.Route for block err %v", err)
	}

	in := []string{
		"10000000000",
		"10019999999",
		"10020000000",
		"10039999999",
		"10040000000",
		"10059999999",
		"10060000000",
		"99999999999",
	}
	expect := []string{
		"192.168.0.1",
		"192.168.0.1",
		"192.168.0.2",
		"192.168.0.2",
		"192.168.0.3",
		"192.168.0.3",
		"192.168.0.3",
		"192.168.0.3",
	}

	for i, id := range in {
		out, _ := router.Route(id)
		if out != expect[i] {
			t.Errorf("test Router.Route  for block out(%s) != expect(%s)", out, expect[i])
		}
	}
}
