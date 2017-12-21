package main

import (
	"fmt"
	"net/http"
	"html/template"
	"os"
	"io/ioutil"
	"encoding/xml"
	"sync"
)

type CSiteMapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type CNews struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type CNewsMap struct {
	KeyWord string
	Location string
}

type CNewsAggPage struct {
	Title string
	News map[string]CNewsMap
}


// A CONCURRENT NEWS AGGREGATION WEB APP
// ALL NEWS CATEGORIES ARE PARSED CONCURRENTLY IN GOROUTINES.



var cNewsAggWaitGroup sync.WaitGroup

func newsRoutine(newsChannel chan CNews, location string) {
	defer cNewsAggWaitGroup.Done()
	news := CNews{}

	resp, _ := http.Get(location)
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	xml.Unmarshal(bytes, &news)

	// put news into channel
	newsChannel <- news
}

func cNewsAggHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Don't hack direct path to file...
	template, err := template.ParseFiles("/home/gareth/go/src/github.com/brycelol/learning-golang/basictemplating.html")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Filepath was probably hard coded - switch it for your dev environment.")
		os.Exit(1)
	}

	siteMap := CSiteMapIndex{}

	resp, _ := http.Get("https://washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	xml.Unmarshal(bytes, &siteMap)

	newsMap := make(map[string]CNewsMap)

	// create a channel to house the news (basically thread safe holder...)
	newsChannel := make(chan CNews, 30)

	for _, location := range siteMap.Locations {
		// COLLECT EVERYTHING CONCURRENTLY
		cNewsAggWaitGroup.Add(1)
		go newsRoutine(newsChannel, location)
	}

	cNewsAggWaitGroup.Wait()
	close(newsChannel)

	for news := range newsChannel {
		for index, _ := range news.Titles {
			newsMap[news.Titles[index]] = CNewsMap{news.Keywords[index], news.Locations[index]}
		}
	}

	newsAggPage := CNewsAggPage{Title: "A News Aggregator!", News: newsMap}

	template.Execute(w, newsAggPage)
}

func cIndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to a web page</h1>")
}

func main() {
	http.HandleFunc("/", cIndexHandler)
	http.HandleFunc("/agg", cNewsAggHandler)
	http.ListenAndServe(":8000", nil)
}