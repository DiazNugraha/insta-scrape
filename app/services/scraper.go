package services

import (
	"fmt"
	"log"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

type Scrape struct {
	Driver  *selenium.WebDriver
	Service *selenium.Service
}

func (s *Scrape) Initial() {
	service, err := selenium.NewChromeDriverService("chromedriver", 4444)
	if err != nil {
		log.Fatal(err)
	}

	s.InitService(*service)

	// configure browser option
	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{Args: []string{
		// "--headless"
	}})

	// create a new remote client with the specified options
	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		log.Fatal(err)
	}
	s.InitSelenium(driver)
}

func (s *Scrape) InitSelenium(selenium selenium.WebDriver) {
	s.Driver = &selenium
}

func (s *Scrape) InitService(service selenium.Service) {
	s.Service = &service
}

func (s *Scrape) Scraper() {

	s.Initial()
	selDriver := *s.Driver
	serviceDriver := *s.Service

	defer serviceDriver.Stop()

	// visit the target page
	err := selDriver.Get("https://scrapingclub.com/exercise/list_infinite_scroll/")
	if err != nil {
		log.Fatal(err)
	}

	// retrieve the page raw html as string
	// and logging it
	html, err := selDriver.PageSource()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(html)

}
