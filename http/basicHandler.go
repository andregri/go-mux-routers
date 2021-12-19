package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func MyServer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hello world!\n")
}

// We use the multiplexer provided in the http package
func main() {
	newMux := http.NewServeMux()

	newMux.HandleFunc("/hello", MyServer)

	newMux.HandleFunc("/randomFloat", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, rand.Float64())
	})

	newMux.HandleFunc("/randomInt", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, rand.Int())
	})

	log.Fatal(http.ListenAndServe(":8000", newMux))
}
