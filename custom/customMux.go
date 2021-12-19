package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

type CustomServeMux struct {
}

// Any struct that has the ServeHTTP function can be a multiplexer
// To add routes, we must add more else if clauses
func (p *CustomServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/hello" {
		hello(w, r)
		return
	} else if r.URL.Path == "/randomInt" {
		giveRandomInt(w, r)
		return
	} else if r.URL.Path == "/randomFloat" {
		giveRandomFloat(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func giveRandomInt(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, rand.Int())
}

func giveRandomFloat(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Your random number is: %f\n", rand.Float64())
}

func main() {
	mux := &CustomServeMux{}
	log.Fatal(http.ListenAndServe(":8000", mux))
}
