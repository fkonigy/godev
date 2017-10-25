package controllers

import (
	"html/template"
	"net/http"

	"github.com/fkonigy/godev/webapp/util"
)

func AboutController(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/about.html"))
	err := tmpl.Execute(w, nil)
	util.CheckErr(err, "Error in rendering about.html template.")
}
