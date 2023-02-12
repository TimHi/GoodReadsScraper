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

	err = chromedp.Run(ctx,
		chromedp.Text(".authorName", &author.Name, chromedp.ByQuery, chromedp.FromNode(authorNode[0])),
		chromedp.Text(".aboutAuthorInfo", &author.Description, chromedp.ByQuery, chromedp.FromNode(authorNode[0])),
	)
	if err != nil {
		log.Panic(err)
	}
	log.Println("Getting authorLeftContainer")
	var imageNode []*cdp.Node
	var ok bool
	err = chromedp.Run(ctx,
		chromedp.Nodes(".authorLeftContainer", &imageNode, chromedp.ByQueryAll),
	)

	if err != nil || len(imageNode) != 1 {
		log.Panic(err)
	}

	log.Println("Getting Image")
	var authorImgAlt string
	err = chromedp.Run(ctx, chromedp.Tasks{
		chromedp.WaitVisible(`//img[@itemprop='image']`, chromedp.BySearch),
		chromedp.AttributeValue(`//img[@itemprop='image']`, "alt", &authorImgAlt, &ok),
		chromedp.AttributeValue(`//img[@itemprop='image']`, "src", &author.PhotoURL, &ok),
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(authorImgAlt)
	author.Books = ScrapeList(ctx)
	return author
}
