package nkspath

import "path/filepath"

func Normalize(path string) string {

	path = filepath.ToSlash(path)
	path = filepath.Clean(path)
	return path
}
