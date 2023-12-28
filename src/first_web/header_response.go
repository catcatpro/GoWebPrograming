package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	User    string
	Threads []string
}

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
				<head>
					<title>Go Web Programming</title>
				</head>
				<body>
					<h1>Hello,world</h1>
				</body>
			</html>
	`
	w.Write([]byte(str))

}

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "No such service, try next door")
}

func handleExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://www.baidu.com")
	w.WriteHeader(302)
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	post := &Post{
		User:    "Sau Sheog",
		Threads: []string{"first", "Second", "Third"},
	}

	json, _ := json.Marshal(post)
	w.Write(json)
}
