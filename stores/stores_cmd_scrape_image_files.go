package main

import (
	"projects/saelections/pkg/sysout"

	"github.com/spf13/cobra"
)

func init() {
	scrapeImageCmd.AddCommand(&cobra.Command{
		Use: "files",
		Run: func(cmd *cobra.Command, args []string) {
			isFreshScrape := false
			if _, err := ScrapeImageFiles(isFreshScrape); err != nil {
				sysout.Fatal(err)
			}
		},
	})
}
