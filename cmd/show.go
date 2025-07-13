/*
CONTRIBUTOR - @thatonecodes
*/

package cmd

import (
	"fmt"
	"os"

	"github.com/ezpieco/gethooky/internal/core"
	"github.com/ezpieco/gethooky/utils"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show [hookname]",
	Short: "Show all hooks or the contents of a specific hook",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Printf("❌ Failed to get current directory path:\n %v\n", err)
			return
		}

		var hookName string
		if len(args) == 1 {
			hookName = args[0]
		}

		hookyDir := utils.GetHookyDir()
		if _, err := os.Stat(hookyDir); os.IsNotExist(err) {
			fmt.Println("❌ .hooky directory not found! Run `hooky init` first!")
			return
		}

		if err := core.ShowHook(pwd, hookName); err != nil {
			fmt.Printf("❌ %v\n", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
