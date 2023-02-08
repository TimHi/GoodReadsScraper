package main

import (
	"fmt"

	"github.com/timhi/goodreadscraper/m/v2/src/scraper"
)

func main() {
	fmt.Println(scraper.ScrapeBook("123"))
}
