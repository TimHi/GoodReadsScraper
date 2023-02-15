package scraper

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/timhi/goodreadscraper/m/v2/src/model"
)

func ScrapeList(ctx context.Context, id string) []model.Book {
	books := []model.Book{}
	//https://www.goodreads.com/author/list/566
	//visit the page is missing!
	//bookIDs := []string{}
	var bookURLs []string

	err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(`https://www.goodreads.com/author/list/` + id)})
	if err != nil {
		log.Fatal("Failed to navigate to list")
	}
	if canListMoreBooks(ctx) {
		log.Println("Can List More")
	} else {
		log.Println("Cant list more")
	}

	err = chromedp.Run(ctx, chromedp.Tasks{
		chromedp.WaitVisible(`.tableList`),
		chromedp.Evaluate(`Array.from(document.querySelectorAll('.tableList tr a.bookTitle')).map(el => el.href)`, &bookURLs),
	})

	if err != nil {
		log.Fatalf("Failed to extract hrefs: %v", err)
	}

	for _, bookURL := range bookURLs {
		books = append(books, ScrapeBook(bookURL, ctx))
	}

	return books
}

func canListMoreBooks(ctx context.Context) bool {
	var canListMoreBooks bool
	var text string
	timeoutCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	err := chromedp.Run(timeoutCtx, chromedp.Tasks{
		chromedp.WaitVisible(`.tableList`),
		chromedp.Text("span.next_page disabled", &text, chromedp.ByQueryAll),
	})

	fmt.Println(text)
	canListMoreBooks = text != "next"
	if err != nil {
		canListMoreBooks = false
	}
	return canListMoreBooks
}
