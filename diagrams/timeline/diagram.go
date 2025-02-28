// Package timeline provides functionality for creating Mermaid timeline diagrams
package timeline

import (
	"strings"

	"github.com/TyphonHill/go-mermaid/diagrams/utils"
	"github.com/TyphonHill/go-mermaid/diagrams/utils/basediagram"
)

// Base string formats for timeline diagrams
const (
	baseDiagramType string = "timeline\n"
)

// Diagram represents a Mermaid timeline diagram
type Diagram struct {
	basediagram.BaseDiagram[TimelineConfigurationProperties]
	Sections []*Section
}

// NewDiagram creates a new timeline diagram
func NewDiagram() *Diagram {
	return &Diagram{
		BaseDiagram: basediagram.NewBaseDiagram(NewTimeLineConfigurationProperties()),
		Sections:    make([]*Section, 0),
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

	sb.WriteString(baseDiagramType)

	for _, section := range d.Sections {
		sb.WriteString(section.String())
	}

	return d.BaseDiagram.String(sb.String())
}

// RenderToFile saves the diagram to a file at the specified path
func (d *Diagram) RenderToFile(path string) error {
	return utils.RenderToFile(path, d.String())
}
