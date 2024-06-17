package main

import "github.com/gocolly/colly"

func collectSubCategories() {
	// sub categories
	SubCategoryCollector.OnHTML("ul.category-sub-menu li", func(e *colly.HTMLElement) {
		onSubCategoryFound(
			Category{
				Name:       e.ChildText("a"),
				Href:       e.ChildAttr("a", "href"),
				ParentHref: src(e),
				Categories: []*Category{},
				Products:   []*Product{},
			},
		)
	})
}
