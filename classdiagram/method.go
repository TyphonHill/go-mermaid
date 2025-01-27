package classdiagram

import (
	"fmt"
	"strings"
)

type methodVisibility string
type methodClassifier string

const (
	MethodVisibilityPublic    methodVisibility = "+"
	MethodVisibilityPrivate   methodVisibility = "-"
	MethodVisibilityProtected methodVisibility = "#"
	MethodVisibilityInternal  methodVisibility = "~"
)

const (
	MethodClassifierAbstract methodClassifier = "*"
	MethodClassifierStatic   methodClassifier = "$"
)

const (
	baseMethodBaseString  string = "\t%s%s(%s)%s %s"
	baseMethodParamString string = "%s:%s,"
)

type Parameter struct {
	Name string
	Type string
}

type Method struct {
	Name       string
	Parameters []Parameter
	ReturnType string
	Visibility methodVisibility
	Classifier methodClassifier
}

func NewMethod(name string) (newMethod *Method) {
	newMethod = &Method{
		Name:       name,
		Visibility: MethodVisibilityPublic,
	}

	return
}

func (m *Method) AddParameter(paramName string, paramType string) {
	m.Parameters = append(m.Parameters, Parameter{Name: paramName, Type: paramType})
}

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
