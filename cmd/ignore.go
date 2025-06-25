/*
CONTRIBUTOR - @flowXM
*/

package cmd

import (
	"fmt"
	"github.com/ezpieco/gethooky/internal/core"
	"github.com/ezpieco/gethooky/utils"
	"github.com/spf13/cobra"
	"os"
)

// ignoreCmd represents the ignore command
var ignoreCmd = &cobra.Command{
	Use:   "ignore",
	Short: "Marks the specified hook as ignored. Automatically reinstalls hooks.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Printf("❌ Failed to get current directory path:\n %v\n", err)
			return
		}

		hookName := args[0]
		hookyDir := utils.GetHookyDir()
		gitHookDir := utils.GetGitHookDir()

		if _, err := os.Stat(hookyDir); os.IsNotExist(err) {
			fmt.Println("🤬 YOU IDIOT! run `hooky init` and then `hooky add` for creating hooks first!")
			return
		}

		if _, err := os.Stat(gitHookDir); os.IsNotExist(err) {
			fmt.Println("⚠️ .git/hooks directory not found! Are you inside a git repository?")
		}

		if err := core.UninstallHooks(pwd); err != nil {
			fmt.Printf("❌ Failed to uninstall: %v", err)
			return
		}

		if err := core.IgnoreHook(pwd, hookName); err != nil {
			fmt.Printf("❌ Failed to ignore hook %s: %v\n", hookName, err)
			return
		}

		if err := core.InstallHooks(pwd); err != nil {
			fmt.Printf("⚠️ Installation failed:\n %s\n", err)
			return
		}

		fmt.Printf("🎉 Hook %s is now ignored!\n", hookName)
	},
}

func init() {
	rootCmd.AddCommand(ignoreCmd)
}
