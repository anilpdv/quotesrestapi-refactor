package utils

import (
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
