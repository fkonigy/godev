package main

import (
	"log"
	"net/http"

	"github.com/fkonigy/godev/webapp/controllers"
)

func main() {
	mx := http.NewServeMux()

	mx.HandleFunc("/", controllers.HomeController)
	mx.HandleFunc("/users", controllers.UsersController)
	mx.HandleFunc("/users/new", controllers.NewUserController)
	mx.HandleFunc("/users/create", controllers.CreateUserController)
	// this form of variables are for gorilla/mux
	// golanf mux has no solution for this and I
	// must implement a solution forthin.
	mx.HandleFunc("/users/show/{id}", controllers.ShowUserController)
	mx.HandleFunc("/about", controllers.AboutController)

	log.Fatal(http.ListenAndServe(":8080", mx))
}
