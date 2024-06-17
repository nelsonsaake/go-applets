package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

func zipFiles(fileName string, files []string) (err error) {

	// create a zip file
	// create a writter for the zip file
	// we go through the list and write files to zip

	zipFile, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("error creating zip file: %v", err)
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, file := range files {
		if err = addAFileToZip(zipWriter, file); err != nil {
			return fmt.Errorf("error adding file to zip: %v", err)
		}
	}

	return
}

func addAFileToZip(zipWriter *zip.Writer, fileName string) (err error) {

	fileToBeZipped, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer fileToBeZipped.Close()

	info, err := fileToBeZipped.Stat()
	if err != nil {
		return fmt.Errorf("error getting statistics of file to be zipped: %v", err)
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return fmt.Errorf("error geting fileinfoheader: %v", err)
	}

	header.Name = fileName
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return fmt.Errorf("error creating header: %v", err)
	}

	_, err = io.Copy(writer, fileToBeZipped)
	if err != nil {
		return fmt.Errorf("error copying contetnal: %v", err)
	}

	return
}

// func zip_usecase_example() {

// 	// get the zip name
// 	fileName := "sample.zip"

// 	// the files we want to zip
// 	files := []string{
// 		"/main.go",
// 		"/procfile",
// 	}

// 	// we pass to the create zip function
// 	if err := zipFiles(fileName, files); err != nil {
// 		fmt.Printf("Error zipping files: %v", err)
// 	}
// }
