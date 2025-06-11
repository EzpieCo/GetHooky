/*
Copyright © 2025 EzpieCo <ezpie.co@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ezpieco/gethooky/utils"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a git hook to the .hook directory",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		hookName := args[0]
		command := strings.Join(args[1:], " ")

		hookyDir := utils.GetHookyDir()

		if _, err := os.Stat(hookyDir); os.IsNotExist(err) {
			fmt.Println("⚠️ Create .hooky directory with `hooky init` first!")
			return
		}

		if strings.Contains(hookName, "/") || strings.Contains(hookName, `\`) {
			fmt.Println("❌ Invalid hook name. Please provide only the hook name like 'pre-commit'")
			return
		}

		filePath := filepath.Join(hookyDir, hookName)
		err := os.WriteFile(filePath, []byte(command+"\n"), 0644)
		if err != nil {
			fmt.Printf("❌ Task Failed successfully: %v\n", err)
			return
		}

		fmt.Printf("✅ Added hook to .hooky/%s 🎉 \n", hookName)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
