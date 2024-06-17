package main

import (
	"github.com/spf13/cobra"
)

var scrapeImageCmd = &cobra.Command{
	Use:   "image",
	Short: "scrape image urls, scrape image files",
}

func init() {
	scrapeCmd.AddCommand(scrapeImageCmd)
}
