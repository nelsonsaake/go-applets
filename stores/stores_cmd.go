package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "stores",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
