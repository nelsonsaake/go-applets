package main

import (
	"projects/saelections/pkg/sysout"
)

func onMainCategoryFound(c Category) {
	sysout.Verbose("main category found: %q\nfound @: %q", c.Href, c.ParentHref)
	if _, ok := cache.MainCategories[c.Href]; !ok {
		cache.MainCategories[c.Href] = &c
	}
	SubCategoryCollector.Visit(c.Href)
}
