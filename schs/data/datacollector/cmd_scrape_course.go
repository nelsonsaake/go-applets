package main

import "github.com/spf13/cobra"

func init() {
	scrapeCmd.AddCommand(
		&cobra.Command{
			Use: "courses",
			Run: func(cmd *cobra.Command, args []string) {
				ScrapeCourses()
			},
		},
	)
}
