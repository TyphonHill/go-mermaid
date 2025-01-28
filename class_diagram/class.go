package class_diagram

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

type Class struct {
	Name       string
	Label      string
	Annotation classAnnotation
	methods    []*Method
	fields     []*Field
}

// Creates a new Class and sets default values to some attributes
func NewClass(name string) (newClass *Class) {
	newClass = &Class{
		Name: name,
	}

	return
}

func (c *Class) AddMethod(name string) (newMethod *Method) {
	newMethod = NewMethod(name)

	c.methods = append(c.methods, newMethod)

	return
}

func (c *Class) AddField(fieldName string, fieldType string) (newField *Field) {
	newField = NewField(fieldName, fieldType)

	c.fields = append(c.fields, newField)

	return
}

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
