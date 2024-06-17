package main

import (
	"github.com/gocolly/colly"
)

func collectProducts() {
	// products
	productsCollector.OnHTML("article.js-product-miniature", func(e *colly.HTMLElement) {
		onProductFound(
			Product{
				Name:         e.ChildText("div.product-description > h3 > a"),
				Price:        e.ChildText("div.product-description > div.product-price-and-shipping > a > span"),
				Description:  e.ChildText("div.product-description > div.product-description-short.text-muted > a"),
				CategoryHref: e.Request.URL.String(),
				CategoryName: e.ChildText("div.product-description > div.product-category-name.text-muted"),
				Href:         e.ChildAttr("div.thumbnail-container > a", "href"),
				Tags:         []*Tag{},
			},
		)
	})
}
