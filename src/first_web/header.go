package main

import (
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header["Accept-Encoding"]
	fmt.Fprintln(w, h)
}
