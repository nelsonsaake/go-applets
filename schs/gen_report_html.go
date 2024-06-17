package main

import (
	"bytes"
	"html/template"
	"projects/saelections/pkg/tmpl"

	"github.com/yosssi/gohtml"
)

var (
	ReportTmplate = "report"
)

func GenReportHtml(tmplDir string, data interface{}) (report string, err error) {
	t := tmpl.NewLoader(tmplDir)
	if err = t.Load(); err != nil {
		return
	}

	tmpl, err := template.ParseFiles(t.TmplFiles...)
	if err != nil {
		return
	}

	outputBuffer := bytes.Buffer{}
	err = tmpl.ExecuteTemplate(&outputBuffer, ReportTmplate, data)
	if err != nil {
		return
	}

	return gohtml.Format(outputBuffer.String()), nil
}
