package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Amanse/sql_blog/data"
	jwt "github.com/dgrijalva/jwt-go"
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
		log.Println(err)
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
		log.Println(err)
		http.Error(rw, "Cant add user", http.StatusInternalServerError)
		return
	}

	token, err := getToken(user)
	if err != nil {
		log.Println(err)
		http.Error(rw, "Can't get token", http.StatusInternalServerError)
		return 
	}

	log.Println(token)
	fmt.Fprintln(rw, token)

}

func getToken(user data.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenStr, err := token.SignedString([]byte("cringe"))

	if err != nil {
		fmt.Println("Something went wrong", err)
		return "", err
	}

	return tokenStr, nil

}
