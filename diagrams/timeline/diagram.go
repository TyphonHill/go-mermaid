// Package timeline provides functionality for creating Mermaid timeline diagrams
package timeline

import (
	"fmt"
	"strings"

	"github.com/TyphonHill/go-mermaid/diagrams/utils"
)

type timelineTheme string

const (
	TimelineThemeBase    timelineTheme = "base"
	TimelineThemeForest  timelineTheme = "forest"
	TimelineThemeDark    timelineTheme = "dark"
	TimelineThemeDefault timelineTheme = "default"
	TimelineThemeNeutral timelineTheme = "neutral"
)

const (
	baseInitString string = "%%%%{init: { 'theme': '%s', 'timeline': {'disableMulticolor': %t}}}%%%%\n"
)

// Diagram represents a Mermaid timeline diagram
type Diagram struct {
	utils.BaseDiagram
	Sections          []*Section
	theme             timelineTheme
	disableMulticolor bool
}

// NewDiagram creates a new timeline diagram
func NewDiagram() *Diagram {
	return &Diagram{
		Sections:          make([]*Section, 0),
		theme:             TimelineThemeBase,
		disableMulticolor: false,
	}
}

// SetTheme sets the theme for the timeline diagram
func (d *Diagram) SetTheme(theme timelineTheme) {
	d.theme = theme
}

// EnableMultiColot enables the multi-color feature for the timeline diagram
func (d *Diagram) EnableMultiColot() {
	d.disableMulticolor = false
}

// DisableMultiColot disables the multi-color feature for the timeline diagram
func (d *Diagram) DisableMultiColot() {
	d.disableMulticolor = true
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

	sb.WriteString(fmt.Sprintf(baseInitString, d.theme, d.disableMulticolor))
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
	return utils.RenderToFile(path, d.String())
}
