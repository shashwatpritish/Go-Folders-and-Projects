package main

import "os"

func main() {
	out, _ := os.Create("main2.go")
	defer out.Close()
}
