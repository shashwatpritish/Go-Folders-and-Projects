package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
func main() {
	http.Handle("/", http.FileServer((http.Dir("./template"))))
	http.ListenAndServe(":8080", nil)
}
