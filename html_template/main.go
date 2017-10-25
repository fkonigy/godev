package main

import (
	"html/template"
	"net/http"
)

type Todo struct {
	Task string
	Done bool
}

func main() {

	tmpl := template.Must(template.ParseFiles("temp1.html"))
	todos := []Todo{
		{"Learning Go", true},
		{"Go Web Development", false},
		{"Go Concurrency", false},
		{"Cryptography", true},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, struct{ Todos []Todo }{todos}) // what the fuck!!!
	})

	http.ListenAndServe(":8080", nil)

}
