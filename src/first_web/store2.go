package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Post4 struct {
	Id         int
	Content    string
	Author     string
	AuthorName string `db: author`
}

var db *sqlx.DB

func init() {
	var err error
	db, err = sqlx.Connect("postgres", "user=gwp dbname=gwp password=123456 sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func GetPost4(id int) (post Post4, err error) {
	post = Post4{}
	err = db.QueryRowx("select id, content,author from posts where id = $1", id).StructScan(&post)
	if err != nil {
		return
	}
	return
}
