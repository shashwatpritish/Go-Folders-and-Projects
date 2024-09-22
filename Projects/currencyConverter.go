package main

import (
	"fmt"
)

func main() {
	var choice int
	fmt.Println("Press 1 to convert to rupees and 2 to convert to dollar")
	fmt.Scan(&choice)

	if choice == 1 {
		var amount float64
		fmt.Println("Enter in dollars")
		fmt.Scan(&amount)
		rupeesAmount := amount * 83.88
		fmt.Printf("$ %v is ₹ %v", amount, rupeesAmount)
	}
	if choice == 2 {
		var amount float64
		fmt.Println("Enter in rupees")
		fmt.Scan(&amount)
		dollarAmount := amount / 83.88
		fmt.Printf("₹ %v is $ %.2f", amount, dollarAmount)
	}
}
