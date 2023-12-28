package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type post struct {
	Id       int        `json: "id"`
	Content  string     `json: "content"`
	Author   author2    `json:"author"`
	Comments []comment2 `json:"comments"`
}

type author2 struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type comment2 struct {
	Id      int    `json:"id"`
	Content string `json: "content"`
	Author  string `json:"author"`
}

func decode_json() {
	jsonFile, err := os.Open("post.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}

	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON data:", err)
		return
	}
	var post post
	json.Unmarshal(jsonData, &post)
	fmt.Println(post)
}

func decoder_decode_json() {
	jsonFile, err := os.Open("post.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}

	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	for {
		var post post
		err := decoder.Decode(&post)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decodeing JSON:", err)
			return
		}
		fmt.Println(post)
	}
}

func encode_json() {
	post := post{
		Id:      1,
		Content: "Hello, world",
		Author: author2{
			Id:   2,
			Name: "Sau Sheong",
		},
		Comments: []comment2{
			comment2{
				Id:      3,
				Content: "Have a grest day!",
				Author:  "Adam",
			},
			comment2{
				Id:      4,
				Content: "How are you today!",
				Author:  "Betty",
			},
		},
	}
	output, err := json.MarshalIndent(&post, "", "\t\t")
	if err != nil {
		fmt.Println("Error marshalling to JSSON:", err)
		return
	}
	err = ioutil.WriteFile("post2.json", output, 0644)
	if err != nil {
		fmt.Println("Error writeing JSON to file:", err)
		return
	}

}

func decoder_encode_json() {
	post := post{
		Id:      1,
		Content: "Hello, world",
		Author: author2{
			Id:   2,
			Name: "Sau Sheong",
		},
		Comments: []comment2{
			comment2{
				Id:      3,
				Content: "Have a grest day!",
				Author:  "Adam",
			},
			comment2{
				Id:      4,
				Content: "How are you today!",
				Author:  "Betty",
			},
		},
	}

	jsonFile, err := os.Create("post3.json")
	if err != nil {
		fmt.Println("Error createing JSON file:", err)
		return
	}
	defer jsonFile.Close()
	encoder := json.NewEncoder(jsonFile)
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("Error encoding JSON to file:", err)
		return
	}
}
