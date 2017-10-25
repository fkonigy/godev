package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home with mx")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About with mx")
}

func main() {
	mx := http.NewServeMux()

	mx.HandleFunc("/", homeHandler)
	mx.HandleFunc("/about", aboutHandler)

	log.Fatal(http.ListenAndServe(":8080", mx))
}
