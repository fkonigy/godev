package handlers

import (
	// "fmt"
	"html/template"
	"net/http"
	// "github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}
