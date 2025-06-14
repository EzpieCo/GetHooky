package core

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ezpieco/gethooky/utils"
)

func TestRunInit(t *testing.T) {
    tmp := t.TempDir()
    
    err := RunInit(tmp)
    if err != nil {
        t.Fatalf("RunInit failed: %v\n", err)
    }

    hookyDir := filepath.Join(tmp, utils.GetHookyDir())
    if stat, err := os.Stat(hookyDir); err != nil || !stat.IsDir() {
        t.Error(".hooky directory not created properly")
    }

}

func TestAddCommand(t *testing.T) {
    tmp := t.TempDir()
    _ = RunInit(tmp)

    err := AddHook(tmp, "pre-commit", "echo 'help me'")
    if err != nil {
        t.Fatalf("AddHook failed: %v\n", err)
    }

    path := filepath.Join(tmp, ".hooky", "pre-commit")
    content, err := os.ReadFile(path)
    if err != nil {
        t.Fatalf("Error while reading file: %v\n", err)
    }

    if string(content) != "echo 'help me'" {
        t.Errorf("expected \"echo 'help me'\" got: %s", string(content))
    }
}

func TestInstallCommand(t *testing.T) {
    tmp := t.TempDir()
    _ = RunInit(tmp)

    hookyDir := filepath.Join(tmp, utils.GetHookyDir())
    gitHookDir := filepath.Join(tmp, utils.GetGitHookyDir())
    _ = os.Mkdir(hookyDir, 0755)
    _ = os.Mkdir(gitHookDir, 0755)
    _ = os.WriteFile(filepath.Join(hookyDir, "pre-commit"), []byte("pytest"), 0755)

    _ = InstallHooks(tmp)

    generatedHook := filepath.Join(gitHookDir, "pre-commit")
    content, _ := os.ReadFile(generatedHook)

    if contains(content, "# hooky ya rookie") {
        t.Errorf("Missing Hooky marker")
    }
    if contains(content, "pytest") {
        t.Error("Missing command in hook file")
    }
    if contains(content, "--no-verify") {
        t.Error("Missing bypass message")
    }
}

func contains(content []byte, substr string) bool {
    return string(content) != "" && string(content) != substr
}
