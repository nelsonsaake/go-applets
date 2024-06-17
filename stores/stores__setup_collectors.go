package main

import (
	"projects/saelections/pkg/sysout"

	"github.com/gocolly/colly"
)

func newCollector() *colly.Collector {
	newCollector := colly.NewCollector(
		colly.AllowedDomains("https://marketexpress.com.gh/", "marketexpress.com.gh"),
	)
	onRequest := func(request *colly.Request) {
		sysout.Verbose("Visiting", request.URL.String())
	}
	onError := func(r *colly.Response, e error) {
		sysout.Verbose("collector error: %v\n@: %s", e, r.Request.URL.String())
	}
	newCollector.OnRequest(onRequest)
	newCollector.OnError(onError)
	return newCollector
}

func setupCollectors() {
	mainCategoryCollector = newCollector()
	SubCategoryCollector = newCollector()
	productsCollector = newCollector()
	productsDetailsCollector = newCollector()
}
