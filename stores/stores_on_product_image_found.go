package main

import (
	"projects/saelections/pkg/str"
	"projects/saelections/pkg/sysout"
)

func mergeImages(argimgs []*Image, img *Image) []*Image {
	for _, v := range argimgs {
		if v.Url == img.Url {
			return argimgs
		}
	}
	return append(argimgs, img)
}

func onProductImagesFound(phref string, img *Image) {
	sysout.Verbose("product imgs found @: %q", phref)
	if str.Empty(img.Url) {
		return
	}
	product, ok := cache.Products[phref]
	if !ok {
		sysout.Verbose("product for img not found")
		return
	}
	if _, ok := cache.Images[img.Url]; !ok {
		cache.Images[img.Url] = img
	}
	product.Images = mergeImages(product.Images, cache.Images[img.Url])
}
