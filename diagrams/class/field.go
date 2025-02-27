package class

import (
	"fmt"
	"strings"

	"github.com/TyphonHill/go-mermaid/diagrams/utils/basediagram"
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
	baseFieldBaseString  string = basediagram.Indentation + "%s%s %s%s"
	baseFieldParamString string = "%s:%s,"
)

// Field represents a class field with visibility and type information
type Field struct {
	Name       string
	Type       string
	Visibility fieldVisibility
	Classifier fieldClassifier
}

// NewField creates a field with the given name and type
func NewField(name string, fieldType string) *Field {
	return &Field{
		Name:       name,
		Type:       fieldType,
		Visibility: FieldVisibilityPublic,
	}
}

// SetVisibility sets the field's visibility
func (f *Field) SetVisibility(visibility fieldVisibility) *Field {
	f.Visibility = visibility
	return f
}

// String returns the Mermaid syntax representation of this field
func (f *Field) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf(string(baseFieldBaseString), f.Visibility, f.Type, f.Name, f.Classifier))

	return sb.String()
}
