package main

import (

	"github.com/rubenwap/book-db/authors"
	"github.com/rubenwap/book-db/books"
	"github.com/rubenwap/book-db/db"
)

func main() {
	a := authors.GetAuthors()
	
	connStr := db.BuildDBURL()
	conn := db.GetDBConnection(connStr)
	
	for _, author := range a {
		db.AddAuthor(conn, author)
		b := books.GetBooks(author)
		for _, book := range b {
			db.AddBook(conn, book, author)
		}
	}
}
