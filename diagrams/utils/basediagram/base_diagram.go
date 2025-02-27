package basediagram

import (
	"fmt"
	"strings"
)

const Indentation = "    "

const (
	baseDiagramSeparator = "---\n"
	baseDiagramTitle     = "title: %s\n"
)

// BaseDiagram provides common functionality for all diagram types
type BaseDiagram struct {
	title string
	ConfigurationProperties
	MarkdownFencer
}

// NewBaseDiagram creates a new BaseDiagram with default settings
func NewBaseDiagram() BaseDiagram {
	return BaseDiagram{
		title: "",
		ConfigurationProperties: ConfigurationProperties{
			theme:       ThemeDefault,
			maxTextSize: 50000,
			maxEdges:    500,
			fontSize:    16,
		},
		MarkdownFencer: MarkdownFencer{
			markdownFence: false,
		},
	}
}

// SetTitle sets the diagram title and returns the diagram for chaining
func (d *BaseDiagram) SetTitle(title string) *BaseDiagram {
	d.title = title
	return d
}

func (d *BaseDiagram) String(diagramText string) string {
	var sb strings.Builder

	sb.WriteString(baseDiagramSeparator)
	sb.WriteString(fmt.Sprintf(baseDiagramTitle, d.title))
	sb.WriteString(d.ConfigurationProperties.String())
	sb.WriteString(baseDiagramSeparator)

	sb.WriteString(diagramText)

	return d.MarkdownFencer.WrapWithFence(sb.String())
}
