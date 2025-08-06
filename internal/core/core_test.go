package core

import (
	"os"
	"path/filepath"
	"strings"
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
	gitHookDir := filepath.Join(tmp, utils.GetGitHookDir())
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

func TestUninstallHooks(t *testing.T) {
	tmp := t.TempDir()

	gitHookDir := filepath.Join(tmp, utils.GetGitHookDir())
	_ = os.MkdirAll(gitHookDir, 0755)

	hookyHook := filepath.Join(gitHookDir, "pre-commit")
	_ = os.WriteFile(hookyHook, []byte("#!/bin/sh\n# hooky ya rookie\npytest"), 0755)

	customHook := filepath.Join(gitHookDir, "pre-push")
	_ = os.WriteFile(customHook, []byte("#!/bin/sh\n echo safe"), 0755)

	err := UninstallHooks(tmp)
	if err != nil {
		t.Errorf("Uninstall Failed: %v", err)
	}

	if _, err := os.Stat(hookyHook); !os.IsNotExist(err) {
		t.Errorf("Expected hooky controlled hook to be removed")
	}

	if _, err := os.Stat(customHook); os.IsNotExist(err) {
		t.Errorf("Expected custom hook to remain")
	}

}

func TestIgnoreHook(t *testing.T) {
	tmp := t.TempDir()

	// Initialize directory structure
	err := RunInit(tmp)
	if err != nil {
		t.Fatalf("RunInit failed: %v", err)
	}

	// Create a test hook file
	hookyDir := filepath.Join(tmp, utils.GetHookyDir())
	hookName := "pre-commit"
	hookPath := filepath.Join(hookyDir, hookName)
	originalContent := "#!/bin/bash\necho Hello\n"

	if err := os.WriteFile(hookPath, []byte(originalContent), 0755); err != nil {
		t.Fatalf("Failed to write hook file: %v", err)
	}

	// Verify the file does not contain the ignore tag initially
	content, _ := os.ReadFile(hookPath)
	if strings.Contains(string(content), "# ignore") {
		t.Fatalf("Hook should not be ignored initially")
	}

	// Call IgnoreHook to add the ignore tag
	err = IgnoreHook(tmp, hookName)
	if err != nil {
		t.Fatalf("IgnoreHook failed: %v", err)
	}

	// Check that the ignore tag is added immediately after the shebang
	content, _ = os.ReadFile(hookPath)
	lines := strings.Split(string(content), "\n")
	if len(lines) < 2 {
		t.Fatal("Hook content too short after ignore")
	}
	if lines[0] != "#!/bin/bash" {
		t.Errorf("First line should be shebang, got: %s", lines[0])
	}
	if strings.TrimSpace(lines[1]) != "# ignore" {
		t.Errorf("Second line should be '# ignore', got: %s", lines[1])
	}

	// Calling IgnoreHook again should return an error indicating it's already ignored
	err = IgnoreHook(tmp, hookName)
	if err == nil || !strings.Contains(err.Error(), "already ignored") {
		t.Errorf("Expected error about already ignored, got: %v", err)
	}
}

func TestShowHook(t *testing.T) {
	tmp := t.TempDir()
	// Initialize directory structure
	err := RunInit(tmp)
	if err != nil {
		t.Fatalf("RunInit failed: %v", err)
	}
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("could not get current directory: %v", err)
	}
	if err := os.Chdir(tmp); err != nil {
		t.Fatalf("could not change to tmp dir: %v", err)
	}
	defer func() {
		if err := os.Chdir(originalDir); err != nil {
			t.Fatalf("could not revert to original dir: %v", err)
		}
	}()
	
	hookyDir := filepath.Join(tmp, utils.GetHookyDir())
	hookName := "pre-commit"
	hookPath := filepath.Join(hookyDir, hookName)
	testContent := "#!/bin/bash\necho 'test hook'"
	if err := os.WriteFile(hookPath, []byte(testContent), 0755); err != nil {
		t.Fatalf("Failed to write hook file: %v", err)
	}

	ogStdout := os.Stdout
	os.Stdout, _ = os.Open(os.DevNull)
	defer func() { os.Stdout = ogStdout }()

	// Test showing specific hook content
	err = ShowHook(tmp, hookName)
	if err != nil {
		t.Fatalf("ShowHook failed: %v", err)
	}
	os.Stdout = ogStdout
	
	// Test showing non-existent hook
	err = ShowHook(tmp, "non-existent")
	if err == nil || !strings.Contains(err.Error(), "specified hook does not exist") {
		t.Errorf("Expected error about non-existent hook, got: %v", err)
	}
}

func TestUnignoreHook(t *testing.T) {
	tmp := t.TempDir()

	// Initialize directory structure
	err := RunInit(tmp)
	if err != nil {
		t.Fatalf("RunInit failed: %v", err)
	}

	hookyDir := filepath.Join(tmp, utils.GetHookyDir())
	hookName := "pre-commit"
	hookPath := filepath.Join(hookyDir, hookName)

	// Write hook file containing the ignore tag
	contentWithIgnore := "#!/bin/bash\n# ignore\necho Hello\n"
	if err := os.WriteFile(hookPath, []byte(contentWithIgnore), 0755); err != nil {
		t.Fatalf("Failed to write hook file: %v", err)
	}

	// Verify the ignore tag exists in the file initially
	content, _ := os.ReadFile(hookPath)
	if !strings.Contains(string(content), "# ignore") {
		t.Fatalf("Hook should be ignored initially")
	}

	// Call UnignoreHook to remove the ignore tag
	err = UnignoreHook(tmp, hookName)
	if err != nil {
		t.Fatalf("UnignoreHook failed: %v", err)
	}

	// Check that the ignore tag was removed
	content, _ = os.ReadFile(hookPath)
	if strings.Contains(string(content), "# ignore") {
		t.Errorf("# ignore tag was not removed")
	}

	// Calling UnignoreHook again should return an error indicating it's not ignored
	err = UnignoreHook(tmp, hookName)
	if err == nil || !strings.Contains(err.Error(), "not marked as ignored") {
		t.Errorf("Expected error about hook not marked ignored, got: %v", err)
	}
}

func contains(content []byte, substr string) bool {
	return string(content) != "" && string(content) != substr
}
