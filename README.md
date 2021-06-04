# go-postgresql-blog
Just a backend(for now) written in golang and using mariadb as database

## setup
- git clone
- go mod tidy
- start postgresql terminal
- CREATE DATABASE go_blog;
- USE go_blog;
- CREATE TABLE posts (id SERIAL PRIMARY KEY, body text, email VARCHAR(255), user_id int);
- CREATE TABLE users (id SERIAL PRIMARY KEY, email VARCHAR(255) UNIQUE, password VARCHAR(200));
- exit
- go to root directory
- go run main.go

## api routes
- :9090/posts [GET] to get all posts
- :9090/posts [POST] to make post
- :9090/posts/{id} [PUT] to update post
- :9090/posts/{id} [DELETE] to delete post
