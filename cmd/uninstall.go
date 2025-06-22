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

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall all hooks from the .git/hooks directory",
	Run: func(cmd *cobra.Command, args []string) {
        pwd, err := os.Getwd()
        if err != nil {
            fmt.Printf("❌ Failed to get current directory path:\n %v\n", err)
            return
        }

        if err := core.UninstallHooks(pwd); err != nil {
            fmt.Printf("❌ Failed to uninstall: %v", err)
            return
        }
        fmt.Println("🎉 All hooks uninstalled successfully!")
    },
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}
