package core

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ezpieco/gethooky/utils"
)

func RunInit(basePath string) error {
    hookyDir := filepath.Join(basePath, utils.GetHookyDir())
    return os.Mkdir(hookyDir, 0755)
}

func AddHook(basePath string,hookName string, command string) error {
    hookyPath := filepath.Join(basePath,utils.GetHookyDir(), hookName)
    return os.WriteFile(hookyPath, []byte(command), 0644)
}

func InstallHooks(basePath string) error {
    hookyPath := filepath.Join(basePath, utils.GetHookyDir())
    gitHookPath := filepath.Join(basePath, utils.GetGitHookyDir())
    
    files, err := os.ReadDir(hookyPath)
    if err != nil {
        return fmt.Errorf("Failed to read .hooky directory:\n %v\n", err)
    }

    for _, file := range files {

        hookName := file.Name()
        hookPath := filepath.Join(hookyPath, hookName)
        gitHookPath := filepath.Join(gitHookPath, hookName)

        commandByte, err := os.ReadFile(hookPath)
        if err != nil {
            fmt.Errorf("⚠️ Failed to read %s: %v", hookPath, err)
            continue
        }

        command := strings.TrimSpace(string(commandByte))
        if command == "" {
            fmt.Errorf("Skipping empty %s hook", hookName)
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
                fmt.Errorf("⚠️ skipping %s! Existing user hook", gitHookPath)
                continue
            }
        }

        if err := os.WriteFile(gitHookPath, []byte(script), 0755); err != nil {
            fmt.Errorf("❌ Failed to write %s hook: %v", hookName, err)
            continue
        }
    }

    return nil
}
