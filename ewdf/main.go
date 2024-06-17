package main

import (
	"bytes"
	"fmt"
	"io/fs"
	"path/filepath"
	"projects/semper-server/pkg/ufs"
)

// we want to get some situational awareness on the .ewdf infection
// we count all other files
// when we find a file with extension .ewdf; we count
// we display some stats and the percentage infection

func main() {

	var (
		ewdfFile  = "ewdf_file.txt"
		statsFile = "stats_file.txt"
	)
	cleanFile(ewdfFile)
	cleanFile(statsFile)

	dir := "d:/"
	fileCount := 0
	ewdfFileCount := 0
	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if ufs.IsFile(path) {
			fileCount++
			if filepath.Ext(path) == ".ewdf" {
				ewdfFileCount++
				appendFile(ewdfFile, fmt.Sprintln(path))
			}
			buff := bytes.Buffer{}
			buff.WriteString(fmt.Sprintln("file count: ", fileCount))
			buff.WriteString(fmt.Sprintln("ewdf file count: ", ewdfFileCount))

			percentage := (float64(ewdfFileCount) * 100.0) / float64(fileCount)
			percentageStr := fmt.Sprintf("%.2f", percentage)
			buff.WriteString(fmt.Sprintln("ewdf infection: ", percentageStr, "%"))

			writeFile(statsFile, buff.String())

			fmt.Println(fileCount, "files scanned; current file: ", path)
		}
		return nil
	})
}
