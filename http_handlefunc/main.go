package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	log.SetFlags(log.Lshortfile)
	// getGoogle()

	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/users/", age)

	err := http.ListenAndServe(":8080", nil) // no handler need yet
	checkErr(err, "Error in starting server.")

}

func checkErr(err error, msg string) {
	if err != nil {
		log.Println(msg, err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	out := []byte("Index")
	w.Write(out)
}

func about(w http.ResponseWriter, r *http.Request) {
	out := []byte("About")
	w.Write(out)
}

func age(w http.ResponseWriter, r *http.Request) {

	userAges := map[string]int{
		"Asal":   6,
		"Ashkan": 3,
	}

	name := r.URL.Path[len("/users/"):]
	age := userAges[name]

	fmt.Fprintf(w, "%s is %d years old.", name, age)
}

// func getGoogle() {
// 	resp, err := http.Get("https://www.google.com")
// 	checkErr(err, "Error in internet connection.")
// 	defer resp.Body.Close()
// 	log.Println(resp.Header["Server"])
// }
