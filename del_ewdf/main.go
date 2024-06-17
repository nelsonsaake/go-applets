package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"projects/semper-server/pkg/ufs"
)

// surgically remove corrupted files only
// corrupted files == files ending with .ewdf

func main() {
	dir := "D:/desk"
	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if ufs.IsFile(path) {
			if filepath.Ext(path) == ".ewdf" {
				delFile(path)
				fmt.Println("file deleted: ", path)
			}
		}
		return nil
	})
}
