package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	userAges := map[string]int{
		"Alice":  10,
		"Bob":    8,
		"Claire": 6,
	}

	r := mux.NewRouter()
	r.HandleFunc("/users/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		age := userAges[name]

		fmt.Fprintf(w, "User %s is %d years old.", name, age)
	}).Methods("GET")

	http.ListenAndServe(":8080", r)

}
