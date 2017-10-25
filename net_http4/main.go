package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Easy Home.")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Easy About.")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
