package main

import (
	"log"
	"net/http"

	"github.com/fkonigy/godev/net_http3/handlers"
)

func main() {
	hh := http.HandlerFunc(handlers.HomeHandler)
	http.Handle("/", hh)

	ha := http.HandlerFunc(handlers.AboutHandler)
	http.Handle("/about", ha)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
