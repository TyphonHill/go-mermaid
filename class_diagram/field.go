package class_diagram

import (
	"fmt"
	"strings"
)

type fieldVisibility string
type fieldClassifier string

const (
	FieldVisibilityPublic    fieldVisibility = "+"
	FieldVisibilityPrivate   fieldVisibility = "-"
	FieldVisibilityProtected fieldVisibility = "#"
	FieldVisibilityInternal  fieldVisibility = "~"
)

const (
	FieldClassifierStatic fieldClassifier = "$"
)

const (
	baseFieldBaseString  string = "\t%s%s %s%s"
	baseFieldParamString string = "%s:%s,"
)

type Field struct {
	Name       string
	Type       string
	Visibility fieldVisibility
	Classifier fieldClassifier
}

func NewField(fieldName string, fieldType string) (newField *Field) {
	newField = &Field{
		Name:       fieldName,
		Type:       fieldType,
		Visibility: FieldVisibilityPublic,
	}

	return
}

func (m *Field) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf(string(baseFieldBaseString), m.Visibility, m.Type, m.Name, m.Classifier))

	return sb.String()
}
