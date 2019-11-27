package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"quotesrestapi/utils"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//PopularQuotes : func(w http.ResponseWriter, r *http.Request) []utils.Quote
func PopularQuotes(w http.ResponseWriter, r *http.Request) {

	var response utils.Resp
	var quotes []utils.Quote

	page, ok := r.URL.Query()["page"]
	if !ok && len(page) < 1 {
		page = []string{"1"}
	}

	format := "json"
	resp, err := http.Get(utils.ParsePopularURL("https://www.goodreads.com/quotes", page, format))
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&response)

	responseReader := strings.NewReader(response.Content)
	doc, err := goquery.NewDocumentFromReader(responseReader)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".quoteContainer").Each(func(i int, s *goquery.Selection) {
		var quote utils.Quote

		img, exist := s.Find(".userIcon").Attr("style")
		if !exist {
			img = "something"
		}

		quote.Img = utils.TransformUrl(img)

		Author := s.Find(".quoteAuthor").Text()
		quote.Author = strings.TrimSpace(Author)

		splitedQuote := s.Find(".quoteBody").Text()
		quote.Quote = strings.TrimSpace(splitedQuote)

		quotes = append(quotes, quote)
	})

	json.NewEncoder(w).Encode(quotes)

}
