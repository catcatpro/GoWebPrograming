package main

import (
	"bytes"
	"encoding/csv"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Post2 struct {
	Id      int
	Content string
	Author  string
}

var PostById map[int]*Post2
var PostByAuthor map[string][]*Post2

func store(post Post2) {
	PostById[post.Id] = &post
	PostByAuthor[post.Author] = append(PostByAuthor[post.Author], &post)
}

func run() {
	PostById = make(map[int]*Post2)
	PostByAuthor = make(map[string][]*Post2)

	post1 := Post2{Id: 1, Content: "Hello World", Author: "Sau Sheong"}
	post2 := Post2{Id: 2, Content: "Bonjour Monde", Author: "Pierre"}
	post3 := Post2{Id: 3, Content: "Hola Mundo", Author: "Padro"}
	post4 := Post2{Id: 4, Content: "Greetings Earthilngs", Author: "Sau Sheong"}

	store(post1)
	store(post2)
	store(post3)
	store(post4)

	fmt.Println(PostById[1])
	fmt.Println(PostById[2])

	for _, post := range PostByAuthor["Sau Sheong"] {
		fmt.Println(post)
	}

	for _, post := range PostByAuthor["Pedro"] {
		fmt.Println(post)
	}
}

func file() {
	data := []byte("Hello, World!")
	err := ioutil.WriteFile("data1", data, 0644)

	if err != nil {
		panic(err)
	}

	read1, _ := ioutil.ReadFile("data1")
	fmt.Println(string(read1))

	file1, _ := os.Create("data2")
	defer file1.Close()

	bytes, _ := file1.Write(data)
	fmt.Printf("Wrote %d bytes to file\n", bytes)

	file2, _ := os.Open("data2")
	defer file2.Close()

	read2 := make([]byte, len(data))
	bytes, _ = file2.Read(read2)
	fmt.Printf("Read %d bytes from file\n", bytes)
	fmt.Println(string(read2))
}

func file_csv() {
	csvFile, err := os.Create("posts.csv")
	if err != nil {
		panic(err)
	}

	defer csvFile.Close()

	allPosts := []Post2{
		Post2{Id: 1, Content: "Hello World", Author: "Sau Sheong"},
		Post2{Id: 2, Content: "Bonjour Monde", Author: "Pierre"},
		Post2{Id: 3, Content: "Hola Mundo", Author: "Padro"},
		Post2{Id: 4, Content: "Greetings Earthilngs", Author: "Sau Sheong"},
	}

	writer := csv.NewWriter(csvFile)
	for _, post := range allPosts {
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()

	file, err := os.Open("posts.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var posts []Post2
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := Post2{Id: int(id), Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}
	fmt.Println(posts[0].Id)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)
}

func store2(data interface{}, filename string) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
	if err != nil {
		panic(err)
	}
}

func load(data interface{}, filename string) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(raw)
	dec := gob.NewDecoder(buffer)
	err = dec.Decode(data)
	if err != nil {
		panic(err)
	}
}

func file_gob() {
	post := Post2{Id: 1, Content: "Hello World", Author: "Sau Sheong"}
	store2(post, "post1")
	var postRead Post2
	load(&postRead, "post1")
	fmt.Println(postRead)
}

func database() {
	post := Post2{Content: "Hello World", Author: "Sau Sheong"}

	fmt.Println(post)
	post.Create()
	fmt.Println(post)

	readPost, _ := GetPost(post.Id)
	fmt.Println(readPost)

	readPost.Content = "Bonjour Monde!"
	readPost.Author = "Pierre"
	readPost.Update()

	posts, _ := Posts(2)
	fmt.Println(posts)
	readPost.Delete()
}
func database2() {
	post := Post3{Content: "Hello, world", Author: "Sau Sheong"}
	post.Create()

	comment := Comment{Content: "Good post!", Author: "Joe", Post: &post}
	comment.Create()
	readPost, _ := GetPost2(post.Id)

	fmt.Println(readPost)
	fmt.Println(readPost.Comments)
	fmt.Println(readPost.Comments[0].Post)

}

func database_for_sqlx() {
	var post Post4
	post, _ = GetPost4(2)
	fmt.Println(post)
}
