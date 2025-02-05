package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// RenderToFile writes content to a file, handling directory creation
func RenderToFile(path string, content string) error {
	// Create directory if needed
	dir := filepath.Dir(path)
	if dir != "" {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
	}

	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
