package scraper

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/timhi/goodreadscraper/m/v2/src/model"
	"github.com/timhi/swiss-army-knife/src/stringutil"
)

func ScrapeBook(url string, ctx context.Context) model.Book {
	log.Printf("Scraping %s \n", url)
	book := model.Book{}
	book.BookURL = url

	var nodes []*cdp.Node

	err := chromedp.Run(ctx,
		chromedp.Navigate(book.BookURL),
		chromedp.Nodes(".BookPage", &nodes, chromedp.ByQueryAll),
	)
	log.Println("Bookpage Node")
	if err != nil {
		log.Panic(err)
	}
	var bookName string
	var rating string

	if len(nodes) != 1 {
		log.Panic("More nodes than expected")
	}
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err = chromedp.Run(timeoutCtx,
		chromedp.Click(`//div[@class="Button__container"]/button[@class="Button Button--inline Button--small"][@aria-label="Book details and editions"]`, chromedp.NodeVisible),
		chromedp.WaitVisible(`//div[@class="DescListItem"]//div[@class="TruncatedContent__text TruncatedContent__text--small"]`, chromedp.BySearch),
		chromedp.Text(".BookPageTitleSection__title", &bookName, chromedp.ByQuery, chromedp.FromNode(nodes[0])),
		chromedp.Text(".RatingStatistics__rating", &rating, chromedp.ByQuery, chromedp.FromNode(nodes[0])),
	)
	if err != nil {
		return model.Book{}
	}

	book.Title = bookName
	book.Rating = stringutil.ParseFloat64(rating)

	log.Println("Getting Book Details 🦫")
	detail, err := getBookDetails(ctx)
	if err != nil {
		log.Println("Scraping Details Failed 💥")
		detail = model.EditionDetail{}
	}

	log.Println("Details scraped successfull ✅")
	book.Details = detail

	log.Println("Getting Book Genres 🦫")
	genres, genreError := getBookGenres(ctx)
	if genreError != nil {
		log.Println("Scraping Genres Failed 💥")
		log.Println(genreError)
		genres = []string{}
	}
	log.Println("Genres scraped successfull ✅")
	book.Genres = genres

	log.Println("Getting Book Authors 🦫")
	authors, authorError := getBookAuthors(ctx, nodes[0])
	if authorError != nil {
		log.Println("Scraping Authors Failed 💥")
		log.Panic(authorError)
	}
	log.Println("Authors scraped successfull ✅")
	book.Authors = authors

	log.Println("Getting Book Cover 🦫")
	coverURL, coverError := getBookCover(ctx, nodes[0])
	if coverError != nil {
		log.Println("Scraping Cover Failed 💥")
		log.Panic(coverError)
	}
	log.Println("Cover scraped successfull ✅")
	book.CoverURL = coverURL

	return book
}

func getBookAuthors(ctx context.Context, node *cdp.Node) ([]string, error) {
	authors := []string{}
	var contributorNodes []*cdp.Node
	err := chromedp.Run(ctx,
		chromedp.Nodes(`.ContributorLink`, &contributorNodes, chromedp.ByQueryAll))
	if err != nil {
		return authors, err
	}
	var authorText string
	for _, node := range contributorNodes {
		err = chromedp.Run(ctx,
			chromedp.Text("span", &authorText, chromedp.ByQuery, chromedp.FromNode(node)))
		if err != nil {
			return authors, err
		}
		authors = append(authors, authorText)
	}

	return authors, err
}

func getBookGenres(ctx context.Context) ([]string, error) {
	genres := []string{}
	var genreNodes []*cdp.Node
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := chromedp.Run(timeoutCtx,
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
	timeoutCtx, cancel := context.WithTimeout(ctx, 4*time.Second)
	defer cancel()
	log.Println("Hanging here?")
	err := chromedp.Run(timeoutCtx,
		chromedp.Nodes(`//div[@class="EditionDetails"]//dl[@class="DescList"]//div[@class="DescListItem"]`, &dtNodes, chromedp.BySearch))
	if err != nil {
		log.Println("err!nil")
		return dtNodes, err
	}
	log.Println("return nodes?")
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

func getBookCover(ctx context.Context, node *cdp.Node) (string, error) {
	cover := ""
	var coverNode []*cdp.Node
	err := chromedp.Run(ctx,
		chromedp.Nodes(`.BookPage__bookCover`, &coverNode, chromedp.ByQueryAll))
	if err != nil || len(coverNode) != 1 {
		return cover, err
	}

	var ok bool
	var obj interface{}
	// Image is not loaded on page load instead there is a placeholder image.
	err = chromedp.Run(ctx,
		chromedp.WaitVisible(".ResponsiveImage", chromedp.BySearch),
		chromedp.Evaluate(`(function() {
        var img = document.querySelector('.ResponsiveImage');
        return new Promise(function(resolve, reject) {
            var interval = setInterval(function() {
                if (img.src !== 'https://d15be2nos83ntc.cloudfront.net/images/no-cover.png') {
                    clearInterval(interval);
                    resolve(true);
                }
            }, 50);
        });
    })()`, obj),
		chromedp.AttributeValue(".ResponsiveImage", "src", &cover, &ok, chromedp.BySearch),
	)

	if err != nil || !ok {
		return cover, err
	}
	return cover, nil
}
