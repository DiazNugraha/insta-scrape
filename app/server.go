package app

import (
	"insta-scrape/app/services"
)

func Run() {
	var scrape = services.Scrape{}
	scrape.Scraper()
}
