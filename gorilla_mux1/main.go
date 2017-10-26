package main

import (
	// "fmt"
	"net/http"

	"github.com/fkonigy/godev/gorilla_mux1/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/about", handlers.AboutHandler)
	r.HandleFunc("/users/{name}", handlers.UsersHandler).Methods("GET")

	http.ListenAndServe(":8080", r)

}
