package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func connect() *sql.DB {
	db, err := sql.Open("postgres", "user=farid password=farid database=godb sslmode=disable")
	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}

func insert(db *sql.DB) int {
	var lastInsertId int
	err := db.QueryRow("INSERT INTO users VALUES($1, $2)", 2, "Asal").Scan(&lastInsertId)
	if err != nil {
		log.Fatal(err.Error())
	}

	return lastInsertId
}

func query1(db *sql.DB) {

	stmt, err := db.Prepare("SELECT name FROM users WHERE id = $1")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(1) // we can also use stmt.Exec
	if err != nil {
		log.Fatal(err.Error())
	}

	for rows.Next() {
		// var id int
		var name string

		err := rows.Scan(&name)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Printf("%s\n", name)
	}
}

func query2(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err.Error())
	}

	for rows.Next() {
		var id int
		var name string

		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Printf("%d, %s\n", id, name)
	}
}

func update(db *sql.DB) {
	stmt, err := db.Prepare("UPDATE users SET name = $1 WHERE id = $2")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stmt.Close()

	res, err := stmt.Exec("Farid Shahy", 1) // we can also use stmt.Query
	if err != nil {
		log.Fatal(err.Error())
	}

	if rows, _ := res.RowsAffected(); rows == 1 {
		log.Printf("Row updated successfuly.")
	} else if err != nil {
		log.Fatal(err.Error())
	}

}

func main() {

	// show date, time and line number in logs
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// first connect to database
	db := connect()
	defer db.Close()

	// then insert a new row
	insert(db)

	// prepare a statement
	// query1(db)

	// do a query
	// query2(db)

	// update a row
	update(db)

}
