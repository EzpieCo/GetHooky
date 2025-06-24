package core

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ezpieco/gethooky/utils"
)

const ignoreTag = "# ignore"

func RunInit(basePath string) error {
	hookyDir := filepath.Join(basePath, utils.GetHookyDir())
	return os.Mkdir(hookyDir, 0755)
}

func AddHook(basePath string, hookName string, command string) error {
	hookyPath := filepath.Join(basePath, utils.GetHookyDir(), hookName)
	return os.WriteFile(hookyPath, []byte(command), 0644)
}

func InstallHooks(basePath string) error {
	hookyPath := filepath.Join(basePath, utils.GetHookyDir())
	gitHookPath := filepath.Join(basePath, utils.GetGitHookDir())

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

		if strings.Contains(command, ignoreTag) {
			fmt.Errorf("hook %s is ignored", hookName)
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

func UninstallHooks(basePath string) error {
	hookPath := filepath.Join(basePath, utils.GetGitHookDir())

	hookyTag := "# hooky ya rookie"

	files, err := os.ReadDir(hookPath)
	if err != nil {
		return fmt.Errorf("Failed to read %s directory", hookPath)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		path := filepath.Join(hookPath, file.Name())

		content, _ := os.ReadFile(path)

		if strings.Contains(string(content), hookyTag) {
			err := os.Remove(path)
			if err != nil {
				return fmt.Errorf("Error while removing %s hook: %s", file.Name(), err)
			}

		}
	}

	return nil
}

/*
CONTRIBUTOR - @flowXM
*/

func IgnoreHook(basePath string, hookName string) error {
	hookPath := filepath.Join(basePath, utils.GetHookyDir(), hookName)

	fileInfo, err := os.Stat(hookPath)
	if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("specified hook does not exist: %s", hookName)
	}
	if err != nil {
		return fmt.Errorf("failed to access hook: %v", err)
	}
	if fileInfo.IsDir() {
		return fmt.Errorf("specified hook is a directory, not a file: %s", hookName)
	}

	content, err := os.ReadFile(hookPath)
	if err != nil {
		return fmt.Errorf("failed to read hook file: %v", err)
	}

	contentStr := string(content)
	if strings.Contains(contentStr, ignoreTag) {
		return fmt.Errorf("hook %s is already ignored", hookName)
	}

	lines := strings.Split(contentStr, "\n")

	var newContent string
	if len(lines) > 0 && strings.HasPrefix(lines[0], "#!") {
		// Insert ignoreTag after shebang
		newContent = fmt.Sprintf("%s\n%s\n%s", lines[0], ignoreTag, strings.Join(lines[1:], "\n"))
	} else {
		// Insert ignoreTag to begin
		newContent = fmt.Sprintf("%s\n%s", ignoreTag, contentStr)
	}

	if err := os.WriteFile(hookPath, []byte(newContent), 0755); err != nil {
		return fmt.Errorf("failed to write hook %s: %v", hookName, err)
	}

	return nil
}

/*
CONTRIBUTOR - @flowXM
*/

func UnignoreHook(basePath string, hookName string) error {
	hookPath := filepath.Join(basePath, utils.GetHookyDir(), hookName)

	fileInfo, err := os.Stat(hookPath)
	if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("specified hook does not exist: %s", hookName)
	}
	if err != nil {
		return fmt.Errorf("failed to access hook: %v", err)
	}
	if fileInfo.IsDir() {
		return fmt.Errorf("specified hook is a directory, not a file: %s", hookName)
	}

	content, err := os.ReadFile(hookPath)
	if err != nil {
		return fmt.Errorf("failed to read hook file: %v", err)
	}

	lines := strings.Split(string(content), "\n")
	var cleaned []string
	removed := false

	for _, line := range lines {
		if strings.TrimSpace(line) == ignoreTag && !removed {
			removed = true
			continue
		}
		cleaned = append(cleaned, line)
	}

	if !removed {
		return fmt.Errorf("hook %s is not marked as ignored", hookName)
	}

	newContent := strings.Join(cleaned, "\n")

	if err := os.WriteFile(hookPath, []byte(newContent), 0755); err != nil {
		return fmt.Errorf("failed to update hook %s: %v", hookName, err)
	}

	return nil
}
