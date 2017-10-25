package main

import (
	"os"
	"text/template"
)

func main() {
	data := struct {
		Name string
	}{
		Name: "Ashkan",
	}

	tmpl, err := template.ParseFiles("templates/main.txt")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
