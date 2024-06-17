package main

import (
	"github.com/gocolly/colly"
)

func src(e *colly.HTMLElement) string {
	return e.Request.URL.String()
}
