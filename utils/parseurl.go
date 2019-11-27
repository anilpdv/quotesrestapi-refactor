package utils

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// Quote : quote struct
type Quote struct {
	Img    string `json:"img"`
	Author string `json:"author"`
	Quote  string `json:"quote"`
}

// Resp : response struct
type Resp struct {
	Ok         bool   `json:"ok"`
	Content    string `json:"content_html"`
	Page       int    `json:"page"`
	PerPage    int    `json:"per_page"`
	NumResults int    `json:"num_results"`
	TotalPages int    `json:"total_pages"`
}

// ParseURL : func(url string, query []string, page []string) string
func ParseURL(url string, query []string, page []string) string {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Print((err))
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("q", query[0])
	q.Add("page", page[0])
	req.URL.RawQuery = q.Encode()

	return req.URL.String()
}

// ParsePopularURL : func(w,r) string
func ParsePopularURL(url string, page []string, format string) string {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)

	}

	q := req.URL.Query()
	q.Add("page", page[0])
	q.Add("format", format)
	req.URL.RawQuery = q.Encode()
	fmt.Println(req.URL.String())
	return req.URL.String()
}
