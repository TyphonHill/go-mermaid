package class

import (
	"fmt"
	"strings"
)

type classAnnotation string

const (
	ClassAnnotationNone        classAnnotation = ""
	ClassAnnotationInterface   classAnnotation = "<<Interface>>"
	ClassAnnotationAbstract    classAnnotation = "<<Abstract>>"
	ClassAnnotationService     classAnnotation = "<<Service>>"
	ClassAnnotationEnumeration classAnnotation = "<<Enumeration>>"
)

const (
	baseClassStartString      string = "\tclass %s%s{\n"
	baseClassEndString        string = "\t}\n"
	baseClassLabelString      string = "[\"%s\"]"
	baseClassAnnotationString string = "\t\t%s\n"
	baseClassMemberString     string = "\t%s\n"
)

// Class represents a class in the Mermaid class diagram,
// containing information about its name, label, annotation, methods, and fields.
type Class struct {
	Name       string
	Label      string
	Annotation classAnnotation
	methods    []*Method
	fields     []*Field
}

// NewClass creates a new Class with the given name.
// It initializes the class with default values and returns a pointer to the new Class.
func NewClass(name string) (newClass *Class) {
	newClass = &Class{
		Name: name,
	}

	return
}

// SetLabel sets the class label and returns the class for chaining
func (c *Class) SetLabel(label string) *Class {
	c.Label = label
	return c
}

// SetAnnotation sets the class annotation and returns the class for chaining
func (c *Class) SetAnnotation(annotation classAnnotation) *Class {
	c.Annotation = annotation
	return c
}

// AddMethod creates and adds a new method to the class
func (c *Class) AddMethod(name string) *Method {
	method := NewMethod(name)
	c.methods = append(c.methods, method)
	return method
}

// AddField creates and adds a new field to the class
func (c *Class) AddField(fieldName string, fieldType string) *Field {
	field := NewField(fieldName, fieldType)
	c.fields = append(c.fields, field)
	return field
}

// String generates the Mermaid syntax representation of the class.
// The curIndentation parameter allows for nested formatting of the class diagram.
// It includes the class name, label, annotation, fields, and methods.
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
