package main

import "fmt"

func main() {
	a := 10
	b := &a

	fmt.Println(*b)
}
