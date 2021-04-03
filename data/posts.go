package data

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	ID    int    `json:"id"`
	Body  string `json:"body"`
	Email string `json:"email"`
}

type Posts []Post

var postList Posts

func (p *Posts) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetAllPosts(db *sql.DB) Posts {
	query := "SELECT * FROM posts"
	res, err := db.Query(query)
	defer res.Close()

	if err != nil {
		log.Println("error: ", err)
	}

	for res.Next() {
		var post Post
		err := res.Scan(&post.ID, &post.Body, &post.Email)

		if err != nil {
			log.Fatal(err)
		}

		postList = append(postList, post)

	}

	return postList
}
