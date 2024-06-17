package main

import (
	"fmt"
	"html/template"
	"path/filepath"
)

type Page struct {
	Heading string
	HTML    template.HTML
}

type ReportData struct {
	Pages []Page
}

// breaks the data in one file into multiple files to avoid overflow in the pdf document
func DeepSplit(file File) []File {
	const maxcharcount = 1652 * 2 / 3

	newfilecopy := func() File {
		f := file
		f.Schs = []Sch{}
		return f
	}

	countchar := func(sch Sch) int {
		return len(fmt.Sprint(sch))
	}

	var (
		files         = []File{}
		tempfile      = newfilecopy()
		tempcharcount = 0
	)

	for _, sch := range file.Schs {

		if tempcharcount+countchar(sch) > maxcharcount {
			files = append(files, tempfile)
			tempfile = newfilecopy()
			tempcharcount = 0
		}

		tempfile.Schs = append(tempfile.Schs, sch)
		tempcharcount += countchar(sch)
	}

	files = append(files, tempfile)

	return files
}

func data() ReportData {
	dir := func(file File) string {
		return filepath.Base(filepath.Dir(file.datafile))
	}

	files := []File{}
	for _, file := range LoadData() {
		files = append(files, DeepSplit(file)...)
	}

	pages := []Page{}
	for _, file := range files {
		page := Page{Heading: dir(file), HTML: template.HTML(file.HTML())}
		pages = append(pages, page)
	}

	return ReportData{pages}
}
