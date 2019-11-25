package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

// Quote : quote struct
type Quote struct {
	Img    string `json:"img"`
	Author string `json:"author"`
	Quote  string `json:"quote"`
}

func main() {
	var quote Quote
	var exist bool
	// var quotes []Quote

	req, err := http.NewRequest("GET", "https://www.goodreads.com/quotes/search", nil)
	if err != nil {
		log.Print((err))
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("q", "walden")
	q.Add("page", "1")
	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())
	resp, err := http.Get(req.URL.String())
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".quote.mediumText").Each(func(i int, s *goquery.Selection) {
		quote.Author = s.Find(".authorOrTitle").Text()
		quote.Img, exist = s.Find("img").Attr("src")
		if !exist {
			quote.Img = "https://something"
		}
		fmt.Printf("Review %d: %s %s \n", i, quote.Author, quote.Img)
	})
}
