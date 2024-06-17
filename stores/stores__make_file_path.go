package main

import (
	"path"
	"projects/saelections/pkg/str"
)

func MakeFilePath(url string) string {
	filename := path.Base(url)
	// url = strings.TrimSuffix(url, filename)
	class := path.Base(url)
	if str.Empty(filename) {
		return ""
	}
	return path.Join(imgdir, class, filename)
}
