package userjourney

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Diagram represents a Mermaid User Journey diagram
type Diagram struct {
	Title         string
	Sections      []*Section
	markdownFence bool
}

// NewDiagram creates a new User Journey diagram
func NewDiagram() *Diagram {
	return &Diagram{
		Sections: make([]*Section, 0),
	}
}

// SetTitle sets the diagram title and returns the diagram for chaining
func (d *Diagram) SetTitle(title string) *Diagram {
	d.Title = title
	return d
}

// EnableMarkdownFence enables markdown fence in output
func (d *Diagram) EnableMarkdownFence() *Diagram {
	d.markdownFence = true
	return d
}

// DisableMarkdownFence disables markdown fence in output
func (d *Diagram) DisableMarkdownFence() {
	d.markdownFence = false
}

// AddSection adds a new section to the diagram
func (d *Diagram) AddSection(title string) *Section {
	section := NewSection(title)
	d.Sections = append(d.Sections, section)
	return section
}

// String generates the Mermaid syntax for the diagram
func (d *Diagram) String() string {
	var sb strings.Builder

	if d.markdownFence {
		sb.WriteString("```mermaid\n")
	}

	sb.WriteString("journey\n")

	if d.Title != "" {
		sb.WriteString(fmt.Sprintf("\ttitle %s\n", d.Title))
	}

	for _, section := range d.Sections {
		sb.WriteString(section.String())
	}

	if d.markdownFence {
		sb.WriteString("```\n")
	}

	return sb.String()
}

// RenderToFile renders the diagram to a file
func (d *Diagram) RenderToFile(path string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	originalFenceState := d.markdownFence
	if strings.ToLower(filepath.Ext(path)) == ".md" {
		d.EnableMarkdownFence()
	}

	content := d.String()
	d.markdownFence = originalFenceState

	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
