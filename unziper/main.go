package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	dir := "."

	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {

		if IsZip(path) {

			src := path
			ext := filepath.Ext(src)
			dst := strings.TrimSuffix(src, ext)

			err := Unzip(dst, src)
			if err != nil {
				panic(err)
			}

			err = os.Remove(src)
			if err != nil {
				panic(err)
			}
		}

		return nil
	})
}
