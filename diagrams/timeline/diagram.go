// Package timeline provides functionality for creating Mermaid timeline diagrams
package timeline

import (
	"fmt"
	"strings"

	"github.com/TyphonHill/go-mermaid/diagrams/utils"
)

// Diagram represents a Mermaid timeline diagram
type Diagram struct {
	utils.BaseDiagram
	Sections []*Section
}

// NewDiagram creates a new timeline diagram
func NewDiagram() *Diagram {
	return &Diagram{
		Sections: make([]*Section, 0),
	}
}

// AddSection creates and adds a new section to the timeline
func (d *Diagram) AddSection(title string) *Section {
	section := NewSection(title)
	d.Sections = append(d.Sections, section)
	return section
}

// String generates the Mermaid syntax for the timeline diagram
func (d *Diagram) String() string {
	var sb strings.Builder

	sb.WriteString("timeline\n")

	if d.Title != "" {
		sb.WriteString(fmt.Sprintf("\ttitle %s\n", d.Title))
	}

	for _, section := range d.Sections {
		sb.WriteString(section.String())
	}

	return d.WrapWithFence(sb.String())
}

// RenderToFile saves the diagram to a file at the specified path
func (d *Diagram) RenderToFile(path string) error {
	return utils.RenderToFile(path, d.String(), d.IsMarkdownFenceEnabled())
}
