package main

import (
	"github.com/spf13/cobra"
)

var patchCmd = &cobra.Command{
	Use:   "patch",
	Short: "A repair tool, for patching up scrapped data",
}

func init() {
	rootCmd.AddCommand(patchCmd)
}
