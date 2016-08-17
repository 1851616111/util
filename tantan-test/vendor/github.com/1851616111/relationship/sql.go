package relationship

import "fmt"

func createSQL(table, from_id, to_id string) string {
	return fmt.Sprintf("INSERT INTO %s (from_id, to_id) VALUES (%s, %s);", table, from_id, to_id)
}

func selectSQL(table, from_id, to_id string) string {
	return fmt.Sprintf("SELECT * FROM %s WHERE from_id = %s AND to_id = %s;", table, from_id, to_id)
}

func deleteSQL(table, from_id, to_id string) string {
	return fmt.Sprintf("DELETE FROM %s WHERE from_id = %s AND to_id = %s;", table, from_id, to_id)
}

func unionRelationShipStr(from_id string, tables []string) string {
	n := len(tables)
	if n == 0 {
		return ""
	}

	queryUnion := []byte{}
	for i, table := range tables {
		queryUnion = append(queryUnion, []byte("SELECT to_id FROM "+table+" WHERE from_id = "+from_id)...)
		if i < n-1 {
			queryUnion = append(queryUnion, []byte(" UNION ")...)
		}
	}

	queryUnion = append(queryUnion, []byte(";")...)

	return string(queryUnion)
}
