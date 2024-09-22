package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	out, _ := os.Create("test.go")
	defer out.Close()

	content, _ := http.Get("https://raw.githubusercontent.com/sd2995/GoLang/main/HelloWorldGo/main.go")
	defer content.Body.Close()
	io.Copy(out, content.Body)

	file, _ := os.Open("test.go")

	filedata := make([]byte, 100)

	len, _ := file.Read(filedata)
	fmt.Println(string(filedata[:len]))
}
