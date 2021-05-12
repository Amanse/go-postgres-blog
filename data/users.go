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
	//See if user exists
	res, err := db.Query("SELECT email FROM users WHERE email=$1 LIMIT 1", u.Email)

	if res.Next() {
		return fmt.Errorf("user alredy exists")
	}

	query := "INSERT INTO users(email, password) VALUES($1, $2)"
	_, err = db.Exec(query, u.Email, u.Password)

	if err != nil {
		log.Println(err)
		return UserNotAdded
	}

	return nil
}
