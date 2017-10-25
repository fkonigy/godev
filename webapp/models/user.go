package models

// import (
// 	"log"
// 	"time"
// )

type User struct {
	Id    int    //`db:"user_id"
	Name  string //`db:",size:50"` // Column size set to 50
	Email string //`db:"size:50"`
	// Created int64
}
