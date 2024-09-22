package main

import "fmt"

type Person struct {
	id      int
	name    string
	age     int
	gender  string
	isCoder bool
	isGamer bool
}

func main() {
	var emp1 Person
	var emp2 Person

	emp1.id = 1
	emp1.name = "Shashwat Pritish"
	emp1.age = 9
	emp1.gender = "Male"
	emp1.isCoder = true
	emp1.isGamer = false

	emp2.id = 2
	emp2.name = "Samanyu Pritish"
	emp2.age = 5
	emp2.gender = "Male"
	emp2.isCoder = false
	emp2.isGamer = true

	fmt.Println("ID: ", emp2.id)
	fmt.Println("Name: ", emp2.name)
	fmt.Println("Age:", emp2.age)
	fmt.Println("gender:", emp2.gender)
	fmt.Println("isCoder :", emp2.isCoder)
	fmt.Println("isGamer:", emp2.isGamer)
}
