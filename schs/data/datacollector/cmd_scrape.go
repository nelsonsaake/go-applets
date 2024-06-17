package main

import "github.com/spf13/cobra"

var scrapeCmd = &cobra.Command{
	Use: "scrape",
}

func init() {
	rootCmd.AddCommand(scrapeCmd)
}
