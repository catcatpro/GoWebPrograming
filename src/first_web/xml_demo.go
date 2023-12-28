package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Post6 struct {
	XMLName xml.Name  `xml:"post"`
	Id      string    `xml:"id,attr"`
	Content string    `xml:"content"`
	Author  Author    `xml:author`
	Xml     string    `xml:",innerxml"`
	Comment []comment `xml:"comments>comment"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

type comment struct {
	Id      string `xml:"id,attr"`
	Content string `xml:"content"`
	Author  Author `xml:author`
}

func decode_xml() {
	xmlFile, err := os.Open("post.xml")
	if err != nil {
		fmt.Println("Error opening XML file", err)
		return
	}

	defer xmlFile.Close()
	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error opening XML data", err)
		return
	}

	var post Post6
	xml.Unmarshal(xmlData, &post)
	fmt.Println(post)
}

func decoder_decode_xml() {
	xmlFile, err := os.Open("post.xml")
	if err != nil {
		fmt.Println("Error opening XML file", err)
		return
	}

	defer xmlFile.Close()
	decoder := xml.NewDecoder(xmlFile)
	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decodeing XML into tokens: ", err)
			return
		}
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "comment" {
				var comment_data comment
				decoder.DecodeElement(&comment_data, &se)
				fmt.Println(comment_data)
			}
		}
	}

}

func encode_xml() {
	post := Post6{
		Id:      "1",
		Content: " Hello, World",
		Author: Author{
			Id:   "2",
			Name: "Sau Sheong",
		},
	}

	// output, err := xml.Marshal(post)
	output, err := xml.MarshalIndent(post, "", "\t")
	if err != nil {
		fmt.Println("Error marshalling to XML:", err)
		return
	}

	// err = ioutil.WriteFile("post_encode.xml", output, 0644)
	err = ioutil.WriteFile("post_encode.xml", []byte(xml.Header+string(output)), 0644)
	if err != nil {
		fmt.Println("Error writing XML to file:", err)
	}
}

func decoder_encode_xml() {
	post := Post6{
		Id:      "1",
		Content: " Hello, World",
		Author: Author{
			Id:   "2",
			Name: "Sau Sheong",
		},
	}

	xmlFile, err := os.Create("post2.xml")
	if err != nil {
		fmt.Println("Error creating XML file: ", err)
		return
	}

	encoder := xml.NewEncoder(xmlFile)
	encoder.Indent("", "\t")
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("Error encoding XML to file: ", err)
		return
	}
}
