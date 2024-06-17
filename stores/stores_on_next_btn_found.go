package main

import (
	"projects/saelections/pkg/str"
	"projects/saelections/pkg/sysout"
)

func onNextBtnFound(next, src string) {
	sysout.Verbose("next found: %q\nfound @: %q", next, src)
	if str.Empty(next) {
		return
	}
	if _, ok := cache.SubCategories[next]; !ok {
		if category, ok := cache.SubCategories[src]; ok {
			cache.SubCategories[next] = category
		} else {
			sysout.Verbose("can't find category for next button: %q", src)
		}
	}
	productsCollector.Visit(next)
}
