package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func main() {
	http.HandleFunc("/index", index)
	http.ListenAndServe(":8080", nil)
}
