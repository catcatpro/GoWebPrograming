package main

import (
	"database/sql"
	"fmt"
)

type Text interface {
	fetch(id int) (err error)
	create() (err error)
	update() (err error)
	delete() (err error)
}

var Db2 *sql.DB

func init() {
	var err2 error
	Db2, err2 = sql.Open("postgres", "user=gwp dbname=gwp password=123456 sslmode=disable")
	if err2 != nil {
		panic(err2)
	}
}

func retrieve(id int) (post Post7, err error) {
	post = Post7{}
	err = Db2.QueryRow("select id,content, author from post7 where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

func (post *Post7) fetch(id int) (err error) {
	err = Db2.QueryRow("select id,content, author from post7 where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

func (post *Post7) create() (err error) {
	statement := "insert into post7 (content, author) values ($1, $2) returning id"
	stmt, err := Db2.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	fmt.Println(post)
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

func (post *Post7) update() (err error) {
	_, err = Db2.Exec("update post7 set content = $3, author = $2 where id = $1", post.Id, post.Author, post.Content)
	return
}

func (post Post7) delete() (err error) {
	_, err = Db2.Exec("delete from post7 where id = $1", post.Id)
	return

}
