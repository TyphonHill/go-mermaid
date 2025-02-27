// Package class provides functionality for creating Mermaid class diagrams
package class

import (
	"fmt"
	"strings"

	"github.com/TyphonHill/go-mermaid/diagrams/utils/basediagram"
)

type classAnnotation string

// Available class annotations
const (
	ClassAnnotationNone        classAnnotation = ""
	ClassAnnotationInterface   classAnnotation = "<<Interface>>"
	ClassAnnotationAbstract    classAnnotation = "<<Abstract>>"
	ClassAnnotationService     classAnnotation = "<<Service>>"
	ClassAnnotationEnumeration classAnnotation = "<<Enumeration>>"
)

// Mermaid class syntax templates
const (
	baseClassStartString      string = basediagram.Indentation + "class %s%s{\n"
	baseClassEndString        string = basediagram.Indentation + "}\n"
	baseClassLabelString      string = "[\"%s\"]"
	baseClassAnnotationString string = basediagram.Indentation + "%s\n"
	baseClassMemberString     string = basediagram.Indentation + "%s\n"
)

// Class represents a class in a Mermaid class diagram
type Class struct {
	Name       string
	Label      string
	Annotation classAnnotation
	methods    []*Method
	fields     []*Field
}

// NewClass creates a new Class with the given name
func NewClass(name string) (newClass *Class) {
	newClass = &Class{
		Name: name,
	}
	return
}

// SetLabel sets the class label
func (c *Class) SetLabel(label string) *Class {
	c.Label = label
	return c
}

// SetAnnotation sets the class annotation
func (c *Class) SetAnnotation(annotation classAnnotation) *Class {
	c.Annotation = annotation
	return c
}

// AddMethod creates and adds a new method
func (c *Class) AddMethod(name string) *Method {
	method := NewMethod(name)
	c.methods = append(c.methods, method)
	return method
}

// AddField creates and adds a new field
func (c *Class) AddField(fieldName string, fieldType string) *Field {
	field := NewField(fieldName, fieldType)
	c.fields = append(c.fields, field)
	return field
}

// String returns the Mermaid syntax representation of this class
func (c *Class) String(curIndentation string) string {
	var sb strings.Builder

	label := ""
	if len(c.Label) > 0 {
		label = fmt.Sprintf(string(baseClassLabelString), c.Label)
	}

	sb.WriteString(fmt.Sprintf(curIndentation, fmt.Sprintf(string(baseClassStartString), c.Name, label)))

	if c.Annotation != ClassAnnotationNone {
		sb.WriteString(fmt.Sprintf(curIndentation, fmt.Sprintf(string(baseClassAnnotationString), string(c.Annotation))))
	}

	for _, field := range c.fields {
		sb.WriteString(fmt.Sprintf(curIndentation, fmt.Sprintf(string(baseClassMemberString), field.String())))
	}

	for _, method := range c.methods {
		sb.WriteString(fmt.Sprintf(curIndentation, fmt.Sprintf(string(baseClassMemberString), method.String())))
	}

	sb.WriteString(fmt.Sprintf(curIndentation, baseClassEndString))

	return sb.String()
}
