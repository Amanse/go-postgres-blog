package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Amanse/sql_blog/handlers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	//Make a new logger to pass into the handler
	l := log.New(os.Stdout, "posts-api", log.LstdFlags)

	//Get posts handler
	ph := handlers.NewPosts(l)

	// Make a new mux/router
	r := mux.NewRouter()

	//Make getRouter
	getRouter := r.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/posts", ph.GetPosts)

	//Post request router
	postRouter := r.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/posts", ph.MakePost)

	//Put Request
	putRouter := r.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/posts/{id:[0-9]+}", ph.UpdatePost)

	//Delete Request
	deleteRouter := r.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/posts/{id:[0-9]+}", ph.DeletePost)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":9090",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	l.Println("Running on 9090")
	log.Fatal(srv.ListenAndServe())

}
