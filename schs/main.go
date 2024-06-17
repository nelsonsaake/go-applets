package main

import (
	"projects/saelections/pkg/htmlpdf"
	"projects/saelections/pkg/sysout"
	"projects/saelections/pkg/ufs"
)

func main() {
	defer cleanup()

	reportHtml, err := GenReportHtml("tmpl", data())
	if err != nil {
		panic(err)
	}

	if err := ufs.WriteFile(_tempfile, reportHtml); err != nil {
		panic(err)
	}

	if err := htmlpdf.ConvertHtmlToPdf(_tempfile, _output); err != nil {
		panic(err)
	}

	sysout.Print("Report generated successfully")
}
