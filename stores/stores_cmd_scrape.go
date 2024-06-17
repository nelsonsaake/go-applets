package main

import (
	"github.com/spf13/cobra"
)

var scrapeCmd = &cobra.Command{
	Use:   "scrape",
	Short: "A scrapping tool, for scrapping products data, and related image files.",
}

func init() {
	rootCmd.AddCommand(scrapeCmd)
}
