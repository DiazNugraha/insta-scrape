package services

import (
	"fmt"
	"insta-scrape/app/helpers"
	"time"

	"github.com/tebeka/selenium"
)

func (s *Scrape) InstagramScrape(account string, hashTag string) {
	driver := *s.Driver
	serviceDriver := *s.Service
	USERNAME := helpers.GetEnv("USERNAME", "")
	PASSWORD := helpers.GetEnv("PASSWORD", "")

	defer serviceDriver.Stop()

	driver.Get("https://www.instagram.com/accounts/login")
	err := driver.Wait(func(driver selenium.WebDriver) (bool, error) {
		time.Sleep(5 * time.Second)
		el, err := driver.FindElement(selenium.ByName, "username")
		if el != nil {
			return true, nil
		}
		return false, err
	})
	if err != nil {
		fmt.Println(err)
	}

	err = driver.Wait(func(driver selenium.WebDriver) (bool, error) {
		username, err := driver.FindElement(selenium.ByName, "username")
		if err != nil {
			fmt.Println(err)
		}
		username.SendKeys(USERNAME)
		password, err := driver.FindElement(selenium.ByName, "password")
		if err != nil {
			fmt.Println(err)
		}
		password.SendKeys(PASSWORD)

		time.Sleep(5 * time.Second)
		return true, nil
	})
	if err != nil {
		fmt.Println(err)
	}

	submit, err := driver.FindElement(selenium.ByXPATH, "//button[@type='submit']")
	if err != nil {
		fmt.Println(err)
	}
	submit.Click()
	time.Sleep(10 * time.Second)

	button, err := driver.FindElement(selenium.ByXPATH, "//div[contains(.,'Not Now')]")
	if err != nil {
		fmt.Println(err)
	}
	button.Click()
	time.Sleep(10 * time.Second)

	fmt.Println("end")
}
