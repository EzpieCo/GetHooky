/*
Copyright © 2025 EzpieCo <ezpie.co@gmail.com>
*/

package utils

import (
	"path/filepath"
)

const (
    HookyDir    = ".hooky"
    GitDir      = ".git"
    HookDir     = "hooks"
)

func GetHookyDir() string {
    return HookyDir
}

func GetGitHookDir() string {
    return filepath.Join(GitDir, HookDir)
}
