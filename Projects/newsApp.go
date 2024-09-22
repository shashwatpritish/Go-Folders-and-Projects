package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

type NewsResponse struct {
	Status   string    `json:"status"`
	Total    int       `json:"totalResults"`
	Articles []Article `json:"articles"`
}

func main() {
	apiKey := "xxxxxxxxxxxxxxxxxxxxxxx"          // Your API key here
	query := "Narendra Modi"
	url := fmt.Sprintf("https://newsapi.org/v2/everything?q=%s&apiKey=%s", query, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to fetch news: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to get valid response: %d %s", resp.StatusCode, resp.Status)
	}

	var newsData NewsResponse
	if err := json.NewDecoder(resp.Body).Decode(&newsData); err != nil {
		log.Fatalf("Failed to decode JSON response: %v", err)
	}

	fmt.Printf("Showing top news for '%s':\n\n", query)
	for i, article := range newsData.Articles {
		fmt.Printf("%d. %s\n", i+1, article.Title)
		fmt.Printf("   %s\n", article.Description)
		fmt.Printf("   Read more: %s\n\n", article.URL)
		if i == 5 {
			break
		}
	}
}
