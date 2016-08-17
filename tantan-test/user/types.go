package user

import "fmt"

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (u *User) create(table string) string {
	return fmt.Sprintf("INSERT INTO %s (id,name) VALUES (%s,'%s');", table, u.Id, u.Name)
}

type UserList []*User
