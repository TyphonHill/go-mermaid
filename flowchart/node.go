package flowchart

import (
	"fmt"
	"strings"
)

type nodeShape string

const (
	NodeShapeRoundEdges       nodeShape = `("%s")`
	NodeShapeStadium          nodeShape = `(["%s"])`
	NodeShapeSubRoutine       nodeShape = `[["%s"]]`
	NodeShapeCylindrical      nodeShape = `[("%s")]`
	NodeShapeCircle           nodeShape = `(("%s"))`
	NodeShapeAsymmetric       nodeShape = `>"%s"]`
	NodeShapeRhombus          nodeShape = `{"%s"}`
	NodeShapeHexagon          nodeShape = `{{"%s"}}`
	NodeShapeParallelogram    nodeShape = `[/"%s"/]`
	NodeShapeParallelogramAlt nodeShape = `[\"%s"\]`
	NodeShapeTrapezoid        nodeShape = `[/"%s"\]`
	NodeShapeTrapezoidAlt     nodeShape = `[\"%s"/]`
	NodeShapeDoubleCircle     nodeShape = `((("%s")))`
)

const (
	baseNodeShapeString string = "\t%d%s"
	baseNodeClassString string = ":::%s"
	baseNodeStyleString string = "\tstyle %d %s\n"
)

type Node struct {
	ID    uint64
	Shape nodeShape
	Text  string
	Style *NodeStyle
	Class *Class
}

func NewNode(id uint64, text string) (newNode *Node) {
	newNode = &Node{
		ID:    id,
		Text:  text,
		Shape: NodeShapeRoundEdges,
	}

	return
}

func (n *Node) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf(string(baseNodeShapeString), n.ID, fmt.Sprintf(string(n.Shape), n.Text)))

	if n.Class != nil {
		sb.WriteString(fmt.Sprintf(string(baseNodeClassString), n.Class.Name))
	}

	sb.WriteByte('\n')

	if n.Style != nil {
		sb.WriteString(fmt.Sprintf(string(baseNodeStyleString), n.ID, n.Style.String()))
	}

	return sb.String()
}
