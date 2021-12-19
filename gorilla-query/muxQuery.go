package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func QueryHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch quey parameters ar a map
	queryParams := r.URL.Query()

	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "Got parameter id: %s\n", queryParams["id"][0])
	fmt.Fprintf(w, "Got parameter category: %s\n", queryParams["category"][0])
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/articles", QueryHandler)
	r.Queries("id", "category")

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
