package nkspath

import (
	"io/fs"
	"path/filepath"
)

// return a list of path of directory files only
// sub-directories are left out
func DirFiles(dir string) []string {
	files := []string{}

	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {

		if !IsDirectParentChild(dir, path) {
			return nil
		}

		if IsExistingFile(path) {
			files = append(files, path)
		}

		return nil
	})

	return files
}
