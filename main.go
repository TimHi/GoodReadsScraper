package main

import (
	"flag"
	"log"

	"github.com/timhi/goodreadscraper/m/v2/src/output"
	"github.com/timhi/goodreadscraper/m/v2/src/scraper"
)

func main() {
	var bookID = flag.String("book", "", "specify the book id, can be obtained from the url.")
	var authorID = flag.String("author", "", "specify the author id, can be obtained from the url.")
	flag.Parse()

	log.Println("Starting headless browser, this might take a second... 🦫")
	if *bookID != "" {
		// E.g go run main.go -book 18144590
		scrapedBook := scraper.ScrapeBook(*bookID)
		output.AsJson(scrapedBook)
	}

	if *authorID != "" {
		// E.g go run main.go -author 3670
		scrapedAuthor := scraper.ScrapeAuthor(*authorID)
		output.AsJson(scrapedAuthor)
	}
}
