package data

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) FromJson(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(u)
}

var UserNotAdded = fmt.Errorf("Couldn add usr")

func AddUser(u User, db *sql.DB) error {
	query := "INSERT INTO users(email, password) VALUES($1, $2)"
	_, err := db.Exec(query, u.Email, u.Password)

	if err != nil {
		log.Fatal(err)
		return UserNotAdded
	}

	return nil
}
