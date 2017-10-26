package handlers

import (
	// "fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var userAges = map[string]int{
	"Alice":  10,
	"Bob":    8,
	"Claire": 6,
}

type User struct {
	Name string
	Age  int
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	age := userAges[name]

	// fmt.Fprintf(w, "User %s is %d years old.", name, age)
	tmpl := template.Must(template.ParseFiles("templates/users.html"))
	user := User{
		Name: name,
		Age:  age,
	}
	err := tmpl.Execute(w, user)
	if err != nil {
		panic(err)
	}
}
