package flowchart

import (
	"fmt"
	"strings"
)

type linkShape string
type linkArrowType string

// List of possible Link shapes.
// Reference: https://mermaid.js.org/syntax/flowchart.html#links-between-nodes
const (
	LinkShapeOpen      linkShape = "--%s"
	LinkShapeDotted    linkShape = "-.%s-"
	LinkShapeThick     linkShape = "==%s"
	LinkShapeInvisible linkShape = "~~%s"
)

// List of possible Link arrow types.
// Reference: https://mermaid.js.org/syntax/flowchart.html#links-between-nodes
const (
	LinkArrowTypeNone      linkArrowType = ""
	LinkArrowTypeArrow     linkArrowType = ">"
	LinkArrowTypeLeftArrow linkArrowType = "<"
	LinkArrowTypeBullet    linkArrowType = "o"
	LinkArrowTypeCross     linkArrowType = "x"
)

const (
	baseLinkString     string = "\t%d %s%s%s%s %d\n"
	baseLinkTextString string = "|%s|"
)

// Nodes can be connected with links/edges. It is possible to have different types of links or attach a text string to a link.
// Reference: https://mermaid.js.org/syntax/flowchart.html#links-between-nodes
type Link struct {
	Shape  linkShape
	Head   linkArrowType
	Tail   linkArrowType
	Text   string
	From   *Node
	To     *Node
	Length int
}

// NewLink creates a new Link and sets default values to some attributes
func NewLink(from *Node, to *Node) (newLink *Link) {
	newLink = &Link{
		From:   from,
		To:     to,
		Shape:  LinkShapeOpen,
		Head:   LinkArrowTypeArrow,
		Tail:   LinkArrowTypeNone,
		Length: 0,
	}

	return
}

// String generates a Mermaid string representation of the link,
// including its shape, arrow types, text, and length.
func (l *Link) String() string {
	var sb strings.Builder

	extension := ""
	for i := 0; i < l.Length; i++ {
		extension += string(l.Shape[1])
	}

	text := ""
	if len(l.Text) > 0 {
		text = fmt.Sprintf(string(baseLinkTextString), l.Text)
	}

	sb.WriteString(fmt.Sprintf(string(baseLinkString), l.From.ID, string(l.Tail), fmt.Sprintf(string(l.Shape), extension), string(l.Head), text, l.To.ID))

	return sb.String()
}
