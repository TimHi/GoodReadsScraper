package main

import (
	"flag"
	"log"

	"github.com/timhi/goodreadscraper/m/v2/src/output"
	"github.com/timhi/goodreadscraper/m/v2/src/scraper"
)

func main() {
	var bookID = flag.String("book", "", "specify the book id, can be obtained from the url.")
	flag.Parse()
	if *bookID != "" {
		// E.g go run main.go -book 18144590
		scrapedBook := scraper.ScrapeBook(*bookID)
		log.Println("Starting headless browser, this might take a second... 🦫")

		output.AsJson(scrapedBook)
	}
}
