package authors

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/google/uuid"
	"log"
	"strings"
)

// Author represents an author
type Author struct {
	ID   uuid.UUID
	Name string
	URL  string
}

var titleCleanup = strings.NewReplacer(
	"Libros Gratis de ", "",
	"Sobre holaebook.com |", "",
	"DMCA", "",
	"Contacto", "",
	"\n", " ",
)

var alphabet = "abcdefghijklmnopqrstuvwxyz"

// GetAuthors retrieve the list of authors to be scraped
func GetAuthors() []Author {

	authors := []Author{}
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	c.OnHTML("a", func(e *colly.HTMLElement) {
		author := Author{
			ID:   uuid.Must(uuid.NewRandom()),
			Name: strings.TrimSpace(titleCleanup.Replace(e.Text)),
			URL:  e.Attr("href"),
		}
		if author.Name != "" {
			authors = append(authors, author)
		}
	})

	c.OnScraped(func(r *colly.Response) {
		log.Println("Scrape finished")
	})

	for _, l := range alphabet {
		c.Visit(fmt.Sprintf("https://www.holaebook.com/autores/%s.html", string(l)))
	}
	return authors
}
