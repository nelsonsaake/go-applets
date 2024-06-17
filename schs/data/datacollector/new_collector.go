package main

import (
	"projects/saelections/pkg/sysout"

	"github.com/gocolly/colly"
)

func NewCollector() *colly.Collector {
	var (
		say = sysout.Print
		ctr = colly.NewCollector()
	)

	ctr.OnRequest(func(r *colly.Request) {
		say("visiting", r.URL)
	})

	ctr.OnError(func(r *colly.Response, e error) {
		say("error scrapping %q: %v", r.Request.URL, e)
	})

	ctr.OnError(func(r *colly.Response, e error) {
		AddFailedScraps(r.Request.URL.String())
	})

	ctr.OnScraped(func(r *colly.Response) {
		say("scraped", r.Request.URL)
	})

	ctr.OnScraped(func(r *colly.Response) {
		AddScrapped(r.Request.URL.String())
	})

	return ctr
}
