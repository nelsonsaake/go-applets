package main

import "github.com/gocolly/colly"

var (
	cache = NewDataManager()

	mainCategoryCollector    *colly.Collector
	SubCategoryCollector     *colly.Collector
	productsCollector        *colly.Collector
	productsDetailsCollector *colly.Collector

	modelstomigrate = []interface{}{}
)

const (
	repofile         = "repo/stores.db"
	jsonfile         = "repo/stores.json"
	productsjsonfile = "repo/products.json"
	imgdir           = "public/imgs/"
	scrapeurl        = "https://marketexpress.com.gh/"
)
