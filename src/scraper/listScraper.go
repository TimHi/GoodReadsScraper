package scraper

import (
	"context"
	"log"

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
		chromedp.Navigate(`https://www.goodreads.com/author/list/` + id),
		chromedp.WaitVisible(`.tableList`),
		chromedp.Evaluate(`Array.from(document.querySelectorAll('.tableList tr a.bookTitle')).map(el => el.href)`, &bookURLs),
	})

	if err != nil {
		log.Fatalf("Failed to extract hrefs: %v", err)
	}

	return books
}
