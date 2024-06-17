package main

import "projects/saelections/pkg/sysout"

func onSubCategoryFound(sub Category) {
	sysout.Verbose("sub category found: %q\nfound @: %q", sub.Href, sub.ParentHref)
	if _, ok := cache.SubCategories[sub.Href]; ok {
		return
	}
	if main, ok := cache.MainCategories[sub.ParentHref]; ok {
		main.Categories = append(main.Categories, &sub)
		cache.SubCategories[sub.Href] = &sub
	} else {
		sysout.Verbose("%q main category not found", sub.ParentHref)
	}
	productsCollector.Visit(sub.Href)
}
