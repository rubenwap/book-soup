package books

import (
	"log"
	"strings"

	"github.com/gocolly/colly"
	"github.com/google/uuid"
	"github.com/rubenwap/book-db/authors"
)

// Book represents a book
type Book struct {
	ID uuid.UUID
	URL string
}

// GetBooks retrieves the url for all books
func GetBooks(author authors.Author) []Book {

	c := colly.NewCollector()
	books := []Book{}
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	c.OnHTML("a", func(e *colly.HTMLElement) {
		book := Book{
			ID:   uuid.Must(uuid.NewRandom()),
			URL: e.Attr("href"),
		}

		if strings.HasPrefix(book.URL, "book/") {
			books = append(books, book)
		}
	})

	c.OnScraped(func(r *colly.Response) {
		log.Println("Scrape finished")
	})

	
		c.Visit("https://www.holaebook.com" + author.URL)
	
	return books
}
