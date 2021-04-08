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
	ID    int    `json:"id"`
	Body  string `json:"body"`
	Email string `json:"email"`
}

type Posts []Post

var postList Posts

func openDBConnection() *sql.DB {
	//Open connection to database
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=learning sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to db")
	return db
}

func (p *Posts) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Post) FromJson(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func GetAllPosts() Posts {

	postList = make(Posts, 0)

	db := openDBConnection()
	defer db.Close()
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

var PostNotMade = fmt.Errorf("Couldn't make post")

func MakePostDB(p Post) error {
	db := openDBConnection()
	defer db.Close()

	query := "INSERT INTO posts(body, email) VALUES($1,$2)"
	_, err := db.Exec(query, p.Body, p.Email)

	if err != nil {
		log.Fatal(err)
		return PostNotMade
	}

	return nil

}

func UpdatePostDB(id int, p Post) error {
	db := openDBConnection()
	defer db.Close()

	query := "UPDATE posts SET body=$1 WHERE id=$2"
	_, err := db.Exec(query, p.Body, id)

	if err != nil {
		log.Fatal(err)
		return PostNotMade
	}

	return nil
}

var CantDeletePost = fmt.Errorf("cant delete post")

func DeletePostDB(id int) error {
	db := openDBConnection()
	defer db.Close()

	query := "DELETE FROM posts WHERE id=$1"
	_, err := db.Exec(query, id)

	if err != nil {
		log.Fatal(err)
		return CantDeletePost
	}

	return nil
}
