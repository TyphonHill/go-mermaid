package flowchart

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type flowchartDirection string
type curveStyle string

// List of possible Flowchart directions.
// Reference: https://mermaid.js.org/syntax/flowchart.html#direction
const (
	FlowchartDirectionTopToBottom flowchartDirection = "TB"
	FlowchartDirectionTopDown     flowchartDirection = "TD"
	FlowchartDirectionBottomUp    flowchartDirection = "BT"
	FlowchartDirectionRightLeft   flowchartDirection = "RL"
	FlowchartDirectionLeftRight   flowchartDirection = "LR"
)

// List of possible Flowchart directions.
// Reference: https://mermaid.js.org/syntax/flowchart.html#styling-line-curves
const (
	CurveStyleNone       curveStyle = ""
	CurveStyleBasis      curveStyle = "basis"
	CurveStyleBumpX      curveStyle = "bumpX"
	CurveStyleBumpY      curveStyle = "bumpY"
	CurveStyleCardinal   curveStyle = "cardinal"
	CurveStyleCatmullRom curveStyle = "catmullRom"
	CurveStyleLinear     curveStyle = "linear"
	CurveStyleMonotoneX  curveStyle = "monotoneX"
	CurveStyleMonotoneY  curveStyle = "monotoneY"
	CurveStyleNatural    curveStyle = "natural"
	CurveStyleStep       curveStyle = "step"
	CurveStyleStepAfter  curveStyle = "stepAfter"
	CurveStyleStepBefore curveStyle = "stepBefore"
)

const (
	baseTitleString              string = "---\ntitle: %s\n---\n\n"
	baseCurveStyleString         string = "%%%%{ init: { 'flowchart': { 'curve': '%s' } } }%%%%\n"
	baseFlowchartDirectionString string = "flowchart %s\n"
)

// Flowcharts are composed of nodes (geometric shapes) and links (arrows or lines).
// The Mermaid code defines how nodes and links are made and accommodates different arrow types,
// multi-directional arrows, and any linking to and from subgraphs.
// Reference: https://mermaid.js.org/syntax/flowchart.html
type Flowchart struct {
	Title         string
	Direction     flowchartDirection
	CurveStyle    curveStyle
	classes       []*Class
	nodes         []*Node
	subgraphs     []*Subgraph
	links         []*Link
	markdownFence bool
	idGenerator   IDGenerator
}

// IDGenerator is an interface for generating unique IDs
type IDGenerator interface {
	NextID() uint64
}

// DefaultIDGenerator is a simple ID generator
type DefaultIDGenerator struct {
	counter uint64
}

// EnableMarkdownFence enables markdown code fencing for the flowchart output.
func (f *Flowchart) EnableMarkdownFence() {
	f.markdownFence = true
}

// DisableMarkdownFence disables markdown code fencing for the flowchart output.
func (f *Flowchart) DisableMarkdownFence() {
	f.markdownFence = false
}

// RenderToFile saves the flowchart diagram to a file at the specified path.
// If the file has a .md extension, markdown fencing is automatically enabled.
func (f *Flowchart) RenderToFile(path string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	originalFenceState := f.markdownFence
	if strings.ToLower(filepath.Ext(path)) == ".md" {
		f.EnableMarkdownFence()
	}

	content := f.String()

	f.markdownFence = originalFenceState

	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// AddSubgraph adds a new subgraph to the flowchart and returns the created subgraph.
func (f *Flowchart) AddSubgraph(title string) (newSubgraph *Subgraph) {
	newSubgraph = NewSubgraph(f.idGenerator.NextID(), title)

	f.subgraphs = append(f.subgraphs, newSubgraph)

	return
}

// AddNode adds a new node to the flowchart and returns the created node.
func (f *Flowchart) AddNode(text string) (newNode *Node) {
	newNode = NewNode(f.idGenerator.NextID(), text)

	f.nodes = append(f.nodes, newNode)

	return
}

// AddLink adds a new link between two nodes in the flowchart and returns the created link.
func (f *Flowchart) AddLink(from *Node, to *Node) (newLink *Link) {
	newLink = NewLink(from, to)

	f.links = append(f.links, newLink)

	return
}

// AddClass adds a new class to the flowchart and returns the created class.
func (f *Flowchart) AddClass(name string) (newClass *Class) {
	newClass = NewClass(name)

	f.classes = append(f.classes, newClass)

	return
}

// String generates a Mermaid flowchart string representation of the flowchart.
func (f *Flowchart) String() string {
	var sb strings.Builder

	if f.markdownFence {
		sb.WriteString("```mermaid\n")
	}

	if len(f.Title) > 0 {
		sb.WriteString(fmt.Sprintf(string(baseTitleString), f.Title))
	}

	if f.CurveStyle != CurveStyleNone {
		sb.WriteString(fmt.Sprintf(string(baseCurveStyleString), string(f.CurveStyle)))
	}

	sb.WriteString(fmt.Sprintf(string(baseFlowchartDirectionString), string(f.Direction)))

	for _, class := range f.classes {
		sb.WriteString(class.String())
	}

	for _, node := range f.nodes {
		sb.WriteString(node.String())
	}

	for _, subgraph := range f.subgraphs {
		sb.WriteString(subgraph.String("%s"))
	}

	for _, link := range f.links {
		sb.WriteString(link.String())
	}

	if f.markdownFence {
		sb.WriteString("```\n")
	}

	return sb.String()
}

// NewFlowchart creates and initializes a new Flowchart with default settings.
func NewFlowchart() (newFlowchart *Flowchart) {
	newFlowchart = &Flowchart{
		Direction:     FlowchartDirectionTopToBottom,
		CurveStyle:    CurveStyleNone,
		classes:       make([]*Class, 0),
		nodes:         make([]*Node, 0),
		subgraphs:     make([]*Subgraph, 0),
		links:         make([]*Link, 0),
		markdownFence: false,
		idGenerator:   &DefaultIDGenerator{},
	}

	return
}

// NextID returns the next unique ID for the DefaultIDGenerator.
func (g *DefaultIDGenerator) NextID() uint64 {
	current := g.counter
	g.counter++
	return current
}
