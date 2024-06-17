package cmd

import (
	"projects/applets/charles/services"

	"github.com/spf13/cobra"
)

func init() {

	// flags

	var (
		size int
		dir  string
	)

	// command

	var groupFilesCmd = &cobra.Command{
		Use: "group",
		Run: func(cmd *cobra.Command, args []string) {
			services.GroupFiles(dir, size)
		},
	}

	// register flags

	groupFilesCmd.PersistentFlags().IntVarP(&size, "size", "s", 10, "size files in a group.")
	groupFilesCmd.PersistentFlags().StringVarP(&dir, "dir", "d", ".", "directory path")

	// register cmd

	rootCmd.AddCommand(groupFilesCmd)
}
