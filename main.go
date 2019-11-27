package main

import (
	"net/http"
	"quotesrestapi/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", routes.SearchQuotes)
	r.HandleFunc("/popular", routes.PopularQuotes)
	r.HandleFunc("/tag/{category}", routes.QuotesWithTag)
	http.ListenAndServe(":3000", r)
}
