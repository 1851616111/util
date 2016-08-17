package router

import "testing"

func Test_NewShard(t *testing.T) {
	hostConfig := &Config{
		Begin: 100000,
		Size:  5000,
		Nodes: []string{
			"192.168.0.1",
			"192.168.0.2",
			"192.168.0.3",
		},
		Compatible: true,
	}
	dbConfig := &Config{
		Nodes: []string{
			"db_user_01",
			"db_user_02",
			"db_user_03",
		},
	}
	tableConfig := &Config{
		Nodes: []string{
			"t_user_01",
			"t_user_02",
			"t_user_03",
		},
	}

	router, err := NewRouter(hostConfig, dbConfig, tableConfig)
	if err != nil {
		t.Errorf("test NewRouter err %v\n", err)
	}

	in := [][]string{
		idGen(15000),
	}

	for i := range in {
		ids := in[i]

		statisResult := map[string]int{}
		for _, id := range ids {
			addr, err := router.Route(id)
			if err != nil {
				t.Errorf("test Router.Route err %v", err)
			}

			statis(addr, id, statisResult)
		}

		t.Logf("idGen(%d)\n", len(ids))

		printStatis(statisResult)

	}

}
