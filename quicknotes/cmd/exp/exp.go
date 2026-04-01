package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

var conn *pgx.Conn

type Post struct {
	id      int
	title   string
	content string
	author  string
}

func main() {
	dbURL := "postgres://postgres:rska2022@localhost:5432/postgres"
	var err error

	conn, err = pgx.Connect(context.Background(), dbURL)
	if err != nil {
		panic(err)
	}
	defer conn.Close(context.Background())

	fmt.Println("The connection to the database was successful")

	createTable()
	//insertPost()
	//insertPostWithReturn()
	//selectPostById(1)
	selectAllPosts()
}

func createTable() {

	query := `CREATE TABLE IF NOT EXISTS posts(
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		content TEXT,
		author TEXT NOT NULL
	)
	`
	_, err := conn.Exec(context.Background(), query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Table posts created")
}

func insertPost() {

	title := "Post 1"
	content := "Conteúdo do post 1"
	author := "Alexander Alves"
	query := `INSERT INTO posts(title, content, author)
		VALUES ($1, $2, $3)`

	//fmt.Println(query)
	_, err := conn.Exec(context.Background(), query, title, content, author)
	if err != nil {
		panic(err)
	}

	fmt.Println("Posts created")
}

func insertPostWithReturn() {

	title := "Post 2"
	content := "Conteúdo do post 2"
	author := "Alexander Alves"
	query := `INSERT INTO posts(title, content, author)
		VALUES ($1, $2, $3) RETURNING id;`

	//fmt.Println(query)
	row := conn.QueryRow(context.Background(), query, title, content, author)
	var id int

	if err := row.Scan(&id); err != nil {
		panic(err)
	}

	fmt.Println("Posts created, Id:", id)
}

func selectPostById(id int) {

	var title, content, author string
	query := `SELECT * FROM posts WHERE id = $1`

	//fmt.Println(query)
	row := conn.QueryRow(context.Background(), query, id)
	err := row.Scan(&title, &content, &author)

	if err == pgx.ErrNoRows {
		fmt.Println("No post found for id = ", id)
		return
	}
	if err != nil {
		panic(err)
	}

	fmt.Printf("Post - Title= %s, Content= %s, Author= %s\n", title, content, author)
}

func selectAllPosts() {

	query := `SELECT * FROM posts`
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Err() != nil {
		panic(rows.Err())
	}

	var posts []Post
	for rows.Next() {
		var post Post
		err = rows.Scan(&post.id, &post.title, &post.content, &post.author)
		if err != nil {
			panic(err)
		}
		posts = append(posts, post)
	}

	for _, post := range posts {
		fmt.Println("Post:", post)
	}
}
