package database

import (
	"database/sql"
	"fmt"
)

type Post struct {
	Id      int
	Title   string
	Content string
	Author  string
}

func ConnectToDB() *sql.DB {
	db, err := sql.Open("postgres", "user=postgres password=1234 dbname=atheros sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	return db
}
