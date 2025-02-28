package utils

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/TyphonHill/go-mermaid/diagrams/utils/testutils"
)

func TestRenderToFile(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		content  string
		setup    func(string) error
		wantErr  bool
		contains []string
	}{
		{
			name:    "Write to simple file",
			path:    "test.txt",
			content: "test content",
			contains: []string{
				"test content",
			},
		},
		{
			name:    "Write to nested directory",
			path:    "test/nested/file.txt",
			content: "nested content",
			contains: []string{
				"nested content",
			},
		},
		{
			name:    "Write with multiple lines",
			path:    "multiline.txt",
			content: "line1\nline2\nline3",
			contains: []string{
				"line1",
				"line2",
				"line3",
			},
		},
		{
			name: "Write to readonly directory",
			path: "/readonly/test.txt",
			setup: func(path string) error {
				dir := filepath.Dir(path)
				if err := os.MkdirAll(dir, 0444); err != nil {
					return err
				}
				return nil
			},
			content: "test content",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create temp file
			tmpFile := testutils.CreateTempFile(t, "render_test")
			defer tmpFile.Cleanup(t)

			// Use temp file path with test filename
			testPath := filepath.Join(filepath.Dir(tmpFile.Path), tt.path)

			// Run any setup
			if tt.setup != nil {
				if err := tt.setup(testPath); err != nil {
					t.Fatalf("Setup failed: %v", err)
				}
			}

			// Test RenderToFile
			err := RenderToFile(testPath, tt.content)

			// Check error expectation
			if (err != nil) != tt.wantErr {
				t.Errorf("RenderToFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			// Verify file content
			testutils.AssertFileContent(t, testPath, tt.content)

			// Verify specific content
			content, err := os.ReadFile(testPath)
			if err != nil {
				t.Fatalf("Failed to read test file: %v", err)
			}
			testutils.AssertContains(t, string(content), tt.contains...)

			// Cleanup created directories
			dir := filepath.Dir(testPath)
			if dir != filepath.Dir(tmpFile.Path) {
				if err := os.RemoveAll(dir); err != nil {
					t.Errorf("Failed to cleanup test directory: %v", err)
				}
			}
		})
	}
}
