package main

import (
	"encoding/json"
	"os"
//	"fmt"
	"testing"
)

func TestParseData(t *testing.T) {
	html := `
		<h2><a href="/article1">Article 1</a></h2>
		<h2><a href="/article2">Article 2</a></h2>
	`
	expected := [][]string{
		{"", "/article1", "Article 1"},
		{"", "/article2", "Article 2"},
	}

	result := parseData(html)
	if len(result) != len(expected) {
		t.Errorf("Expected %d articles, got %d", len(expected), len(result))
	}

	for i, article := range result {
		if article[1] != expected[i][1] || article[2] != expected[i][2] {
			t.Errorf("Expected %v, got %v", expected[i], article)
		}
	}
}

func TestSaveAsJSON(t *testing.T) {
	articles := []Article{
		{Title: "Article 1", Link: "/article1"},
		{Title: "Article 2", Link: "/article2"},
	}

	saveAsJSON(articles)

	data, err := os.ReadFile("Articles.json")
	if err != nil {
		t.Fatalf("Error reading temp file: %v", err)
	}

	var decodedArticles []Article
	err = json.Unmarshal(data, &decodedArticles)
	if err != nil {
		t.Fatalf("Error decoding JSON: %v", err)
	}

	if len(decodedArticles) != len(articles) {
		t.Errorf("Expected %d articles, got %d", len(articles), len(decodedArticles))
	}
	for i, article := range decodedArticles {
		if article.Title != articles[i].Title || article.Link != articles[i].Link {
			t.Errorf("Expected %v, got %v", articles[i], article)
		}
	}
}