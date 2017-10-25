package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
	"gopkg.in/gorp.v2"
)

type Post struct {
	// db tag lets you specify the column name if it differs from the struct field
	Id      int64 `db:"post_id"`
	Created int64
	Title   string `db:",size:50"`               // Column size set to 50
	Body    string `db:"article_body,size:1024"` // Set both column name and size
}

func main() {

	// err object used by checkErr func
	var err error

	// config log
	log.SetFlags(log.Lshortfile)

	//init database
	dbmap := initDB()
	defer dbmap.Db.Close()

	// delete any existing rows
	err = dbmap.TruncateTables()
	checkErr(err, "dbmap.TruncateTables() failed.")

	// define new posts
	p1 := Post{
		Created: time.Now().UnixNano(),
		Title:   "My First Post.",
		Body:    "Some awesome text here.1",
	}

	p2 := Post{
		Created: time.Now().UnixNano(),
		Title:   "My Second Post.",
		Body:    "Some awesome text here.2",
	}

	// insert new post
	err = dbmap.Insert(&p1, &p2)
	checkErr(err, "dbmap.Insert(&p1) failed.")

	// SelectInt
	count, err := dbmap.SelectInt("SELECT COUNT(*) FROM posts")
	checkErr(err, "dbmap.SelectInt() failed.")
	log.Println("Total Rows: ", count)

	// update record
	p2.Title = "This post is updated now."
	count, err = dbmap.Update(&p2)
	checkErr(err, "dbmap.Updated() failed.")
	log.Println("Number of rows updated: ", count)

	//SelectOne
	var temp Post
	err = dbmap.SelectOne(&temp, "SELECT * FROM posts WHERE post_id = $1", p1.Id)
	checkErr(err, "dbmap.SelectOne() failed.")
	log.Println("You selected this Id: ", temp.Id)
	// now update its title
	temp.Title = "Updated by selecting."
	count, err = dbmap.Update(&temp)
	checkErr(err, "dbmap.Update() failed.")

	// delete a row
	// count, err = dbmap.Delete(&temp)
	// checkErr(err, "dbmap.Delete() failed.")
	// log.Println("Number of rows deleted: ", count)

	// select multiple row
	var posts []Post
	_, err = dbmap.Select(&posts, "SELECT * FROM posts")
	checkErr(err, "dbmap.Select() failed.")
	for _, p := range posts {
		log.Printf("Id: %d - Title: %s", p.Id, p.Title)
	}

}

func initDB() *gorp.DbMap {
	db, err := sql.Open("postgres", "user=farid password=farid database=gorp sslmode=disable")
	checkErr(err, "sql.Open failed.")

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	// define tables
	dbmap.AddTableWithName(Post{}, "posts").SetKeys(true, "Id")

	// create tables
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "dbmap.CreateTablesIfNotExists() failed.")

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
