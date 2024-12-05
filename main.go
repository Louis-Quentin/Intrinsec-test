package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func displayResults(data [][]string) {

	fmt.Println("Results");
	for _, result := range data {
		link := result[1];
		title := result[2];
		fmt.Printf("Title = %s\nLien = %s\n\n", title, link);
	}
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
	displayResults(res);
}