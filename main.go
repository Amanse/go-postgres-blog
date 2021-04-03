package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/learning")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB connected")

	state := "INSERT INTO posts(body, email) VALUES('verysad', 'everyone@earth.com')"
	res, err := db.Query(state)
	defer res.Close()

	if err != nil {
		log.Fatal(err)
	}

}
