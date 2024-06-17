package main

import (
	"fmt"
	"projects/saelections/pkg/htmlpdf"
	"projects/saelections/pkg/sysout"
	"projects/saelections/pkg/ufs"
)

func delfile(file string) {
	if err := ufs.DelFile(file); err != nil {
		fmt.Println(err)
	}
}

func main() {
	var (
		htmlfile = "report.html"
		pdffile  = "out.pdf"
	)

	delfile(htmlfile)
	defer delfile(htmlfile)

	reportHtml, err := GenReportHtml("tmpl", data())
	if err != nil {
		panic(err)
	}

	if err := ufs.WriteFile(htmlfile, reportHtml); err != nil {
		panic(err)
	}

	if err := htmlpdf.ConvertHtmlToPdf(htmlfile, pdffile); err != nil {
		panic(err)
	}

	sysout.Print("Report generated successfully")
}
