package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("-*-*-*-*-*-*-*-*-*Welcome to Todo-*-*-*-*-*-*-*-*-*")
	tasks := make([]string, 0)
	for {
		fmt.Println("------------------------")
		fmt.Println("Enter your query")
		fmt.Println("1. Add task")
		fmt.Println("2. List tasks")
		fmt.Println("3. Quit")
		fmt.Println("------------------------")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var task string

			fmt.Println("What is your task")
			fmt.Scan(&task)
			tasks = append(tasks, task)
		case 2:
			fmt.Println("Your task")
			for i := 0; i <= len(tasks)-1; i++ {
				fmt.Println(tasks[i])
			}
			//fmt.Println(tasks)
		case 3:
			os.Exit(0)
		}
	}
}
