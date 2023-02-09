package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/timhi/goodreadscraper/m/v2/src/scraper"
)

func main() {
	log.Println("Starting headless browser, this might take a second... 🦫")
	result := scraper.ScrapeBook("123")
	res2B, _ := json.Marshal(result)
	fmt.Println(string(res2B))

}
