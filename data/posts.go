package data

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type Post struct {
	ID     int    `json:"id"`
	Body   string `json:"body"`
	Email  string `json:"email"`
	UserId int    `json:"user_id"`
}

type Posts []Post

var postList Posts

func (p *Posts) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Post) FromJson(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func GetAllPosts(db *sql.DB) Posts {

	postList = make(Posts, 0)

	query := "SELECT * FROM posts"
	res, err := db.Query(query)
	defer res.Close()

	if err != nil {
		log.Println("error: ", err)
	}

	for res.Next() {
		var post Post
		err := res.Scan(&post.ID, &post.Body, &post.Email, &post.UserId)

		if err != nil {
			log.Println(err)
		}

		postList = append(postList, post)

	}

	return postList
}

var PostNotMade = fmt.Errorf("Couldn't make post")

func MakePostDB(p Post, db *sql.DB) error {

	// getting the email

	equery := "SELECT email FROM users WHERE id=$1"
	res, err := db.Query(equery, p.UserId)

	if err != nil {
		log.Println(err)
		return PostNotMade
	}

	for res.Next() {
		var email string
		err = res.Scan(&email)
		if err != nil {
			log.Println(err)
			return PostNotMade
		}
		p.Email = email
	}

	query := "INSERT INTO posts(body, email, user_id) VALUES($1,$2, $3)"
	_, err = db.Exec(query, p.Body, p.Email, p.UserId)

	if err != nil {
		log.Println(err)
		return PostNotMade
	}

	return nil

}

func UpdatePostDB(id int, p Post, db *sql.DB) error {

	query := "UPDATE posts SET body=$1 WHERE id=$2"
	_, err := db.Exec(query, p.Body, id)

	if err != nil {
		log.Println(err)
		return PostNotMade
	}

	return nil
}

var CantDeletePost = fmt.Errorf("cant delete post")

func DeletePostDB(id int, db *sql.DB) error {
	query := "DELETE FROM posts WHERE id=$1"
	_, err := db.Exec(query, id)

	if err != nil {
		log.Println(err)
		return CantDeletePost
	}

	return nil
}
