package main

import "fmt"

func main() {
	var day string
	fmt.Println("Enter the day")
	fmt.Scanf("%s", &day)
	switch day {
	case "Monday":
		fmt.Println("Som war")
	case "Tuesday":
		fmt.Println("Mangal war")
	case "Wednesday":
		fmt.Println("World war")
	case "Thursday":
		fmt.Println("Shyam war")
	case "Friday":
		fmt.Println("Paka war")
	case "Saturday":
		fmt.Println("Kalinga war")
	case "Sunday":
		fmt.Println("Ravi war")

	}
}
