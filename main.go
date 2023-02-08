package main

import (
	"fmt"
	"log"

	"github.com/timhi/goodreadscraper/m/v2/src/scraper"
)

func main() {
	log.Println("Start scraping 🦫")
	fmt.Println(scraper.ScrapeBook("123"))
}
