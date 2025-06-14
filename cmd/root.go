/*
GetHooky - a simple CLI tool to help you manage your git hooks and share them with your team

Copyright © 2025 EzpieCo <ezpie.co@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hooky",
	Short: "A simple git hook manager that can handle your hooks",
    CompletionOptions: cobra.CompletionOptions{
        DisableDefaultCmd: true, // this gets rid of the default completion cmd
    },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
