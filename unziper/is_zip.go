package main

import (
	"path/filepath"
	"strings"
)

func IsZip(path string) bool {

	ext := filepath.Ext(path)

	ext = strings.ToLower(ext)

	return ext == ".zip"
}
