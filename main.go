package main

import (
	"encoding/json"
	"fmt"

	"github.com/gocolly/colly/v2"
)

type ContextPair struct {
	Symbol           string
	MarketCap        string
	Price            string
	Volume24H        string
	PercentChange1H  string
	PercentChange24H string
	PercentChange7D  string
}

func scrape(phrase, srcLang, targLang string) string {
	c := colly.NewCollector()

	c.OnHTML("", func(e *colly.HTMLElement) {

	})

	url := fmt.Sprintf("https://context.reverso.net/translation/%v-%v/%v", targLang, srcLang, phrase)

	c.Visit(url)

	fmt.Println("Scraping finished")

	jsonData, err := json.MarshalIndent(url, "", " ")
	if err != nil {
		panic(err)
	}

	return string(jsonData)
}

func main() {
	scrape("hello", "english", "french")
}
