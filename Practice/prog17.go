package main

import "fmt"

func generic_circumference[r int | float32](radius r) {

	c := 2 * 3 * radius
	fmt.Println("The circumference is: ", c)

}

func main() {
	var radius1 int = 8
	var radius2 float32 = 9.5

	generic_circumference(radius1)
	generic_circumference(radius2)
}
