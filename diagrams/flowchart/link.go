package flowchart

import (
	"fmt"
	"strings"

	"github.com/TyphonHill/go-mermaid/diagrams/utils/basediagram"
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
	baseLinkString     string = basediagram.Indentation + "%s %s%s%s%s %s\n"
	baseLinkTextString string = "|%s|"
)

// Link represents a connection between nodes in a flowchart
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

// SetText sets the link text and returns the link for chaining
func (l *Link) SetText(text string) *Link {
	l.Text = text
	return l
}

// SetShape sets the link shape and returns the link for chaining
func (l *Link) SetShape(shape linkShape) *Link {
	l.Shape = shape
	return l
}

// SetLength sets the link length and returns the link for chaining
func (l *Link) SetLength(length int) *Link {
	l.Length = length
	return l
}

// SetHead sets the head arrow type and returns the link for chaining
func (l *Link) SetHead(arrowType linkArrowType) *Link {
	l.Head = arrowType
	return l
}

// SetTail sets the tail arrow type and returns the link for chaining
func (l *Link) SetTail(arrowType linkArrowType) *Link {
	l.Tail = arrowType
	return l
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
