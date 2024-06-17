package main

import (
	"projects/saelections/pkg/sysout"

	"github.com/spf13/cobra"
)

func init() {
	patchCmd.AddCommand(&cobra.Command{
		Use: "draft",
		Run: func(cmd *cobra.Command, args []string) {
			if err := RepoDraftFromFile(jsonfile); err != nil {
				sysout.Fatal(err)
			}
		},
	})
}
