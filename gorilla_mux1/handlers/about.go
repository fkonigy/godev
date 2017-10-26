package handlers

import (
	"html/template"
	"net/http"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/about.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}
