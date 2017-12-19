package main

import (
	"fmt"
	"net/http"
	"html/template"
	"os"
	"io/ioutil"
	"encoding/xml"
)

// CRAWLS THE WALL STREET JOURNAL SITEMAP AND BUILDS AN AGGREGATED TABLE FOR VIEWING

type GSiteMapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type GNews struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type GNewsMap struct {
	KeyWord string
	Location string
}

type NewsAggPage struct {
	Title string
	News map[string]GNewsMap
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Don't hack direct path to file...
	template, err := template.ParseFiles("/home/gareth/go/src/github.com/brycelol/hello/basictemplating.html")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Filepath was probably hard coded - switch it for your dev environment.")
		os.Exit(1)
	}

	siteMap := GSiteMapIndex{}

	resp, _ := http.Get("https://washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	xml.Unmarshal(bytes, &siteMap)

	news := GNews{}
	newsMap := make(map[string]GNewsMap)

	for _, location := range siteMap.Locations {
		fmt.Printf("\n%s", location)

		resp, _ := http.Get(location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		xml.Unmarshal(bytes, &news)

		// loop through all news articles.
		for index, _ := range news.Titles {
			// add map entry containing keyword and location.
			newsMap[news.Titles[index]] = GNewsMap{news.Keywords[index], news.Locations[index]}
		}
	}

	newsAggPage := NewsAggPage{Title: "A News Aggregator!", News: newsMap}

	template.Execute(w, newsAggPage)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to a web page</h1>")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg", newsAggHandler)
	http.ListenAndServe(":8000", nil)
}