package main

import (
	"projects/saelections/pkg/ufs"
)

func writeOutData() (err error) {
	if err = ufs.WriteFile(jsonfile, tojson(cache.MainCategories)); err != nil {
		return
	}
	if err = RepoCategoriesCreateFromMap(cache.MainCategories); err != nil {
		return
	}
	return nil
}
