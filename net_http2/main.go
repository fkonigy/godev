package main

import (
	"log"
	"net/http"

	"github.com/fkonigy/godev/net_http2/handlers"
)

func main() {
	http.Handle("/", handlers.HomeHandler{})
	http.Handle("/about", handlers.AboutHandler{})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
