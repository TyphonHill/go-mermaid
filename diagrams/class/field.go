package class

import (
	"fmt"
	"strings"
)

// fieldVisibility represents the access modifier for a field in a class diagram.
type fieldVisibility string

// fieldClassifier represents additional modifiers for a field, such as static.
type fieldClassifier string

// Field visibility constants define the access levels for class fields.
const (
	FieldVisibilityPublic    fieldVisibility = "+"
	FieldVisibilityPrivate   fieldVisibility = "-"
	FieldVisibilityProtected fieldVisibility = "#"
	FieldVisibilityInternal  fieldVisibility = "~"
)

// Field classifier constants define additional field characteristics.
const (
	FieldClassifierStatic fieldClassifier = "$"
)

// Formatting constants for field string representation.
const (
	baseFieldBaseString  string = "\t%s%s %s%s"
	baseFieldParamString string = "%s:%s,"
)

// Field represents a field (attribute) in a class diagram.
// It contains information about the field's name, type, visibility, and classifier.
type Field struct {
	Name       string
	Type       string
	Visibility fieldVisibility
	Classifier fieldClassifier
}

// NewField creates a new Field with the given name and type.
// It initializes the field with default public visibility.
func NewField(fieldName string, fieldType string) (newField *Field) {
	newField = &Field{
		Name:       fieldName,
		Type:       fieldType,
		Visibility: FieldVisibilityPublic,
	}

	return
}

// String generates the Mermaid syntax representation of the field.
// It includes the field's visibility, type, name, and classifier.
func (m *Field) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf(string(baseFieldBaseString), m.Visibility, m.Type, m.Name, m.Classifier))

	return sb.String()
}
