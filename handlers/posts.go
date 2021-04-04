package handlers

import (
	"log"
	"net/http"

	"github.com/Amanse/sql_blog/data"
	_ "github.com/go-sql-driver/mysql"
)

type PostHandler struct {
	l *log.Logger
}

func NewPosts(l *log.Logger) *PostHandler {
	return &PostHandler{l}
}

func (p *PostHandler) GetPosts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle get all posts")

	var posts data.Posts
	posts = data.GetAllPosts()
	err := posts.ToJson(rw)
	if err != nil {
		http.Error(rw, "Couldn't decode json from db", http.StatusInternalServerError)
	}
}
