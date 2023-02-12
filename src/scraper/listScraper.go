package scraper

import (
	"context"
	"fmt"
	"log"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/timhi/goodreadscraper/m/v2/src/model"
)

func ScrapeList(ctx context.Context) []model.Book {
	books := []model.Book{}
	//https://www.goodreads.com/author/list/566
	//visit the page is missing!
	//bookIDs := []string{}
	var nodes []*cdp.Node

	//tableList
	err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Nodes(`table`, &nodes, chromedp.ByQueryAll),
	})
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(len(nodes))
	fmt.Println(nodes[0].Attributes)
	return books
}
