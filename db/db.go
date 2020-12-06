package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/rubenwap/book-db/authors"
	"github.com/rubenwap/book-db/books"
	"github.com/rubenwap/book-db/config"
)

func GetDBConnection(connStr string) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Print(err)
		log.Print("Could not connect to Postgres.")
		return &pgx.Conn{}
	}
	return conn
}

func AddAuthor(conn *pgx.Conn, author authors.Author) error {
	_, err := conn.Exec(context.Background(), "insert into authors(id, name, url) values($1, $2, $3);", author.ID, author.Name, author.URL)
	return err
}

func AddBook(conn *pgx.Conn, book books.Book, author authors.Author) error {
	_, err := conn.Exec(context.Background(), "insert into books(id, url, author_id) values($1, $2, $3)", book.ID, book.URL, author.ID)
	return err
}

func BuildDBURL() string {
	conf := config.Config()
	host, _ := conf.Get("db.host")
	port, _ := conf.Get("db.port")
	name, _ := conf.Get("db.name")
	user, _ := conf.Get("db.user")
	pass, _ := conf.Get("db.pass")
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, pass, host, port, name)
}
