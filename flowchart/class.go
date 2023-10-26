package flowchart

import (
	"fmt"
	"strings"
)

const (
	baseClassString string = "\tclassDef %s %s\n"
)

// Classes are a convenient way of creating a node style since you can attach them directly to a node.
// Reference: https://mermaid.js.org/syntax/flowchart.html#classes
type Class struct {
	Name  string
	Style *NodeStyle
}

// Creates a new Class and sets default values to some attributes
func NewClass(name string) (newClass *Class) {
	newClass = &Class{
		Name:  name,
		Style: NewNodeStyle(),
	}

	return
}

// Builds a new string based on the current elements
func (c *Class) String() string {
	var sb strings.Builder

	if c.Style != nil {
		sb.WriteString(fmt.Sprintf(string(baseClassString), c.Name, c.Style.String()))
	}

	return sb.String()
}
