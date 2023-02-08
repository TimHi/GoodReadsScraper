package scraper

import (
	"context"
	"fmt"
	"log"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/timhi/goodreadscraper/m/v2/src/model"
)

func ScrapeBook(id string) model.Book {
	book := model.Book{}

	// initializing a chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// navigate to the target web page and select the HTML elements of interest
	var nodes []*cdp.Node

	err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.goodreads.com/book/show/18144590"),
		chromedp.Nodes(".BookPage", &nodes, chromedp.ByQueryAll),
	)
	if err != nil {
		fmt.Println(err)
		chromedp.Cancel(ctx)
	}
	var bookName string

	err = chromedp.Run(ctx,
		chromedp.Click(`//div[@class="Button__container"]/button[@class="Button Button--inline Button--small"][@aria-label="Book details and editions"]`, chromedp.NodeVisible),
		chromedp.WaitVisible(`//div[@class="DescListItem"]//div[@class="TruncatedContent__text TruncatedContent__text--small"]`, chromedp.BySearch),
		chromedp.Text(".BookPageTitleSection__title", &bookName, chromedp.ByQuery, chromedp.FromNode(nodes[0])),
	)
	if err != nil {
		fmt.Println(err)
		chromedp.Cancel(ctx)
	}

	book.Title = bookName

	log.Println("Getting Book Details 🦫")
	detail, err := getBookDetails(ctx)
	if err != nil {
		log.Println("Scraping Details Failed 💥")
		log.Panic(err)
	}

	log.Println("Details scraped successfull ✅")
	book.Details = detail

	log.Println("Getting Book Genres 🦫")
	genres, genreError := getBookGenres(ctx)
	if genreError != nil {
		log.Println("Scraping Genres Failed 💥")
		log.Panic(err)
	}
	log.Println("Genres scraped successfull ✅")
	book.Genres = genres

	return book
}

func getBookGenres(ctx context.Context) ([]string, error) {
	genres := []string{}
	var genreNodes []*cdp.Node

	err := chromedp.Run(ctx,
		chromedp.Nodes(`//div[@class="BookPageMetadataSection__genres"]//ul[@class="CollapsableList"]//span//span[@class="BookPageMetadataSection__genreButton"]`, &genreNodes, chromedp.BySearch))

	if err != nil {
		return nil, err
	}

	for _, node := range genreNodes {
		var text string
		chromedp.Run(ctx,
			chromedp.Text("a", &text, chromedp.ByQuery, chromedp.FromNode(node)))
		genres = append(genres, text)
	}

	if err != nil {
		return nil, err
	}

	return genres, nil
}

func getBookDetails(ctx context.Context) (model.EditionDetail, error) {
	dtNodes, err := getEditionDetailListNodes(ctx)
	if err != nil {
		return model.EditionDetail{}, err
	}
	return getDetailValues(ctx, dtNodes)
}

func getEditionDetailListNodes(ctx context.Context) ([]*cdp.Node, error) {
	var dtNodes []*cdp.Node
	err := chromedp.Run(ctx,
		chromedp.Nodes(`//div[@class="EditionDetails"]//dl[@class="DescList"]//div[@class="DescListItem"]`, &dtNodes, chromedp.BySearch))
	if err != nil {
		return dtNodes, err
	}
	return dtNodes, nil
}

func getDetailValues(ctx context.Context, dtNodes []*cdp.Node) (model.EditionDetail, error) {
	detail := model.EditionDetail{}
	var key, value string
	for _, node := range dtNodes {
		err := chromedp.Run(ctx,
			chromedp.Text("dt", &key, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.Text("dd", &value, chromedp.ByQuery, chromedp.FromNode(node)),
		)
		if err != nil {
			return model.EditionDetail{}, err
		}
		switch key {
		case "Format":
			detail.Format = value
		case "Published":
			detail.Published = value
		case "ISBN":
			detail.ISBN = value
		case "Language":
			detail.Language = value
		}
	}

	return detail, nil
}
