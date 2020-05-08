package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Requesting")
	})

	c.OnHTML("svg", func(e *colly.HTMLElement) {
		fmt.Print(e.Response.Save("test2.svg"))
	})

	c.Visit("https://github.com/users/ablades/contributions")

}
