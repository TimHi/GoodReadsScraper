package scraper

import "github.com/timhi/goodreadscraper/m/v2/src/model"

func ScrapeAuthor(id string) model.Author {
	author := model.Author{}
	author.AuthorURL = "https://www.goodreads.com/author/show/" + id

	return author
}
