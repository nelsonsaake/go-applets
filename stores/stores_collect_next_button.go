package main

import (
	"github.com/gocolly/colly"
)

func collectNextButton() {
	productsCollector.OnHTML("#infinity-url", func(e *colly.HTMLElement) {
		onNextBtnFound(e.Attr("href"), src(e))
	})
}
