package scraper

import (
	"context"
	"fmt"
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

	var imageNode []*cdp.Node
	var ok bool
	err = chromedp.Run(ctx,
		chromedp.Nodes(".authorLeftContainer", &imageNode, chromedp.ByQueryAll),
	)

	if err != nil || len(imageNode) != 1 {
		log.Panic(err)
	}

	err = chromedp.Run(ctx,
		chromedp.AttributeValue("img", "src", &author.PhotoURL, &ok, chromedp.FromNode(imageNode[0])),
	)
	if err != nil || !ok {
		//log.Panic(err)
	}

	author.Books = getBooks(ctx)
	return author
}

func getBooks(ctx context.Context) []model.Book {
	books := []model.Book{}
	var bookNodes []*cdp.Node
	err := chromedp.Run(ctx,
		chromedp.Nodes(".stacked tableList", &bookNodes, chromedp.ByQueryAll), // get table nodes
	)
	fmt.Println(len(bookNodes))
	if err != nil || len(bookNodes) == 0 {
		log.Panic(err)
	}
	var ok bool
	for _, node := range bookNodes {
		var bookId string
		err = chromedp.Run(ctx,
			chromedp.AttributeValue("a", "href", &bookId, &ok, chromedp.FromNode(node)))
		fmt.Println(bookId)
		books = append(books, ScrapeBook(bookId, ctx))
	}

	return books
}
