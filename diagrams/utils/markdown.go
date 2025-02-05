package utils

import "strings"

// MarkdownFencer provides common functionality for handling markdown fence state
type MarkdownFencer struct {
	markdownFence bool
}

// EnableMarkdownFence enables markdown fence in output and returns receiver for chaining
func (m *MarkdownFencer) EnableMarkdownFence() *MarkdownFencer {
	m.markdownFence = true
	return m
}

// DisableMarkdownFence disables markdown fence in output
func (m *MarkdownFencer) DisableMarkdownFence() {
	m.markdownFence = false
}

// IsMarkdownFenceEnabled returns current markdown fence state
func (m *MarkdownFencer) IsMarkdownFenceEnabled() bool {
	return m.markdownFence
}

// WrapWithFence wraps content with markdown fence if enabled
func (m *MarkdownFencer) WrapWithFence(content string) string {
	if !m.markdownFence {
		return content
	}

	var sb strings.Builder
	sb.WriteString("```mermaid\n")
	sb.WriteString(content)
	sb.WriteString("\n```\n")
	return sb.String()
}
