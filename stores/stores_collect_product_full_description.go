package main

import (
	"github.com/gocolly/colly"
)

func collectProductFullDescription() {
	productsDetailsCollector.OnHTML(
		"div.product-description > div.rte-content",
		func(e *colly.HTMLElement) {
			onProductFullDescriptionFound(src(e), e.Text)
		},
	)
}
