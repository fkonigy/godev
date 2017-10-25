package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/fkonigy/godev/webapp/models"
	"github.com/fkonigy/godev/webapp/util"
)

// create a store for users. Will be replaced by a persisted store
// like Postgres database.
var userStore = make(map[string]models.User)

// used for generating user ids
var id int = 0

func UsersController(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/users.html"))
	err := tmpl.Execute(w, userStore)
	util.CheckErr(err, "Error in rendering users.html template.")
}

func NewUserController(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/new_user.html"))
	err := tmpl.Execute(w, nil)
	util.CheckErr(err, "Error in rendering new_user.html template.")
}

func CreateUserController(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.PostFormValue("name")
	email := r.PostFormValue("email")

	id++
	user := models.User{id, name, email}
	k := strconv.Itoa(id)
	userStore[k] = user

	http.Redirect(w, r, "/users", 302)
}

func ShowUserController(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm()
	fmt.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	tmpl := template.Must(template.ParseFiles("templates/ssshow_user.html"))
	err := tmpl.Execute(w, nil)
	util.CheckErr(err, "Error in rendering show_user.html template.")
}
