package main

import (
	"encoding/json"
	"os"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

type Article struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

func saveAsJSON(articles []Article) {

	filename := "Articles.json";
	file, err := os.Create(filename)

	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(articles);

	if err != nil {
		log.Fatalf("Error encoding JSON: %v", err)
	}
}

func saveResults(data [][]string) {

	var articles []Article;

	for _, result := range data {
		link := result[1];
		title := result[2];
		fmt.Printf("Title = %s\nLien = %s\n\n", title, link);

		articles = append(articles, Article{
			Title: title,
			Link:  link,
		});
	}
	saveAsJSON(articles);
}

func parseData(body string) [][]string {

	regex := regexp.MustCompile(`<h2[^>]*>\s*<a[^>]*href=["'](.*?)["'][^>]*>(.*?)</a>\s*</h2>`);
	res := regex.FindAllStringSubmatch(body, -1);

	return res;
}

func fetchData() [][]string {

	url := "https://www.intrinsec.com/blog/";

	resp, err := http.Get(url);
	if err != nil {
		log.Fatalf("Error fetching data : %v", err);
	}
	defer resp.Body.Close();

	body, err := ioutil.ReadAll(resp.Body);
	if err != nil {
		log.Fatalf("Error reading resp : %v", err);
	}

	html := string(body);
	res := parseData(html);

	return res;
}

func main() {

	res := fetchData();
	saveResults(res);
}