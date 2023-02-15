package main

import (
	"context"
	"flag"
	"log"

	"github.com/chromedp/chromedp"
	"github.com/timhi/goodreadscraper/m/v2/src/output"
	"github.com/timhi/goodreadscraper/m/v2/src/scraper"
)

func main() {
	var bookID = flag.String("book", "", "specify the book id, can be obtained from the url.")
	var authorID = flag.String("author", "", "specify the author id, can be obtained from the url.")
	flag.Parse()

	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	log.Println("Starting headless browser, this might take a second... 🦫")
	if *bookID != "" {
		// E.g go run main.go -book 18144590
		bookURL := "https://www.goodreads.com/book/show/" + *bookID
		scrapedBook := scraper.ScrapeBook(bookURL, ctx)
		output.AsJson(scrapedBook)
	}

	if *authorID != "" {
		// E.g go run main.go -author 3670
		scrapedAuthor := scraper.ScrapeAuthor(*authorID, ctx)
		output.AsJson(scrapedAuthor)
	}
}
