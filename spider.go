package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/xml")

type SiteMapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	KeyWord string
	Location string
}

func main() {
	// Load the news sitemap XML
	siteMap := SiteMapIndex{}

	resp, _ := http.Get("https://washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	xml.Unmarshal(bytes, &siteMap)

	news := News{}
	newsMap := make(map[string]NewsMap)

	for _, location := range siteMap.Locations {
		fmt.Printf("\n%s", location)

		resp, _ := http.Get(location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		xml.Unmarshal(bytes, &news)

		// loop through all news articles.
		for index, _ := range news.Titles {
			// add map entry containing keyword and location.
			newsMap[news.Titles[index]] = NewsMap{news.Keywords[index], news.Locations[index]}
		}
	}

	for idx, data := range newsMap {
		fmt.Println("\n\n\n", idx)
		fmt.Println("\n KEYWORD: ", data.KeyWord)
		fmt.Println("\n LOCATION: ", data.Location)
	}
}

// [5 5]type == array
// []int == slice
// fyi := means we can modify var.
