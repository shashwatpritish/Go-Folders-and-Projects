package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func handle(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100)
	fmt.Fprintf(w, "Random Number: %d", randomNumber)
}
func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8080", nil)
}
