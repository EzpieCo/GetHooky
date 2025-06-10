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

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install all hooks into .git/hooks",
	Run: func(cmd *cobra.Command, args []string) {
		hookyDir := utils.GetHookyDir()
		gitHookDir := utils.GetGitHookyDir()

		if _, err := os.Stat(hookyDir); os.IsNotExist(err) {
			fmt.Println("🤬 YOU IDIOT! run `hooky init` and then `hooky add` for creating hooks first!")
			return
		}

		if _, err := os.Stat(gitHookDir); os.IsNotExist(err) {
			fmt.Println("⚠️ .git/hooks directory not found! Are you inside a git repository?")
		}

		files, _ := os.ReadDir(hookyDir)

		for _, file := range files {
			if file.IsDir() {
				continue
			}

			hookName := file.Name()
			hookPath := filepath.Join(hookyDir, hookName)
			gitHookPath := filepath.Join(gitHookDir, hookName)

			commandByte, err := os.ReadFile(hookPath)
			if err != nil {
				fmt.Printf("⚠️ Failed to read %s: %v", hookPath, err)
				continue
			}

			command := strings.TrimSpace(string(commandByte))
			if command == "" {
				fmt.Printf("Skipping empty %s hook", hookName)
				continue
			}

			script := fmt.Sprintf(`#!/bin/sh
# hooky ya rookie

%s

if [ $? -ne 0 ]; then
  echo ""
  echo "🚫 Hook '%s' failed."
  echo "👉 To bypass, use: git commit --no-verify"
  echo ""
  exit 1
fi
`, command, hookName)

			if existing, err := os.ReadFile(gitHookPath); err == nil {
				if !strings.Contains(string(existing), "# hooky ya rookie") {
					fmt.Printf("⚠️ skipping %s! Existing user hook", gitHookPath)
					continue
				}
			}

			if err := os.WriteFile(gitHookPath, []byte(script), 0755); err != nil {
				fmt.Printf("❌ Failed to write %s hook: %v", hookName, err)
				continue
			}

			fmt.Printf("✅ Installed %s hook successfully!\n", hookName)
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
