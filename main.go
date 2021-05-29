package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
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
	uh := handlers.NewUser(l)

	// Make a new mux/router
	r := mux.NewRouter()

	//Make getRouter
	getRouter := r.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/posts", ph.GetPosts)
	getRouter.HandleFunc("/posts/{id}", ph.GetPost)
	//TODO
	// - Add titles to posts

	//Post request router
	postRouter := r.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/posts", ph.MakePost)
	postRouter.HandleFunc("/users/add", uh.AddUser)
	postRouter.HandleFunc("/users/login", uh.LoginUser)

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

	go func() {
		//err := srv.ListenAndServe()
		//if err != nil {
		//	log.Fatal(err)
		//}
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	l.Println("Serving")

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Shutting down", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	srv.Shutdown(tc)

}
