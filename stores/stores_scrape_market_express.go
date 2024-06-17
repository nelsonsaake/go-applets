package main

func ScrapeMarketExpress() (err error) {
	setupCollectors()

	collectMainCategories()

	collectSubCategories()

	collectProducts()

	collectNextButton()

	collectProductTags()

	collectProductImages()

	collectProductFullDescription()

	startCollecting(scrapeurl)

	writeOutData()

	return
}
