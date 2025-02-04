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
	NodeShapeProcess          nodeShape = `rect`       // Process (Rectangle)
	NodeShapeEvent            nodeShape = `rounded`    // Event (Rounded rectangle)
	NodeShapeTerminal         nodeShape = `stadium`    // Terminal (Stadium-shaped)
	NodeShapeSubprocess       nodeShape = `fr-rect`    // Subprocess (Framed rectangle)
	NodeShapeDatabase         nodeShape = `cyl`        // Database (Cylinder)
	NodeShapeStart            nodeShape = `circle`     // Start (Circle)
	NodeShapeOdd              nodeShape = `odd`        // Odd shape (Asymmetric)
	NodeShapeDecision         nodeShape = `diam`       // Decision (Diamond)
	NodeShapePrepare          nodeShape = `hex`        // Prepare (Hexagon)
	NodeShapeInputOutput      nodeShape = `lean-r`     // Input/Output (Parallelogram)
	NodeShapeOutputInput      nodeShape = `lean-l`     // Output/Input (Alt Parallelogram)
	NodeShapeManualOperation  nodeShape = `trap-b`     // Manual Operation (Trapezoid)
	NodeShapeManual           nodeShape = `trap-t`     // Manual (Alt Trapezoid)
	NodeShapeStopDouble       nodeShape = `dbl-circ`   // Stop (Double Circle)
	NodeShapeText             nodeShape = `text`       // Text block
	NodeShapeCard             nodeShape = `notch-rect` // Card
	NodeShapeLinedProcess     nodeShape = `lin-rect`   // Lined Process (Rectangle with shadow)
	NodeShapeStartSmall       nodeShape = `sm-circ`    // Start (Small circle)
	NodeShapeStopFramed       nodeShape = `fr-circ`    // Stop (Circle with frame)
	NodeShapeForkJoin         nodeShape = `fork`       // Fork/Join
	NodeShapeCollate          nodeShape = `hourglass`  // Collate (Hourglass)
	NodeShapeComment          nodeShape = `brace`      // Comment (Left brace)
	NodeShapeCommentRight     nodeShape = `brace-r`    // Comment Right (Right brace)
	NodeShapeCommentBothSides nodeShape = `braces`     // Comment (Both braces)
	NodeShapeComLink          nodeShape = `bolt`       // Com Link (Lightning bolt)
	NodeShapeDocument         nodeShape = `doc`        // Document
	NodeShapeDelay            nodeShape = `delay`      // Delay
	NodeShapeStorage          nodeShape = `h-cyl`      // Storage (Horizontal cylinder)
	NodeShapeDiskStorage      nodeShape = `lin-cyl`    // Disk Storage (Lined cylinder)
	NodeShapeDisplay          nodeShape = `curv-trap`  // Display (Curved trapezoid)
	NodeShapeDividedProcess   nodeShape = `div-rect`   // Divided Process
	NodeShapeExtract          nodeShape = `tri`        // Extract (Triangle)
	NodeShapeInternalStorage  nodeShape = `win-pane`   // Internal Storage
	NodeShapeJunction         nodeShape = `f-circ`     // Junction (Filled circle)
	NodeShapeLinedDocument    nodeShape = `lin-doc`    // Lined Document
	NodeShapeLoopLimit        nodeShape = `notch-pent` // Loop Limit
	NodeShapeManualFile       nodeShape = `flip-tri`   // Manual File
	NodeShapeManualInput      nodeShape = `sl-rect`    // Manual Input (Sloped rectangle)
	NodeShapeMultiDocument    nodeShape = `docs`       // Multi-Document
	NodeShapeMultiProcess     nodeShape = `st-rect`    // Multi-Process
	NodeShapePaperTape        nodeShape = `flag`       // Paper Tape
	NodeShapeStoredData       nodeShape = `bow-rect`   // Stored Data
	NodeShapeSummary          nodeShape = `cross-circ` // Summary (Circle with cross)
	NodeShapeTaggedDocument   nodeShape = `tag-doc`    // Tagged Document
	NodeShapeTaggedProcess    nodeShape = `tag-rect`   // Tagged Process
)

const (
	baseNodeShapeString string = "\t%s@{ shape: %s, label: \"%s\"}"
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
