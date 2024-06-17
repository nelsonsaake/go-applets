package main

import (
	"projects/saelections/pkg/str"
	"projects/saelections/pkg/sysout"
)

func mergeTags(argtags []*Tag, tag *Tag) []*Tag {
	for _, v := range argtags {
		if v.Text == tag.Text {
			return argtags
		}
	}
	return append(argtags, tag)
}

func onProductTagsFound(argtags []*Tag, phref string) {
	sysout.Verbose("product tags found @: %q", phref)
	product, ok := cache.Products[phref]
	if !ok {
		sysout.Fatal("product for tag not found")
		return
	}
	for _, v := range argtags {
		if str.Empty(v.Text) {
			continue
		}
		if _, ok := cache.Tags[v.Text]; !ok {
			cache.Tags[v.Text] = v
		}
		product.Tags = mergeTags(product.Tags, cache.Tags[v.Text])
	}
}
