package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"projects/semper-server/pkg/ufs"

	_ "embed"
)

//go:embed _readme.txt
var _readme string

func main() {
	dir := "c:/"
	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if d.Name() == "_readme.txt" {
			fileContent, err := ufs.ReadFile(path)
			if err != nil {
				panic(err)
			}

			if fileContent == _readme {
				err := ufs.DelFile(path)
				if err != nil {
					panic(err)
				}

				fmt.Println("deleted ransom note: ", path)
			}

		}
		return nil
	})
}
