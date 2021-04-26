package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/Amanse/sql_blog/data"
)

type UserHandler struct {
	l  *log.Logger
	db *sql.DB
}

func NewUser(l *log.Logger) *UserHandler {

	//Open connection to database

	var conn string

	if os.Getenv("DB_CONN") != "" {
		log.Println(os.Getenv("DB_CONN"))
		conn = os.Getenv("DB_CONN")
	} else {
		conn = "host=localhost user=me dbname=go_blog sslmode=disable"
	}

	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}
	return &UserHandler{l, db}
}

func (u *UserHandler) AddUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	u.l.Println("Add user handler")

	var user data.User

	err := user.FromJson(r.Body)

	if err != nil {
		u.l.Println("err", err)
		http.Error(rw, "Couldnt decode json", http.StatusBadRequest)
		return
	}

	err = data.AddUser(user, u.db)

	if err != nil {
		log.Fatal(err)
		http.Error(rw, "Cant add user", http.StatusInternalServerError)
		return
	}

}
