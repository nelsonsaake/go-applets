package main

import (
	"github.com/gocolly/colly"
)

func collectProductImages() {
	productsDetailsCollector.OnHTML(
		".js-easyzoom-trigger",
		func(e *colly.HTMLElement) {
			onProductImagesFound(src(e), &Image{Url: e.Attr("href")})
		},
	)
	productsDetailsCollector.OnHTML(
		".thumb.js-thumb.img-fluid",
		func(e *colly.HTMLElement) {
			onProductImagesFound(src(e), &Image{Url: e.Attr("src")})
		},
	)
	productsDetailsCollector.OnHTML(
		".easyzoom > a",
		func(e *colly.HTMLElement) {
			onProductImagesFound(src(e), &Image{Url: e.Attr("href")})
		},
	)
	productsDetailsCollector.OnHTML(
		"#product-images-large > div > div > div > div > div > div > a",
		func(e *colly.HTMLElement) {
			onProductImagesFound(src(e), &Image{Url: e.Attr("href")})
		},
	)
	productsDetailsCollector.OnHTML(
		"#main-product-wrapper > div.row.product-info-row > div.col-md-5.col-product-info > div.product_header_container.clearfix > div.product-manufacturer.product-manufacturer-next.float-right > a > img",
		func(e *colly.HTMLElement) {
			onProductImagesFound(src(e), &Image{Url: e.Attr("src")})
		},
	)
	productsDetailsCollector.OnHTML(
		"#product-images-large > div > div > div > div > div > img",
		func(e *colly.HTMLElement) {
			onProductImagesFound(src(e), &Image{Url: e.Attr("src")})
		},
	)
}
