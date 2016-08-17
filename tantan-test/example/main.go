package main

import (
	"fmt"

	"gopkg.in/pg.v4"
)

type User struct {
	Id     int64
	Name   string
	Emails []string
}

func (u User) String() string {
	return fmt.Sprintf("User<%d %s %v>", u.Id, u.Name, u.Emails)
}

type Story struct {
	Id       int64
	Title    string
	AuthorId int64
	Author   *User
}

func (s Story) String() string {
	return fmt.Sprintf("Story<%d %s %s>", s.Id, s.Title, s.Author)
}

func createSchema(db *pg.DB) error {
	queries := []string{
		`CREATE TEMP TABLE users (id serial, name text, emails jsonb)`,
		`CREATE TEMP TABLE stories (id serial, title text, author_id bigint)`,
	}
	for _, q := range queries {
		_, err := db.Exec(q)
		if err != nil {
			return err
		}
	}
	return nil
}

type relation struct {
	From_id string
	To_id   string
}

func main() {
	db := pg.Connect(&pg.Options{
		User:     "tantan007",
		Addr:     "10.1.235.98:5432",
		Database: "db_relationship_like_01",
	})

	rs := []relation{}
	_, err := db.Query(&rs, "SELECT to_id FROM t_relationship_like_01 WHERE from_id = 10000000000 UNION SELECT to_id FROM t_relationship_like_02 WHERE from_id = 10000000000;")
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println(rs)

	//if err != nil {
	//	fmt.Print(err)
	//}
	//users := []*User{}
	//res, err := db.Query(&users, "SELECT id, name FROM t_user_01 UNION SELECT id, name FROM t_user_02;")
	//if err != nil {
	//	fmt.Print(err)
	//}
	//
	//fmt.Println(res.Affected())
	//fmt.Println(users)
	//_, err := db.Exec("SELECT * FROM usertbl;")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//_, err = db.Exec("INSERT INTO usertbl(name, signupdate) VALUES('2222', '2013-12-22');")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//err := createSchema(db)
	//if err != nil {
	//	panic(err)
	//}
	//
	//user1 := &User{
	//	Name:   "admin",
	//	Emails: []string{"admin1@admin", "admin2@admin"},
	//}
	//err = db.Create(user1)
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = db.Create(&User{
	//	Name:   "root",
	//	Emails: []string{"root1@root", "root2@root"},
	//})
	//if err != nil {
	//	panic(err)
	//}
	//
	//story1 := &Story{
	//	Title:    "Cool story",
	//	AuthorId: user1.Id,
	//}
	//err = db.Create(story1)
	//if err != nil {
	//	panic(err)
	//}
	//
	//// Select user by primary key.
	//user := User{Id: user1.Id}
	//err = db.Select(&user)
	//if err != nil {
	//	panic(err)
	//}

	// Select all users.
	//var users []User
	//err = db.Model(&users).Select()
	//if err != nil {
	//	panic(err)
	//}
	//
	//// Select story and associated author in one query.
	//var story Story
	//err = db.Model(&story).
	//	Column("story.*", "Author").
	//	Where("story.id = ?", story1.Id).
	//	Select()
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(user)
	//fmt.Println(users)
	//fmt.Println(story)
	//// Output: User<1 admin [admin1@admin admin2@admin]>
	// [User<1 admin [admin1@admin admin2@admin]> User<2 root [root1@root root2@root]>]
	// Story<1 Cool story User<1 admin [admin1@admin admin2@admin]>>
}
