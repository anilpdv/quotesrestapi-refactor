package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

// Quote : quote struct
type Quote struct {
	Img    string `json:"img"`
	Author string `json:"author"`
	Quote  string `json:"quote"`
}

// Reqinfo : reqinfo struct
type Reqinfo struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	NumResults int `json:"num_results"`
	TotalPages int `json:"total_pages"`
}

// Respquotes : Respquotes struct
type Respquotes struct {
	Quotes []Quote `json:"quotes"`
	Info   Reqinfo `json:"req_info"`
}

// URL : NEW REQ url
type URL string

// Parse : interface
type Parse interface {
	ParseURL(p string) string
	String() string
}

// String : receiver
func (u URL) String() string {
	return string(u)
}

// ParseURL : receiver
func (u URL) ParseURL(p string) URL {
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		log.Print((err))
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("page", p)
	req.URL.RawQuery = q.Encode()

	return URL(req.URL.String())
}

// RandomQuotes : handlefunc
func RandomQuotes(w http.ResponseWriter, r *http.Request) {
	var quotes Respquotes
	var url URL

	rand.Seed(time.Now().UnixNano())
	rpage := randomInt(0, 99)
	rquote := randomInt(0, 29)
	url = "http://www.quotesapi.ml/quotes"

	url = url.ParseURL(string(rpage))
	_, _ = rpage, rquote
	resp, err := http.Get(url.String())
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&quotes)

	json.NewEncoder(w).Encode(&quotes.Quotes[rquote])
}

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}
