package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	var num int
	var compNum int
	tries := 10

	compNum = rand.Intn(10)

	for i := 0; i < 10; i++ {
		fmt.Println("Enter a number b/w 0 to 10")
		fmt.Scan(&num)
		if num == compNum {
			fmt.Println("You win")
			os.Exit(0)
		} else {
			fmt.Println("Try again")
		}
		tries += i
	}
	fmt.Println("Game over!")
}
