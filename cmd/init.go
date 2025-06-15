/*
Copyright © 2025 EzpieCo <ezpie.co@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/ezpieco/gethooky/internal/core"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize a .hooky directory in the current directory",
	Run: func(cmd *cobra.Command, args []string) {

        pwd, err := os.Getwd()
        if err != nil {
            fmt.Printf("❌ Failed to get current directory path:\n %v\n", err)
        }

        if err := core.RunInit(pwd); err != nil {
            fmt.Printf("❌ Could not create .hooky directory:\n %v\n", err)
            return
        }

		fmt.Println("✅ Initialized .hooky directory! Now Git Hooking! 🚀")

	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
