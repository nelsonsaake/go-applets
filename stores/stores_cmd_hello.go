package main

import (
	"projects/saelections/pkg/sysout"

	"github.com/spf13/cobra"
)

func init() {
	hwfunc := func(cmd *cobra.Command, args []string) {
		sysout.Print("hello world...")
		if len(args) != 0 {
			sysout.Print(args)
		}
	}
	rootCmd.AddCommand(&cobra.Command{
		Use: "hello",
		Run: hwfunc,
	})
	rootCmd.AddCommand(&cobra.Command{
		Use: "hw",
		Run: hwfunc,
	})
}
