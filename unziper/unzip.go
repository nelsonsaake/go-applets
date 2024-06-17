package main

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// compressed file
func Unzip(dst, src string) error {

	ext := filepath.Ext(src)
	ext = strings.ToLower(ext)
	if ext != ".zip" {
		return errors.New("source is not a zip file")
	}

	fmt.Println("src ", src, "...")
	fmt.Println("dst ", dst, "...")

	// open zip file
	songszip, err := zip.OpenReader(src)
	if err != nil {
		return fmt.Errorf("err opening zip: %v", err)
	}
	defer songszip.Close()

	// iterate through the file
	for _, entry := range songszip.File {

		fmt.Println()

		path := filepath.Join(dst, entry.Name)

		// create dir
		if entry.FileInfo().IsDir() {

			fmt.Println("creating dir ", path, "...")

			err = os.MkdirAll(path, os.ModePerm)
			if err != nil {
				return fmt.Errorf("err creating dir: %v", err)
			}

			fmt.Println("created dir ", path)
			continue
		}

		// create file
		fmt.Println("creating file ", path, "...")

		dir := filepath.Dir(path)

		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("err creating path for file: %v", err)
		}

		file, err := entry.Open()
		if err != nil {
			return fmt.Errorf("err opening file: %v", err)
		}
		defer file.Close()

		dstFile, err := os.Create(path)
		if err != nil {
			return fmt.Errorf("err creating dst file: %v", err)
		}
		defer dstFile.Close()

		_, err = io.Copy(dstFile, file)
		if err != nil {
			return fmt.Errorf("err copy file: %v", err)
		}

		fmt.Println("created file ", path)
	}

	return nil
}
