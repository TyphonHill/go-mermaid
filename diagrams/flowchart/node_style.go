package flowchart

import (
	"fmt"
	"strings"
)

const (
	baseNodeStyleColorString       string = "color:%s,"
	baseNodeStyleFillString        string = "fill:%s,"
	baseNodeStyleStrokeString      string = "stroke:%s,"
	baseNodeStyleStrokeWidthString string = "stroke-width:%d,"
	baseNodeStyleStrokeDashString  string = "stroke-dasharray:%s"
)

type NodeStyle struct {
	Color       string
	Fill        string
	Stroke      string
	StrokeWidth int
	StrokeDash  string
}

// NewNodeStyle creates a new NodeStyle with default stroke width and dash settings.
func NewNodeStyle() (newNodeStyle *NodeStyle) {
	newNodeStyle = &NodeStyle{
		StrokeWidth: 1,
		StrokeDash:  "0",
	}

	return
}

// String generates a formatted string representation of the node style,
// including color, fill, stroke, width, and dash properties.
func (n *NodeStyle) String() string {
	var sb strings.Builder

	if n.Color != "" {
		sb.WriteString(fmt.Sprintf(string(baseNodeStyleColorString), n.Color))
	}

	if n.Fill != "" {
		sb.WriteString(fmt.Sprintf(string(baseNodeStyleFillString), n.Fill))
	}

	if n.Stroke != "" {
		sb.WriteString(fmt.Sprintf(string(baseNodeStyleStrokeString), n.Stroke))
	}

	if n.StrokeWidth > 0 {
		sb.WriteString(fmt.Sprintf(string(baseNodeStyleStrokeWidthString), n.StrokeWidth))
	}

	if n.StrokeDash != "" {
		sb.WriteString(fmt.Sprintf(string(baseNodeStyleStrokeDashString), n.StrokeDash))
	}

	return strings.Trim(sb.String(), ",")
}
