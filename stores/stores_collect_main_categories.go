package main

import (
	"github.com/gocolly/colly"
)

func collectMainCategories() {
	// main categories
	mainCategoryCollector.OnHTML(".cbp-hrmenu-tab", func(e *colly.HTMLElement) {
		onMainCategoryFound(
			Category{
				Name:       e.ChildText("a > span"),
				Href:       e.ChildAttr("a", "href"),
				ParentHref: src(e),
				Categories: []*Category{},
				Products:   []*Product{},
			},
		)
	})
}
