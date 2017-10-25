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

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	myTasks := []Todo{
		{"Learning Go", true},
		{"Go Web Development", false},
		{"Go Concurrency", false},
		{"Cryptography", true},
	}

	data := struct{ Tasks []Todo }{myTasks} // It must be clear now.

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":8080", nil)

}
