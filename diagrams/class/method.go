package class

import (
	"fmt"
	"strings"

	"github.com/TyphonHill/go-mermaid/diagrams/utils/basediagram"
)

// methodVisibility represents the access modifier for a method
type methodVisibility string

// methodClassifier represents additional modifiers for a method
type methodClassifier string

// Method visibility constants
const (
	MethodVisibilityPublic    methodVisibility = "+"
	MethodVisibilityPrivate   methodVisibility = "-"
	MethodVisibilityProtected methodVisibility = "#"
	MethodVisibilityInternal  methodVisibility = "~"
)

// Method classifier constants
const (
	MethodClassifierAbstract methodClassifier = "*"
	MethodClassifierStatic   methodClassifier = "$"
)

// Mermaid method syntax templates
const (
	baseMethodBaseString  string = basediagram.Indentation + "%s%s(%s)%s %s"
	baseMethodParamString string = "%s:%s,"
)

// Parameter represents a method parameter
type Parameter struct {
	Name string
	Type string
}

// Method represents a class method
type Method struct {
	Name       string
	Parameters []Parameter
	ReturnType string
	Visibility methodVisibility
	Classifier methodClassifier
}

// NewMethod creates a method with the given name
func NewMethod(name string) (newMethod *Method) {
	newMethod = &Method{
		Name:       name,
		Visibility: MethodVisibilityPublic,
	}
	return
}

// SetVisibility sets the method's visibility
func (m *Method) SetVisibility(visibility methodVisibility) *Method {
	m.Visibility = visibility
	return m
}

// SetReturnType sets the method's return type
func (m *Method) SetReturnType(returnType string) *Method {
	m.ReturnType = returnType
	return m
}

// SetClassifier sets the method's classifier
func (m *Method) SetClassifier(classifier methodClassifier) *Method {
	m.Classifier = classifier
	return m
}

// AddParameter adds a parameter to this method
func (m *Method) AddParameter(paramName string, paramType string) {
	m.Parameters = append(m.Parameters, Parameter{Name: paramName, Type: paramType})
}

// String returns the Mermaid syntax representation of this method
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
