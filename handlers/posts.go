package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Amanse/sql_blog/data"
	_ "github.com/go-sql-driver/mysql"
)

type PostHandler struct {
	l  *log.Logger
	db *sql.DB
}

func NewPosts(l *log.Logger) *PostHandler {
	db := openDBConnection()
	return &PostHandler{l, db}
}

func openDBConnection() *sql.DB {
	//Open connection to database
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/learning")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func (p *PostHandler) GetPosts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle get all posts")

	posts := data.GetAllPosts(p.db)
	i := 0
	for i <= len(posts)-1 {
		p.l.Println("ID: ", posts[i].ID)
		p.l.Println("Body: ", posts[i].Body)
		p.l.Println("Email: ", posts[i].Email)
		p.l.Println("-----------------------")
		i = i + 1
	}
}
