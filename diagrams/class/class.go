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

// AddMethod creates and adds a new method to the class.
// It returns the newly created Method, allowing for further configuration.
func (c *Class) AddMethod(name string) (newMethod *Method) {
	newMethod = NewMethod(name)

	c.methods = append(c.methods, newMethod)

	return
}

// AddField creates and adds a new field to the class.
// It returns the newly created Field, allowing for further configuration.
func (c *Class) AddField(fieldName string, fieldType string) (newField *Field) {
	newField = NewField(fieldName, fieldType)

	c.fields = append(c.fields, newField)

	return
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
