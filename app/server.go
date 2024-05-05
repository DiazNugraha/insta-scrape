package app

import (
	"insta-scrape/app/services"

	"github.com/joho/godotenv"
)

func Run() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	var scrape = services.Scrape{}
	scrape.Scraper()
}
