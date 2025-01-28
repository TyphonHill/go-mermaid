package flowchart

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type flowchartDirection string
type curveStyle string

var Counter uint64 = 0

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
}

// EnableMarkdownFence enables markdown fencing in the output
func (f *Flowchart) EnableMarkdownFence() {
	f.markdownFence = true
}

// DisableMarkdownFence disables markdown fencing in the output
func (f *Flowchart) DisableMarkdownFence() {
	f.markdownFence = false
}

// Creates a new Flowchart and sets default values to some attributes
func NewFlowchart() (newFlowchart *Flowchart) {
	newFlowchart = &Flowchart{
		Direction:  FlowchartDirectionTopToBottom,
		CurveStyle: CurveStyleNone,
	}

	return
}

// Returns a new ID to be used when registering new elements that require an ID
func NewID() (newID uint64) {
	newID = Counter
	Counter++

	return
}

// Builds a new string based on the current Flowchart elements
func (f *Flowchart) String() string {
	var sb strings.Builder

	// Add markdown fence if enabled
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

	// Close markdown fence if enabled
	if f.markdownFence {
		sb.WriteString("```\n")
	}

	return sb.String()
}

// RenderToFile saves the diagram to a file at the specified path
// If the file extension is .md, markdown fencing is automatically enabled
func (f *Flowchart) RenderToFile(path string) error {
	// Create directory if it doesn't exist
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// If file has .md extension, enable markdown fencing
	originalFenceState := f.markdownFence
	if strings.ToLower(filepath.Ext(path)) == ".md" {
		f.EnableMarkdownFence()
	}

	// Generate diagram content
	content := f.String()

	// Restore original fence state
	f.markdownFence = originalFenceState

	// Write to file
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// Adds a new Subgraph to the Flowchart
func (f *Flowchart) AddSubgraph(title string) (newSubgraph *Subgraph) {
	newSubgraph = NewSubgraph(NewID(), title)

	f.subgraphs = append(f.subgraphs, newSubgraph)

	return
}

// Adds a new Node to the Flowchart
func (f *Flowchart) AddNode(text string) (newNode *Node) {
	newNode = NewNode(NewID(), text)

	f.nodes = append(f.nodes, newNode)

	return
}

// Adds a new Link to the Flowchart
func (f *Flowchart) AddLink(from *Node, to *Node) (newLink *Link) {
	newLink = NewLink(from, to)

	f.links = append(f.links, newLink)

	return
}

// Adds a new Class to the Flowchart
func (f *Flowchart) AddClass(name string) (newClass *Class) {
	newClass = NewClass(name)

	f.classes = append(f.classes, newClass)

	return
}
