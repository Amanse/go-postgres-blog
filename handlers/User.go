package handlers

import (
	"context"
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

func (u *UserHandler) LoginUser(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("Login request")
	rw.Header().Set("Access-Control-Allow-Origin", "null")
	rw.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var user data.User

	err := user.FromJson(r.Body)

	if err != nil {
		log.Println(err)
		http.Error(rw, "Can't decode json", http.StatusBadRequest)
		return
	}

	isLog, err := data.LoginUser(user, u.db)

	if err != nil {
		u.l.Println(err)
		http.Error(rw, "cant login", http.StatusInternalServerError)
		return
	}

	if isLog {
		token, err := getToken(user)
		if err != nil {
			log.Println("Token issue")
			http.Error(rw, "No can do", http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(rw, token)
		u.l.Println("Login complete")
		return
	}

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

type KeyProduct struct{}

func (u *UserHandler) IsAuth(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Access-Control-Allow-Origin", "*, null")
		rw.Header().Set("Access-Control-Allow-Header", "Token, *")
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Error")
				}
				return []byte("cringe"), nil
			})

			if err != nil {
				fmt.Fprintln(rw, err)
			}

			if token.Valid {
				if claims, ok := token.Claims.(jwt.MapClaims); ok {
					ctx := context.WithValue(r.Context(), KeyProduct{}, claims["email"])
					r = r.WithContext(ctx)
					endpoint(rw, r)
				}
			}
		} else {
			log.Println("Notoken")
		}
	})
}
