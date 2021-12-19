package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jxskiss/base62"
)

type payload struct {
	Url string `json:"url"`
}

// Map pretends to be a database that maps from shortUrl to longUrl
var db = make(map[string]string)

func PostUrl(w http.ResponseWriter, r *http.Request) {
	// Decode the url in the json payload
	var p payload
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Add the shorten url to the db
	w.WriteHeader(http.StatusOK)
	shortUrl := base62.EncodeToString([]byte(p.Url))
	db[shortUrl] = p.Url

	// Return all entries in the db
	fmt.Fprintln(w, db)
}

func GetShortUrl(w http.ResponseWriter, r *http.Request) {
	// Get the request params as a map
	params := mux.Vars(r)

	if longUrl, ok := db[params["url"]]; ok {
		// If the short url is found, return original url
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, longUrl)
	} else {
		// Short url not found
		http.NotFound(w, r)
	}
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/new", PostUrl).Methods("POST")
	r.HandleFunc("/api/v1/{url}", GetShortUrl).Methods("GET")

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
