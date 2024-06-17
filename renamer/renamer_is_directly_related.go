package main

import (
	"path/filepath"
	"strings"
)

func isDirectChild(parent, child string) bool {
	clean := func(p string) string {
		p = filepath.Clean(p)
		p = filepath.ToSlash(p)
		return p
	}
	cdir := strings.TrimSuffix(child, filepath.Base(child))
	return clean(parent) == clean(cdir)
}

func IsDirectlyRelated(a, b string) bool {
	return isDirectChild(a, b) || isDirectChild(b, a)
}

func IsNotDirectlyRelated(a, b string) bool {
	return !IsDirectlyRelated(a, b)
}
