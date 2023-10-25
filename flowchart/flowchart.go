package flowchart

import (
	"fmt"
	"strings"
)

type flowchartDirection string
type curveStyle string

var Counter uint64 = 0

// Flowchart directions, defined at https://mermaid.js.org/syntax/flowchart.html#direction
const (
	FlowchartDirectionTopToBottom flowchartDirection = "TB"
	FlowchartDirectionTopDown     flowchartDirection = "TD"
	FlowchartDirectionBottomUp    flowchartDirection = "BT"
	FlowchartDirectionRightLeft   flowchartDirection = "RL"
	FlowchartDirectionLeftRight   flowchartDirection = "LR"
)

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

type Flowchart struct {
	Title      string
	Direction  flowchartDirection
	CurveStyle curveStyle
	classes    []*Class
	nodes      []*Node
	subgraphs  []*Subgraph
	links      []*Link
}

func NewFlowchart() (newFlowchart *Flowchart) {
	newFlowchart = &Flowchart{
		Direction:  FlowchartDirectionTopToBottom,
		CurveStyle: CurveStyleNone,
	}

	return
}

func NewID() (newID uint64) {
	newID = Counter
	Counter++

	return
}

func (f *Flowchart) String() string {
	var sb strings.Builder

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

	return sb.String()
}

func (f *Flowchart) AddSubgraph(title string) (newSubgraph *Subgraph) {
	newSubgraph = NewSubgraph(NewID(), title)

	f.subgraphs = append(f.subgraphs, newSubgraph)

	return
}

func (f *Flowchart) AddNode(text string) (newNode *Node) {
	newNode = NewNode(NewID(), text)

	f.nodes = append(f.nodes, newNode)

	return
}

func (f *Flowchart) AddLink(from *Node, to *Node) (newLink *Link) {
	newLink = NewLink(from, to)

	f.links = append(f.links, newLink)

	return
}

func (f *Flowchart) AddClass(name string) (newClass *Class) {
	newClass = NewClass(name)

	f.classes = append(f.classes, newClass)

	return
}
