package main

import (
	"fmt"
	"path/filepath"
)

var count = 1

func Index(name string) string {
	name = fmt.Sprintf("%02d%s", count, filepath.Ext(name)) // eg. 01.png
	count++                                                 // increment
	return name
}
