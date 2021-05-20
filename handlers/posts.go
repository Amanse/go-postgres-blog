package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Amanse/sql_blog/data"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type PostHandler struct {
	l  *log.Logger
	db *sql.DB
}

type UserClaims struct {
	Authorized bool   `json:"authorized"`
	Email      string `json:"email"`
	Exp        string `json:"exp"`
}

func NewPosts(l *log.Logger) *PostHandler {
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
	return &PostHandler{l, db}
}

func (p *PostHandler) GetPosts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle get all posts")
	rw.Header().Set("Access-Control-Allow-Origin", "*")

	var posts data.Posts
	posts = data.GetAllPosts(p.db)
	err := posts.ToJson(rw)
	if err != nil {
		http.Error(rw, "Couldn't decode json from db", http.StatusInternalServerError)
	}
}

func (p *PostHandler) MakePost(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST request")

	rw.Header().Set("Access-Control-Allow-Origin", "*")

	claims := r.Context().Value(KeyProduct{}).(string)

	email := claims
	log.Println(email)

	var post data.Post

	err := post.FromJson(r.Body)
	if err != nil {
		p.l.Println("err", err)
		http.Error(rw, "Couldn't decode json", http.StatusBadRequest)
		return
	}

	err = data.MakePostDB(post, p.db, email)

	if err != nil {
		log.Fatal(err)
		http.Error(rw, "Nothing big duh", http.StatusInternalServerError)
		return
	}
}

func (p *PostHandler) UpdatePost(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	i, _ := strconv.Atoi(id)
	p.l.Println("Handle PUT request", id)

	var post data.Post

	err := post.FromJson(r.Body)

	if err != nil {
		http.Error(rw, "Unmarchln't json", http.StatusBadRequest)
		return
	}
	err = data.UpdatePostDB(i, post, p.db)
	if err != nil {
		http.Error(rw, "couldn't update post", http.StatusBadRequest)
		return
	}
}

func (p *PostHandler) DeletePost(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	i, _ := strconv.Atoi(id)
	p.l.Println("Handle delete request", i)

	err := data.DeletePostDB(i, p.db)
	if err != nil {
		http.Error(rw, "not deltee", http.StatusBadRequest)
		return
	}
}
