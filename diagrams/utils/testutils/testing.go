package testutils

import (
	"os"
	"strings"
	"testing"
)

// TestFile represents a temporary test file
type TestFile struct {
	Path string
}

// CreateTempFile creates a temporary file for testing
func CreateTempFile(t *testing.T, prefix string) *TestFile {
	t.Helper()
	f, err := os.CreateTemp("", prefix)
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	f.Close()
	return &TestFile{Path: f.Name()}
}

// Cleanup removes the temporary file
func (tf *TestFile) Cleanup(t *testing.T) {
	t.Helper()
	if err := os.Remove(tf.Path); err != nil {
		t.Errorf("Failed to remove temp file: %v", err)
	}
}

// AssertFileContent checks if file content matches expected content
func AssertFileContent(t *testing.T, path string, want string) {
	t.Helper()
	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	if got := string(content); got != want {
		t.Errorf("File content mismatch:\nwant:\n%s\ngot:\n%s", want, got)
	}
}

// AssertContains checks if output contains all expected strings
func AssertContains(t *testing.T, output string, wants ...string) {
	t.Helper()
	for _, want := range wants {
		if !strings.Contains(output, want) {
			t.Errorf("Output missing expected content %q in:\n%s", want, output)
		}
	}
}
