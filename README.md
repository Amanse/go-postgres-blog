# go-mariadb-blog
Just a backend(for now) written in golang and using mariadb as database

## setup
- git clone
- go mod tidy
- start postgresql terminal
- CREATE DATABASE learning;
- USE learning;
- CREATE TABLE posts (id int PRIMARY KEY AUTO_INCREMENT, body text, email VARCHAR(255));
- exit
- go to root directory
- go run main.go

## api routes
- :9090/posts [GET] to get all posts
- :9090/posts [POST] to make post
- :9090/posts/{id} [PUT] to update post
- :9090/posts/{id} [DELETE] to delete post
