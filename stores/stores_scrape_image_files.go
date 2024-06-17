package main

import (
	"fmt"
	"projects/saelections/pkg/download"
	"projects/saelections/pkg/sysout"
	"projects/saelections/pkg/ufs"
)

type ImageFileScrapper struct {
	faileddownloads     int64
	failedcreate        int64
	successfuldownloads int64
	urls                []string
	ReDownload          bool
}

func (s *ImageFileScrapper) SayRepoSaveProductFailed(img Image, err error) {
	f := "failed to create product img %q: err %v"
	sysout.Verbose(f, img.Url, err)
}

func (s *ImageFileScrapper) SayRepoSaveProductSuccessful(img Image) {
	sysout.Verbose("new img save to db: %q", img.Url)
}

func (s *ImageFileScrapper) RepoGetImageUrlsFromDB() (err error) {
	s.urls, err = RepoImageUrls()
	if err != nil {
		return fmt.Errorf("error gettings product imgs urls: %v", err)
	}
	return nil
}

func (s *ImageFileScrapper) LoadURLS() (err error) {
	if err := s.RepoGetImageUrlsFromDB(); err != nil {
		return err
	}
	return nil
}

func (s *ImageFileScrapper) HandleRepoSaveProductError(img Image, err error) {
	if err != nil {
		s.SayRepoSaveProductFailed(img, err)
		s.failedcreate++
	} else {
		s.SayRepoSaveProductSuccessful(img)
	}
}

func (s *ImageFileScrapper) SaveImage(url, fpath string) (img Image, err error) {
	img = Image{Url: url, Local: fpath}
	err = RepoImageSave(&img)
	s.HandleRepoSaveProductError(img, err)
	return
}

func (s *ImageFileScrapper) SayDownloadImgFailed(url string, err error) {
	f := "failed to download img %q: err %v"
	sysout.Verbose(f, url, err)
}

func (s *ImageFileScrapper) SayDownloadImgSuccessful(url string) {
	sysout.Verbose("new img downloaded: %q", url)
}

func (s *ImageFileScrapper) HandleDownloadImgError(url string, err error) {
	if err != nil {
		s.SayDownloadImgFailed(url, err)
		s.faileddownloads++
	} else {
		s.SayDownloadImgSuccessful(url)
		s.successfuldownloads++
	}
}

func (s *ImageFileScrapper) ShouldSkipDownloadUrl(url string) bool {
	fpath := MakeFilePath(url)
	mustNotDownload := !s.ReDownload
	yes, _ := ufs.Exists(fpath)
	return yes && mustNotDownload
}

func (s *ImageFileScrapper) SaySkippingDownload(url string) {
	sysout.Verbose("already downloaded, skipping: %v", url)
}

func (s *ImageFileScrapper) DownloadImg(url string) (fpath string, err error) {
	fpath = MakeFilePath(url)
	if s.ShouldSkipDownloadUrl(url) {
		s.SaySkippingDownload(url)
		return
	}
	err = download.Image(url, fpath)
	s.HandleDownloadImgError(url, err)
	return
}

func (s *ImageFileScrapper) SayScrapingImages() {
	sysout.Verbose("scrapping product imgs")
}

func (s *ImageFileScrapper) SayScrapingImagesSummary() {
	f := "%v images failed to download and %v create failed\n%v products where successfully downloaded"
	sysout.Verbose(f, s.faileddownloads, s.failedcreate, s.successfuldownloads)
}

func (s *ImageFileScrapper) ScrapeProductsImgs() (imgs []*Image, err error) {
	s.SayScrapingImages()

	for _, url := range s.urls {
		fpath, err := s.DownloadImg(url)
		if err != nil {
			continue
		}

		img, err := s.SaveImage(url, fpath)
		if err != nil {
			continue
		}

		imgs = append(imgs, &img)
	}

	s.SayScrapingImagesSummary()
	return imgs, nil
}

func ScrapeImageFiles(isFreshScrape bool) (imgs []*Image, err error) {
	s := ImageFileScrapper{ReDownload: isFreshScrape}
	if err = s.RepoGetImageUrlsFromDB(); err != nil {
		return nil, err
	}
	return s.ScrapeProductsImgs()
}

func ScrapeImageFilesFor(p Product) (imgs []*Image, err error) {
	s := ImageFileScrapper{ReDownload: false}
	for _, img := range p.Images {
		s.urls = append(s.urls, img.Url)
	}
	return s.ScrapeProductsImgs()
}
