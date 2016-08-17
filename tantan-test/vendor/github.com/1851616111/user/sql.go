package user

func unionUsersStr(tables []string) string {
	n := len(tables)
	if n == 0 {
		return ""
	}

	queryUnion := []byte{}
	for i, table := range tables {
		queryUnion = append(queryUnion, []byte("SELECT id, name FROM "+table)...)
		if i < n-1 {
			queryUnion = append(queryUnion, []byte(" UNION ")...)
		}
	}

	queryUnion = append(queryUnion, []byte(";")...)

	return string(queryUnion)
}
