package main

import (
	"github.com/spf13/cobra"
)

func init() {
	scrapeImageCmd.AddCommand(&cobra.Command{
		Use: "urls",
		Run: func(cmd *cobra.Command, args []string) {
			ScrapeImageUrls()
		},
	})
}
