package main

import "github.com/gocolly/colly"

func collectProductTags() {
	productsDetailsCollector.OnHTML("div.iqitproducttags", func(e *colly.HTMLElement) {
		ptags := make([]*Tag, 0)
		e.ForEach("ul > li > a", func(i int, el *colly.HTMLElement) {
			ptags = append(ptags, &Tag{Text: el.Text})
		})
		onProductTagsFound(ptags, src(e))
	})
}
