package main

import (
	"encoding/json"
	"io/ioutil"
	"projects/saelections/pkg/ufs"
)

var (
	scrappedfile = "_scrappedfile.json"
	scrapped     = map[string]string{}

	failedscrapsfile = "_failedscrapsfile.json"
	failedscraps     = map[string]string{}
)

// ---

func AddScrapped(url string) {
	if IsValidUrl(url) {
		scrapped[url] = url
	}
}

func IsScrapped(url string) bool {
	_, ok := scrapped[url]
	return ok
}

func LoadScrappedData() {
	if !ufs.IsFile(scrappedfile) {
		return
	}

	b, err := ioutil.ReadFile(scrappedfile)
	if err != nil {
		return
	}

	json.Unmarshal(b, &scrapped)
}

func SaveScrappedData() {
	b, err := json.Marshal(scrapped)
	if err != nil {
		return
	}

	ioutil.WriteFile(scrappedfile, b, 0777) // 0777 allow all except ui
}

// ---

func AddFailedScraps(url string) {
	if IsValidUrl(url) {
		failedscraps[url] = url
	}
}

func IsFailedScraps(url string) bool {
	_, ok := failedscraps[url]
	return ok
}

func LoadFailedScrapsData() {
	if !ufs.IsFile(failedscrapsfile) {
		return
	}

	b, err := ioutil.ReadFile(failedscrapsfile)
	if err != nil {
		return
	}

	json.Unmarshal(b, &failedscraps)
}

func SaveFailedScrapsData() {
	b, err := json.Marshal(failedscraps)
	if err != nil {
		return
	}

	ioutil.WriteFile(failedscrapsfile, b, 0777) // 0777 allow all except ui
}

// ---

func init() {
	LoadScrappedData()
	OnTearDown(SaveScrappedData)
	OnTearDown(SaveFailedScrapsData)
}
