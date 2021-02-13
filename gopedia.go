package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	forbidden := []string{".mw-parser-output"}
	lang := flag.String("l", "en", "ISO-639-1 Language code.")
	paragraphs := flag.Int("p", 3, "Max number of paragraphs.")
	help := flag.Bool("h", false, "Print help")
	flag.Parse()
	query := strings.Join(flag.Args(), "_")
	if *help {
		flag.PrintDefaults()
	}
	c := colly.NewCollector()

	url := fmt.Sprintf("https://%s.wikipedia.org/wiki/%s", *lang, query)
	var times int = 1
	c.OnHTML(".mw-parser-output", func(e *colly.HTMLElement) {
		e.ForEach("p", func(i int, elem *colly.HTMLElement) {
			text := strings.TrimSpace(elem.Text)
			for i := range forbidden {
				if strings.Contains(text, forbidden[i]) {
					return
				}
			}
			if times <= *paragraphs && text != "" {
				fmt.Println(text)
				times++
			}
		})
	})
	c.Visit(url)
}
