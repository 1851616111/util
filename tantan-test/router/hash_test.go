package router

import (
	"fmt"
	"sort"
	"testing"
)

func Test_Route_Hash(t *testing.T) {
	nodes := [][]string{
		[]string{
			"db_user_01",
		},
		[]string{
			"db_user_01",
			"db_user_02",
			"db_user_03",
		},
		[]string{
			"db_user_01",
			"db_user_02",
			"db_user_03",
			"db_user_04",
			"db_user_05",
			"db_user_06",
			"db_user_07",
			"db_user_08",
			"db_user_09",
			"db_user_10",
		},
		[]string{
			"db_user_01",
			"db_user_02",
			"db_user_03",
			"db_user_04",
			"db_user_05",
			"db_user_06",
			"db_user_07",
			"db_user_08",
			"db_user_09",
			"db_user_10",
			"db_user_11",
			"db_user_12",
			"db_user_13",
			"db_user_14",
			"db_user_15",
			"db_user_16",
			"db_user_17",
			"db_user_18",
			"db_user_19",
			"db_user_20",
		},
	}

	in := [][]string{
		idGen(1),
		idGen(100),
		idGen(1000),
		idGen(10000),
	}

	for i := range nodes {
		ids := in[i]

		router, err := NewHashRouter(nodes[i])
		if err != nil {
			t.Errorf("test Router.Route for hash err %v", err)
		}

		statisResult := map[string]int{}

		for _, id := range ids {
			node, err := router.Route(id)
			if err != nil {
				t.Errorf("test Router.Route for hash err %v", err)
			}
			statis(node, id, statisResult)
		}

		t.Logf("idGen(%d) ---> nodes(%d)\n", len(ids), len(nodes))
		printStatis(statisResult)
	}
}

func idGen(num int) []string {
	l := []string{}
	for i := 100000; i < 100000+num; i++ {
		l = append(l, fmt.Sprintf("%d", i))
	}
	return l
}

func statis(node, key string, result map[string]int) {
	if _, ok := result[node]; !ok {
		result[node] = 1
		return
	}

	result[node] += 1
}

func printStatis(result map[string]int) {
	rows := sortStringList{}
	for node, num := range result {
		rows = append(rows, fmt.Sprintf("statis: node(%s) ---> id(%d)\n", node, num))
	}

	sort.Sort(rows)
	println("-----------------------------------")
	for i := range rows {
		fmt.Printf("%d:\t %s", i+1, rows[i])
	}
}

type sortStringList []string

func (l sortStringList) Len() int {
	return len(l)
}

func (l sortStringList) Less(i, j int) bool {
	return l[i] < l[j]
}

func (l sortStringList) Swap(i, j int) {
	tmp := l[i]
	l[j], l[i] = tmp, l[j]
}
