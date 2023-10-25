package flowchart

import (
	"fmt"
	"strings"
)

type linkShape string
type linkArrowType string

const (
	LinkShapeOpen      linkShape = "--%s"
	LinkShapeDotted    linkShape = "-.%s-"
	LinkShapeThick     linkShape = "==%s"
	LinkShapeInvisible linkShape = "~~%s"
)

const (
	LinkArrowTypeNone      linkArrowType = ""
	LinkArrowTypeArrow     linkArrowType = ">"
	LinkArrowTypeLeftArrow linkArrowType = "<"
	LinkArrowTypeBullet    linkArrowType = "o"
	LinkArrowTypeCross     linkArrowType = "x"
)

const (
	baseLinkString     string = "\t%d %s%s%s%s %d"
	baseLinkTextString string = "|%s|"
)

type Link struct {
	Shape  linkShape
	Head   linkArrowType
	Tail   linkArrowType
	Text   string
	From   *Node
	To     *Node
	Length int
}

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
