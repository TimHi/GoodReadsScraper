package main

import (
	"flag"

	"github.com/timhi/goodreadscraper/m/v2/src/scraper"
)

func main() {
	var bookID = flag.String("book", "18144590", "specify the book id, can be obtained from the url.")
	scraper.ScrapBook(*bookID)
}
