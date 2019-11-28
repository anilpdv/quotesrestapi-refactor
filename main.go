package main

import (
	"log"
	"net/http"
	"os"
	"quotesrestapi/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", routes.SearchQuotes)
	r.HandleFunc("/popular", routes.PopularQuotes)
	r.HandleFunc("/tag/{category}", routes.QuotesWithTag)
	//Port
	getport := os.Getenv("PORT")

	if getport == "" {
		log.Fatal("$PORT must be set")
	}

	http.ListenAndServe(getport, r)
}
