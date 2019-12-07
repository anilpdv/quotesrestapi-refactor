package main

import (
	"log"
	"net/http"
	"os"
	"quotesrestapi/routes"

	"github.com/gorilla/mux"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4747"
		log.Println("[-] No PORT environment variable detected. Setting to ", port)
	}
	return ":" + port
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeRoute)
	r.HandleFunc("/search", routes.SearchQuotes)
	r.HandleFunc("/popular", routes.PopularQuotes)
	r.HandleFunc("/tag/{category}", routes.QuotesWithTag)
	r.HandleFunc("/random", routes.RandomQuotes)

	port := getPort()
	log.Println("[-] Listening on...", port)
	http.ListenAndServe(port, r)
}
