/*
Copyright © 2025 EzpieCo <ezpie.co@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/ezpieco/gethooky/internal/core"
	"github.com/ezpieco/gethooky/utils"
	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install all hooks into .git/hooks",
	Run: func(cmd *cobra.Command, args []string) {
        pwd, err := os.Getwd()
        if err != nil {
            fmt.Printf("❌ Failed to get current directory path:\n %v\n", err)
            return
        }

        hookyDir := utils.GetHookyDir()
		gitHookDir := utils.GetGitHookyDir()

		if _, err := os.Stat(hookyDir); os.IsNotExist(err) {
			fmt.Println("🤬 YOU IDIOT! run `hooky init` and then `hooky add` for creating hooks first!")
			return
		}

		if _, err := os.Stat(gitHookDir); os.IsNotExist(err) {
			fmt.Println("⚠️ .git/hooks directory not found! Are you inside a git repository?")
		}

        if err := core.InstallHooks(pwd); err != nil {
            fmt.Printf("⚠️ Installation failed:\n %s\n", err)
            return
        }

        fmt.Println("🎉 All hooks installed successfully!")
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
