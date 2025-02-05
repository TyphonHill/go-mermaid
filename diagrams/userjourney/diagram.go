// Package userjourney provides functionality for creating Mermaid user journey diagrams
package userjourney

import (
	"fmt"
	"strings"

	"github.com/TyphonHill/go-mermaid/diagrams/utils"
)

// Base string formats for user journey diagrams
const (
	baseDiagramType  string = "journey\n"
	baseDiagramTitle string = "\ttitle %s\n"
)

// Diagram represents a Mermaid User Journey diagram
type Diagram struct {
	utils.BaseDiagram
	Sections []*Section
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

// AddSection adds a new section to the diagram
func (d *Diagram) AddSection(title string) *Section {
	section := NewSection(title)
	d.Sections = append(d.Sections, section)
	return section
}

// String generates the Mermaid syntax for the diagram
func (d *Diagram) String() string {
	var sb strings.Builder

	sb.WriteString(baseDiagramType)

	if d.Title != "" {
		sb.WriteString(fmt.Sprintf(baseDiagramTitle, d.Title))
	}

	for _, section := range d.Sections {
		sb.WriteString(section.String())
	}

	return d.WrapWithFence(sb.String())
}

// RenderToFile renders the diagram to a file
func (d *Diagram) RenderToFile(path string) error {
	return utils.RenderToFile(path, d.String())
}
