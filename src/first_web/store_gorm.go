package main

import (
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Post5 struct {
	Id        int
	Content   string
	Author    string     `sql:"not null"`
	Comments  []Comment2 `gorm:"foreignKey:PostId"`
	CreatedAt time.Time
}

type Comment2 struct {
	Id        int
	Content   string
	Author    string `spl:"not null"`
	PostId    int
	CreatedAt time.Time
}

var db2 *gorm.DB

func init() {
	var err error
	db2, err = gorm.Open(postgres.Open("host=localhost user=gwp dbname=gwp password=123456 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db2.AutoMigrate(&Post5{}, &Comment2{})
}

func store_for_gorm() {
	fmt.Println("store_for_gorm")
	post := Post5{Content: "Hello, world", Author: "Sao Shaong"}
	fmt.Println(post)

	db2.Create(&post)
	fmt.Println(post)

	comment := Comment2{Content: "Good post", Author: "joe"}
	db2.Model(&post).Association("Comments").Append(comment)

	var readPost Post5
	db2.Where("author = $1", "Sao Shaong").First(&readPost)
	var comments []Comment2
	db2.Model(&readPost).Find(&comments)
	fmt.Println(comments)
}
