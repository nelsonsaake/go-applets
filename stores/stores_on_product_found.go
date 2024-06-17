package main

import "projects/saelections/pkg/sysout"

func onProductFound(p Product) {
	sysout.Verbose("product found: %q\nfound @: %q", p.Href, p.CategoryHref)
	if _, ok := cache.Products[p.Href]; !ok {
		cache.Products[p.Href] = &p
	}
	if category, ok := cache.SubCategories[p.CategoryHref]; ok {
		category.Products = append(category.Products, &p)
	} else {
		sysout.Verbose("%q product category not found", p.CategoryHref)
	}
	productsDetailsCollector.Visit(p.Href)
}
