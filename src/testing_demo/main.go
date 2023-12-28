package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Post struct {
	Id       int       `json: "id"`
	Content  string    `json: "content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json: "content"`
	Author  string `json:"author"`
}

func decode(filename string) (post Post, err error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening JSON file.", err)
		return
	}

	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&post)
	if err != nil {
		fmt.Println("Error decoding file JSON: ", err)
		return
	}

	return
}

func main() {
	_, err := decode("post.json")
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func unmarshal(filename string) (post Post, err error) {
	jsonfile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}

	defer jsonfile.Close()

	stat, err := jsonfile.Stat()
	if err != nil {
		fmt.Println("failed to get file size:", err)
		return
	}

	jsonData := make([]byte, stat.Size())

	_, err = io.ReadFull(jsonfile, jsonData)
	if err != nil {
		fmt.Println("Error reading JSON data:", err)
		return
	}

	json.Unmarshal(jsonData, &post)
	return
}
