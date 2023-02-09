package scraper

import (
	"context"
	"log"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/timhi/goodreadscraper/m/v2/src/model"
)

func ScrapeAuthor(id string, ctx context.Context) model.Author {
	author := model.Author{}
	author.AuthorURL = "https://www.goodreads.com/author/show/" + id
	var authorNode []*cdp.Node

	err := chromedp.Run(ctx,
		chromedp.Navigate(author.AuthorURL),
		chromedp.Nodes(".mainContent", &authorNode, chromedp.ByQueryAll),
	)

	if err != nil || len(authorNode) != 1 {
		log.Panic(err)
	}

	return author
}
