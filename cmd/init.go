/*
Copyright © 2025 EzpieCo <ezpie.co@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

    "github.com/ezpieco/gethooky/utils"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize a .hooky directory in the current directory",
	Run: func(cmd *cobra.Command, args []string) {

		hookyDir := utils.GetHookyDir()

		if _, err := os.Stat(hookyDir); err == nil {
			fmt.Println("🎉 .hooky directory already exist! Nothing for me to do")
			return
		} else if !os.IsNotExist(err) {
			fmt.Printf("⚠️ error while trying to find .hooky: %v\n", err)
		}

		err := os.Mkdir(hookyDir, 0755)
		if err != nil {
			fmt.Printf("❌ Task Failed successfully: %v\n", err)
			return
		}

		fmt.Println("✅ Initialized .hooky directory! Now Git Hooking! 🚀")

	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
