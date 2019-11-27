package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"quotesrestapi/utils"
	"reflect"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// SearchQuotes : func(w http.ResponseWriter, r *http.Request) []Quote
func SearchQuotes(w http.ResponseWriter, r *http.Request) {
	var exist bool
	var quotes []utils.Quote

	query, ok := r.URL.Query()["q"]
	fmt.Println(reflect.TypeOf(query), ok)
	if !ok && len(query) < 1 {
		http.Error(w, "should provide a query", 400)
		return

	}

	page, ok := r.URL.Query()["page"]
	if !ok && len(page) < 1 {
		page = []string{"1"}
	}

	resp, err := http.Get(utils.ParseURL("https://www.goodreads.com/quotes/search", query, page))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".quote.mediumText").Each(func(i int, s *goquery.Selection) {

		var quote utils.Quote

		Author := s.Find(".authorOrTitle").Text()
		quote.Author = strings.TrimSpace(Author)

		quote.Img, exist = s.Find("img").Attr("src")
		if !exist {
			quote.Img = "https://something"
		}

		text := s.Find(".quoteText").Text()
		splitedQuote := strings.Split(text, "\n")[1]
		quote.Quote = strings.TrimSpace(splitedQuote)

		quotes = append(quotes, quote)
	})

	json.NewEncoder(w).Encode(quotes)
}
