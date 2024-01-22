package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)

func main() {
	// Create a new collector
	c := colly.NewCollector()

	// Set up a callback for when a visited HTML element is found
	c.OnHTML("div.quote", func(e *colly.HTMLElement) {
		// Extract information from the HTML element
		quote := e.ChildText("span.text")
		author := e.ChildText("small.author")

		// Print the extracted data
		fmt.Printf("Quote: %s\nAuthor: %s\n\n", quote, author)
	})

	// Set up a callback for when a visited link is found
	c.OnHTML("li.next a", func(e *colly.HTMLElement) {
		// Visit the next page recursively
		link := e.Attr("href")
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Start the scraping by visiting the initial URL
	err := c.Visit("http://quotes.toscrape.com")
	if err != nil {
		log.Fatal(err)
	}
}
