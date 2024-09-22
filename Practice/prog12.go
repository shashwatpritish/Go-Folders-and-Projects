package main

import (
	"fmt"
	"os"
)

func main() {
	// Create a new file
	file, err := os.OpenFile("newfile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close() // Ensure file is closed

	// Write to the file
	data := []byte("Hello, World!\n")
	_, err = file.WriteString(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("File created and written successfully.")
}
