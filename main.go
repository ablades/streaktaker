package main

import (
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

func main() {

	// Get Path

	if len(os.Args) != 2 {
		fmt.Println("Add Filepath/Name")
		return
	}

	// Get working directory and add it to file path
	wd, _ := os.Getwd()
	filePath := wd + fmt.Sprintf(os.Args[1])

	fmt.Print(filePath)
	// Fetch Streak
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Requesting")
	})

	c.OnHTML(".js-calendar-graph-svg", func(e *colly.HTMLElement) {
		streak, _ := os.Create(filePath)
		str, _ := e.DOM.Parent().Html()
		// Write SVG to file
		streak.WriteString(str)
	})

	c.Visit("https://github.com/users/ablades/contributions")
}
