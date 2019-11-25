package main

import (
	"net/http"
	"quotesrestapi/routes"
)

func main() {
	http.HandleFunc("/", routes.SearchQuotes)

	http.ListenAndServe(":3000", nil)
}
