package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"projects/saelections/pkg/str"
	"projects/saelections/pkg/ufs"
	"strings"

	"github.com/spf13/cobra"
)

//go:embed _query
var _query string

func CourseToString(c Course) string {
	ln := fmt.Sprintln
	s := ln(c.Name) +
		ln(c.School) +
		ln(c.ClassStyle) +
		ln(c.Duration) +
		ln(c.ProgramStyle) +
		ln(c.LongReq) +
		ln(c.URL)
	if strings.Contains(s, "\n\n") {
		return ""
	}
	return ln(s)
}

func WriteOut() {
	courses := []Course{}
	sqlite.Find(&courses)

	b := bytes.Buffer{}
	for _, course := range courses {
		s := CourseToString(course)
		if !str.Empty(s) {
			b.WriteString(s)
		}
	}

	ufs.WriteFile(_query+"/_findamasters", b.String())
}

func init() {
	rootCmd.AddCommand(
		&cobra.Command{
			Use: "writeout",
			Run: func(cmd *cobra.Command, args []string) {
				WriteOut()
			},
		},
	)
}
