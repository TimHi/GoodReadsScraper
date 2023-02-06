package scraper

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/timhi/goodreadscraper/m/v2/src/model"
	"github.com/timhi/swiss-army-knife/src/stringutil"
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

func ScrapBook(id string) model.Book {
	c := colly.NewCollector()
	scrapedBook := model.Book{}

	c.OnHTML("div.RatingStatistics__rating", func(e *colly.HTMLElement) {
		fmt.Println("Wasd")
		fmt.Println(e.Text)
		scrapedBook.Rating = stringutil.ParseFloat64(e.Text)
	})

	c.OnHTML("div.ContributorLinksList", func(e *colly.HTMLElement) {
		e.ForEach("a.ContributorLink", func(_ int, el *colly.HTMLElement) {
			scrapedBook.Authors = append(scrapedBook.Authors, el.Text)
		})
	})

	c.Visit(BASE_URL + BOOK_ENDPOINT + id)

	return scrapedBook
}
