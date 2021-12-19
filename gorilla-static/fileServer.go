package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))

	log.Fatal(http.ListenAndServe(":8000", r))
}
