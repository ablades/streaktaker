package main

import (
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Requesting")
	})

	c.OnHTML(".js-calendar-graph-svg", func(e *colly.HTMLElement) {
		streak, _ := os.Create("streak.svg")
		str, _ := e.DOM.Parent().Html()
		streak.WriteString(str)
	})

	c.Visit("https://github.com/users/ablades/contributions")
}
