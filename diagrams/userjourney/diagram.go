// Package userjourney provides functionality for creating Mermaid user journey diagrams
package userjourney

import (
	"strings"

	"github.com/TyphonHill/go-mermaid/diagrams/utils"
	"github.com/TyphonHill/go-mermaid/diagrams/utils/basediagram"
)

// Base string formats for user journey diagrams
const (
	baseDiagramType string = "journey\n"
)

// Diagram represents a Mermaid User Journey diagram
type Diagram struct {
	basediagram.BaseDiagram[JourneyConfigurationProperties]
	Sections []*Section
}

// NewDiagram creates a new User Journey diagram
func NewDiagram() *Diagram {
	return &Diagram{
		BaseDiagram: basediagram.NewBaseDiagram(NewJourneyConfigurationProperties()),
		Sections:    make([]*Section, 0),
	}
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

	for _, section := range d.Sections {
		sb.WriteString(section.String())
	}

	return d.BaseDiagram.String(sb.String())
}

// RenderToFile renders the diagram to a file
func (d *Diagram) RenderToFile(path string) error {
	return utils.RenderToFile(path, d.String())
}
