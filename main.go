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
	r.HandleFunc("/search", routes.SearchQuotes)
	r.HandleFunc("/popular", routes.PopularQuotes)
	r.HandleFunc("/tag/{category}", routes.QuotesWithTag)
	port := getPort()
	log.Println("[-] Listening on...", port)
	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
