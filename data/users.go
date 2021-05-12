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

func LoginUser(u User, db *sql.DB) (bool, error){
	//See if user exists
	res, err := db.Query("SELECT email FROM users WHERE email=$1 LIMIT 1", u.Email)

	if err != nil {
		log.Println(err)
		return false, err
	}

	if res.Next() != true {
		return false, fmt.Errorf("User doesn't exits")
	}

	var pass string

	query := "SELECT password FROM users WHERE email=$1 LIMIT 1"
	resp, err := db.Query(query, u.Email)

	if err != nil {
		log.Println(err)
		return false, err
	}

	for resp.Next(){
		resp.Scan(&pass)
		if pass != u.Password {
			return false, fmt.Errorf("Invalid credentials")
		}
	}

	return true, nil
	

}
