package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
)

type ContextPair struct {
	SourceSentence string
	TargetSentence string
}

func contains(s []string, i string) bool {
	for _, v := range s {
		if v == i {
			return true
		}
	}
	return false
}

func scrape(phrase, sourceLanguage, targetLanguage string) (string, error) {
	sourceLanguage = strings.ToLower(sourceLanguage)
	targetLanguage = strings.ToLower(targetLanguage)
	possibleLanguages := [15]string{"arabic", "german", "english", "spanish", "french", "hebrew", "italian", "japanese", "dutch", "polish", "portuguese", "romanian", "russian", "turkish", "chinese"}

	if !contains(possibleLanguages[:], sourceLanguage) || !contains(possibleLanguages[:], targetLanguage) {
		return "", errors.New("one or more of your language choices is unavailable")
	}

	c := colly.NewCollector()
	url := fmt.Sprintf("https://context.reverso.net/translation/%v-%v/%v", sourceLanguage, targetLanguage, phrase)

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
	return string(jsonData), nil
}

func main() {
	scrape("peak", "ENGLISH", "FreNCh")
}
