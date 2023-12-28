package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm()
	// fmt.Fprintln(w, r.Form)
	// r.ParseMultipartForm(1024)
	// fmt.Fprintln(w, r.MultipartForm)

	// fmt.Fprintln(w, r.PostFormValue("user"))
	// fmt.Fprintln(w, r.Form)
	// fmt.Fprintln(w, "(1)", r.FormValue("user"))
	// fmt.Fprintln(w, "(2)", r.PostFormValue("user"))
	// fmt.Fprintln(w, "(3)", r.PostForm)
	// fmt.Fprintln(w, "(4)", r.MultipartForm)

	// r.ParseMultipartForm(1024)
	// fileHeader := r.MultipartForm.File["uploaded"][0]
	// file, err := fileHeader.Open()
	// if err == nil {
	// 	data, err := ioutil.ReadAll(file)
	// 	if err == nil {
	// 		fmt.Fprintln(w, string(data))
	// 	}
	// }

	file, _, err := r.FormFile("uploaded")
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}

}
