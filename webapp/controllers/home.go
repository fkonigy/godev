package controllers

import (
	"html/template"
	"net/http"

	// "github.com/fkonigy/godev/webapp/models"
	"github.com/fkonigy/godev/webapp/util"
)

func HomeController(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	err := tmpl.Execute(w, nil)
	util.CheckErr(err, "Error in rendering index.html template.")
}
