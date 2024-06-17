package main

import (
	"projects/saelections/pkg/sysout"

	"github.com/spf13/cobra"
)

func init() {
	scrapeCmd.AddCommand(&cobra.Command{
		Use: "data",
		Run: func(cmd *cobra.Command, args []string) {
			if err := ScrapeMarketExpress(); err != nil {
				sysout.Fatal(err)
			}
		},
	})
}
