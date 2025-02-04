package flowchart

import (
	"fmt"
	"strings"
)

type nodeShape string

// List of possible Node shapes.
// Reference: https://mermaid.js.org/syntax/flowchart.html#complete-list-of-new-shapes
const (
	// Basic shapes
	NodeShapeProcess     nodeShape = `rect`       // Rectangle
	NodeShapeEvent       nodeShape = `rounded`    // Rounded rectangle
	NodeShapeTerminal    nodeShape = `stadium`    // Stadium/pill
	NodeShapeSubprocess  nodeShape = `fr-rect`    // Framed rectangle
	NodeShapeDatabase    nodeShape = `cyl`        // Cylinder
	NodeShapeStart       nodeShape = `circle`     // Circle
	NodeShapeOdd         nodeShape = `odd`        // Odd shape
	NodeShapeDecision    nodeShape = `diam`       // Diamond
	NodeShapePrepare     nodeShape = `hex`        // Hexagon
	NodeShapeInputOutput nodeShape = `lean-r`     // Lean right
	NodeShapeOutputInput nodeShape = `lean-l`     // Lean left
	NodeShapePriority    nodeShape = `trap-b`     // Trapezoid bottom
	NodeShapeOperation   nodeShape = `trap-t`     // Trapezoid top
	NodeShapeStop        nodeShape = `dbl-circ`   // Double circle
	NodeShapeText        nodeShape = `text`       // Text block
	NodeShapeCard        nodeShape = `notch-rect` // Notched rectangle
	NodeShapeProcess2    nodeShape = `lin-rect`   // Lined/shaded rectangle
	NodeShapeSmallStart  nodeShape = `sm-circ`    // Small circle
	NodeShapeStopAlt     nodeShape = `fr-circ`    // Framed circle
	NodeShapeFork        nodeShape = `fork`       // Fork/join
	NodeShapeCollate     nodeShape = `hourglass`  // Hourglass
	NodeShapeComment     nodeShape = `brace`      // Curly brace
	NodeShapeCommentR    nodeShape = `brace-r`    // Right curly brace
	NodeShapeCommentLR   nodeShape = `braces`     // Both curly braces
	NodeShapeComLink     nodeShape = `bolt`       // Lightning bolt
	NodeShapeDocument    nodeShape = `doc`        // Document
	NodeShapeDelay       nodeShape = `delay`      // Half-rounded rectangle
	NodeShapeStorage     nodeShape = `h-cyl`      // Horizontal cylinder
	NodeShapeDisk        nodeShape = `lin-cyl`    // Lined cylinder
	NodeShapeDisplay     nodeShape = `curv-trap`  // Curved trapezoid
	NodeShapeDivided     nodeShape = `div-rect`   // Divided rectangle
	NodeShapeExtract     nodeShape = `tri`        // Triangle
	NodeShapeInternal    nodeShape = `win-pane`   // Window pane
	NodeShapeJunction    nodeShape = `f-circ`     // Filled circle
	NodeShapeLinedDoc    nodeShape = `lin-doc`    // Lined document
	NodeShapeLoopLimit   nodeShape = `notch-pent` // Notched pentagon
	NodeShapeManualFile  nodeShape = `flip-tri`   // Flipped triangle
	NodeShapeManualInput nodeShape = `sl-rect`    // Sloped rectangle
	NodeShapeMultiDoc    nodeShape = `docs`       // Stacked document
	NodeShapeMultiProc   nodeShape = `st-rect`    // Stacked rectangle
	NodeShapePaperTape   nodeShape = `flag`       // Flag/paper tape
	NodeShapeStoredData  nodeShape = `bow-rect`   // Bow tie rectangle
	NodeShapeSummary     nodeShape = `cross-circ` // Crossed circle
	NodeShapeTaggedDoc   nodeShape = `tag-doc`    // Tagged document
	NodeShapeTaggedProc  nodeShape = `tag-rect`   // Tagged rectangle
)

const (
	baseNodeShapeString string = "%s@{ shape: %s label: \"%s\"}"
	baseNodeClassString string = ":::%s"
	baseNodeStyleString string = "\tstyle %s %s\n"
)

// Node represents a node in a flowchart
type Node struct {
	ID    string
	Shape nodeShape
	Text  string
	Style *NodeStyle
	Class *Class
}

// NewNode creates a new Node with the given ID and text, setting default shape to round edges.
func NewNode(id string, text string) (newNode *Node) {
	newNode = &Node{
		ID:    id,
		Text:  text,
		Shape: NodeShapeProcess,
	}

	return
}

// SetClass sets the node class and returns the node for chaining
func (n *Node) SetClass(class *Class) *Node {
	n.Class = class
	return n
}

// SetText sets the node text and returns the node for chaining
func (n *Node) SetText(text string) *Node {
	n.Text = text
	return n
}

// SetStyle sets the style for the node and returns the node for chaining
func (n *Node) SetStyle(style *NodeStyle) *Node {
	n.Style = style
	return n
}

// SetShape sets the node shape and returns the node for chaining
func (n *Node) SetShape(shape nodeShape) *Node {
	n.Shape = shape
	return n
}

// String generates a Mermaid string representation of the node, including its shape, class, and style.
func (n *Node) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf(string(baseNodeShapeString), n.ID, string(n.Shape), n.Text))

	if n.Class != nil {
		sb.WriteString(fmt.Sprintf(string(baseNodeClassString), n.Class.Name))
	}

	sb.WriteByte('\n')

	if n.Style != nil {
		sb.WriteString(fmt.Sprintf(string(baseNodeStyleString), n.ID, n.Style.String()))
	}

	return sb.String()
}
