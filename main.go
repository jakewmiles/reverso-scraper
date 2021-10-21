package main

import (
	"encoding/json"
	"fmt"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
)

type ContextPair struct {
	SourceSentence string
	TargetSentence string
}

func scrape(phrase, sourceLanguage, targetLanguage string) string {
	c := colly.NewCollector()
	url := fmt.Sprintf("https://context.reverso.net/translation/%v-%v/%v", targetLanguage, sourceLanguage, phrase)

	extensions.RandomUserAgent(c)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "*/*")
		r.Headers.Set("Connection", "keep-alive")
	})

	sentences := []ContextPair{}

	c.OnHTML("#examples-content > .example", func(e *colly.HTMLElement) {
		source_sentence := e.ChildText(".src > span")
		target_sentence := e.ChildText(".trg > span")

		sentencePair := ContextPair{
			SourceSentence: source_sentence,
			TargetSentence: target_sentence,
		}
		sentences = append(sentences, sentencePair)
	})

	error := c.Visit(url)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println("Scraping finished")

	jsonData, err := json.MarshalIndent(sentences, "", " ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonData))
	return string(jsonData)
}

func main() {
	scrape("food", "english", "spanish")
}
