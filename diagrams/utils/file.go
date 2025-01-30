package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// RenderToFile writes content to a file, handling directory creation and markdown fencing
func RenderToFile(path string, content string, markdownFence bool) error {
	// Create directory if needed
	dir := filepath.Dir(path)
	if dir != "" {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
	}

	// Add markdown fence if needed
	if markdownFence || strings.HasSuffix(path, ".md") {
		if !strings.Contains(content, "```mermaid") {
			content = "```mermaid\n" + content + "```\n"
		}
	}

	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
