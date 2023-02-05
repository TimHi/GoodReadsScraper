package scraper

import (
	"fmt"

	"github.com/gocolly/colly"
)

var BASE_URL = "https://www.goodreads.com/"
var AUTHOR_ENDPOINT = "author/show/"
var BOOK_ENDPOINT = "book/show/"

func ScrapeAuthor(id string) {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(BASE_URL + AUTHOR_ENDPOINT + id)
}

func ScrapBook(id string) {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("div.BookPageMetadataSection__ratingStats", func(e *colly.HTMLElement) {
		fmt.Println("FUCK")
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(BASE_URL + BOOK_ENDPOINT + id)
}
