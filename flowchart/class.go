package flowchart

import (
	"fmt"
	"strings"
)

const (
	baseClassString string = "\tclassDef %s %s\n"
)

type Class struct {
	Name  string
	Style *NodeStyle
}

func NewClass(name string) (newClass *Class) {
	newClass = &Class{
		Name: name,
	}

	return
}

func (c *Class) String() string {
	var sb strings.Builder

	if c.Style != nil {
		sb.WriteString(fmt.Sprintf(string(baseClassString), c.Name, c.Style.String()))
	}

	return sb.String()
}
