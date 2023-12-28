package main

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
)

var Db *sql.DB

type Post3 struct {
	Id       int
	Content  string
	Author   string
	Comments []Comment
}

type Comment struct {
	Id      int
	Content string
	Author  string
	Post    *Post3
}

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=123456 sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func Posts(limit int) (posts []Post2, err error) {
	rows, err := Db.Query("select id, content,author from posts limit $1", limit)
	if err != nil {
		return
	}

	for rows.Next() {
		post := Post2{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

func GetPost(id int) (post Post2, err error) {
	post = Post2{}
	err = Db.QueryRow("select id, content,author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

func (post *Post2) Create() (err error) {
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	stat, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stat.Close()
	err = stat.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

func (post *Post2) Update() (err error) {
	_, err = Db.Exec("update posts set content = $2, author = $3 where id = $1", post.Id, post.Content, post.Author)
	return
}

func (post *Post2) Delete() (err error) {
	_, err = Db.Exec("delete from posts where id = $1", post.Id)
	return
}

func (comment *Comment) Create() (err error) {
	if comment.Post == nil {
		err = errors.New("Post not found")
		return
	}

	err = Db.QueryRow("insert into comments (content, author, post_id) values ($1, $2, $3) returning id", comment.Id, comment.Author, comment.Post.Id).Scan(&comment.Id)
	return
}

func GetPost2(id int) (post Post3, err error) {
	post = Post3{}
	post.Comments = []Comment{}
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)

	rows, err := Db.Query("select id, content, author from comments")
	if err != nil {
		return
	}

	for rows.Next() {
		comment := Comment{Post: &post}
		err = rows.Scan(&comment.Id, &comment.Content, &comment.Author)
		if err != nil {
			return
		}
		post.Comments = append(post.Comments, comment)
	}

	rows.Close()
	return

}

func (post *Post3) Create() (err error) {
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	stat, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stat.Close()
	err = stat.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}
