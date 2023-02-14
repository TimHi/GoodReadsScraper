package scraper

import (
	"context"
	"fmt"
	"log"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/timhi/goodreadscraper/m/v2/src/model"
)

func ScrapeList(ctx context.Context, id string) []model.Book {
	books := []model.Book{}
	//https://www.goodreads.com/author/list/566
	//visit the page is missing!
	//bookIDs := []string{}
	var nodes []*cdp.Node
	err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(`https://www.goodreads.com/author/list/` + id),
		chromedp.Nodes(`tr`, &nodes, chromedp.ByQueryAll),
	})
	if err != nil {
		log.Fatal(err)
	}

	bookURLs := []string{}
	for _, node := range nodes {

	}

	fmt.Println(len(nodes))
	fmt.Println(nodes[0].Attributes)
	return books
}
