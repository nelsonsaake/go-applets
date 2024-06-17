package main

import (
	"projects/saelections/pkg/sysout"
	"projects/saelections/pkg/ufs"
)

type ImageUrlScrapper struct {
	Products            []Product
	VisitErrorCount     int
	SaveImageErrorCount int
}

func (s *ImageUrlScrapper) StartScrapping() {
	handleError := func(err error) {
		if err != nil {
			sysout.Verbose(err)
			s.VisitErrorCount++
		}
	}
	for _, p := range s.Products {
		cache.Products[p.Href] = &p
		err := productsDetailsCollector.Visit(p.Href)
		handleError(err)
	}
}

func (s *ImageUrlScrapper) SaveImagesToRepo() {
	handleError := func(err error) {
		if err != nil {
			sysout.Verbose(err)
			s.SaveImageErrorCount++
		}
	}
	for _, p := range cache.Products {
		for _, i := range p.Images {
			i.ProductID = p.ID
			err := RepoImageSave(i)
			handleError(err)
		}
	}
}

func (s *ImageUrlScrapper) SaySummary() {
	sysout.Print("%v products visit errors", s.VisitErrorCount)
	sysout.Print("%v save product errors", s.SaveImageErrorCount)
}

func (s *ImageUrlScrapper) WriteToJson() (err error) {
	err = ufs.WriteFile(productsjsonfile, tojson(cache.Products))
	return
}

func (s *ImageUrlScrapper) LoadProducts() (err error) {
	s.Products, err = RepoProductsWithNoImages()
	return err
}

func (s *ImageUrlScrapper) ScrapeImageUrls() {
	setupCollectors()

	collectProductImages()

	s.StartScrapping()

	s.SaveImagesToRepo()

	s.SaySummary()
}

// collect product-images for products with no images
func ScrapeImageUrls() {
	s := ImageUrlScrapper{}

	if err := s.LoadProducts(); err != nil {
		sysout.Fatal(err)
	}

	s.ScrapeImageUrls()

	if err := s.WriteToJson(); err != nil {
		sysout.Fatal(err)
	}
}

func ScrapeImageUrlsFor(p Product) {
	s := ImageUrlScrapper{}
	s.Products = append(s.Products, p)
	s.ScrapeImageUrls()
}
