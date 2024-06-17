package main

import (
	"projects/saelections/pkg/str"
	"projects/saelections/pkg/sysout"
)

func onProductFullDescriptionFound(phref, des string) {
	sysout.Verbose("product full description found @: %q", phref)
	if str.Empty(des) {
		return
	}
	if product, ok := cache.Products[phref]; ok {
		product.FullDescription = des
	} else {
		sysout.Verbose("product for description not found")
	}
}
