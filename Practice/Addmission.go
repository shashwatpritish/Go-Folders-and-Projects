package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func createFile(data string) {
	file, err := os.OpenFile("Data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Write some content to the file
	if _, err := file.WriteString(data); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Content appended successfully!")
}

type Student struct {
	name   string
	age    int
	class  int
	gender string
}

func studentAdmission(name string, age int, class int, gender string) {
	var student Student
	student.name = name
	student.age = age
	student.class = class
	student.gender = gender
}

func main() {
	var name string
	var age int
	var class int
	var gender string

	fmt.Println("Enter name of student")
	fmt.Scan(&name)
	fmt.Println("Enter age of student")
	fmt.Scan(&age)
	fmt.Println("Enter class of student")
	fmt.Scan(&class)
	fmt.Println("Enter gender of student")
	fmt.Scan(&gender)

	profile := "Name: " + string(name) + "\n" + "Age: " + strconv.Itoa(age) + "\n" + "Class: " + strconv.Itoa(class) + "\n" + "Gender: " + string(gender)
	createFile(profile + "\n-------\n")
}
