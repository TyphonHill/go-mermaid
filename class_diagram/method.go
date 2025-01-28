package classdiagram

import (
	"fmt"
	"strings"
)

// methodVisibility represents the access modifier for a method in a class diagram.
type methodVisibility string

// methodClassifier represents additional modifiers for a method, such as abstract or static.
type methodClassifier string

// Method visibility constants define the access levels for class methods.
const (
	MethodVisibilityPublic    methodVisibility = "+"
	MethodVisibilityPrivate   methodVisibility = "-"
	MethodVisibilityProtected methodVisibility = "#"
	MethodVisibilityInternal  methodVisibility = "~"
)

// Method classifier constants define additional method characteristics.
const (
	MethodClassifierAbstract methodClassifier = "*"
	MethodClassifierStatic   methodClassifier = "$"
)

// Formatting constants for method string representation.
const (
	baseMethodBaseString  string = "\t%s%s(%s)%s %s"
	baseMethodParamString string = "%s:%s,"
)

// Parameter represents a single parameter in a method signature.
type Parameter struct {
	Name string
	Type string
}

// Method represents a method in a class diagram.
// It contains information about the method's name, parameters, return type,
// visibility, and classifier.
type Method struct {
	Name       string
	Parameters []Parameter
	ReturnType string
	Visibility methodVisibility
	Classifier methodClassifier
}

// NewMethod creates a new Method with the given name.
// It initializes the method with default public visibility.
func NewMethod(name string) (newMethod *Method) {
	newMethod = &Method{
		Name:       name,
		Visibility: MethodVisibilityPublic,
	}

	return
}

// AddParameter adds a new parameter to the method's parameter list.
func (m *Method) AddParameter(paramName string, paramType string) {
	m.Parameters = append(m.Parameters, Parameter{Name: paramName, Type: paramType})
}

// String generates the Mermaid syntax representation of the method.
// It includes the method's visibility, name, parameters, classifier, and return type.
func (m *Method) String() string {
	var sb strings.Builder

	params := ""
	for _, param := range m.Parameters {
		params += fmt.Sprintf(string(baseMethodParamString), param.Name, param.Type)
	}
	params = strings.Trim(params, ",")

	sb.WriteString(fmt.Sprintf(string(baseMethodBaseString), m.Visibility, m.Name, params, m.Classifier, m.ReturnType))

	return sb.String()
}
