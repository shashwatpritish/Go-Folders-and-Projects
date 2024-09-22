package main

import (
	"fmt"
	"github.com/russross/blackfriday/v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Function to render Markdown file as HTML
func renderMarkdown(filename string) (string, error) {
	// Read the Markdown file
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	// Convert Markdown content to HTML using blackfriday
	html := blackfriday.Run(content)

	// Wrap the HTML content in a basic HTML structure
	return fmt.Sprintf(`
        <!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>Markdown Viewer</title>
        </head>
        <body>
            %s
        </body>
        </html>
    `, html), nil
}

// HTTP handler to serve the HTML content
func markdownHandler(w http.ResponseWriter, r *http.Request) {
	filename := "README.md" // Change this to the markdown file you want to display

	htmlContent, err := renderMarkdown(filename)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading file: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(htmlContent))
}

func main() {
	// Verify if Markdown file is passed as an argument
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <markdown-file>")
	}

	filename := os.Args[1]

	// Check if the file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		log.Fatalf("File does not exist: %s", filename)
	}

	// Serve the converted Markdown on a local web server
	http.HandleFunc("/", markdownHandler)
	fmt.Println("Starting server at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
