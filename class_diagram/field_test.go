package classdiagram

import (
	"strings"
	"testing"
)

func TestNewField(t *testing.T) {
	tests := []struct {
		name      string
		fieldName string
		fieldType string
		want      *Field
	}{
		{
			name:      "Create new field",
			fieldName: "testField",
			fieldType: "string",
			want: &Field{
				Name:       "testField",
				Type:       "string",
				Visibility: FieldVisibilityPublic,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewField(tt.fieldName, tt.fieldType)

			if got.Name != tt.want.Name {
				t.Errorf("NewField() Name = %v, want %v", got.Name, tt.want.Name)
			}

			if got.Type != tt.want.Type {
				t.Errorf("NewField() Type = %v, want %v", got.Type, tt.want.Type)
			}

			if got.Visibility != tt.want.Visibility {
				t.Errorf("NewField() Visibility = %v, want %v", got.Visibility, tt.want.Visibility)
			}
		})
	}
}

func TestField_String(t *testing.T) {
	tests := []struct {
		name     string
		field    *Field
		contains []string
	}{
		{
			name:  "Public field with default visibility",
			field: NewField("name", "string"),
			contains: []string{
				"+string name",
			},
		},
		{
			name: "Private field",
			field: func() *Field {
				f := NewField("age", "int")
				f.Visibility = FieldVisibilityPrivate
				return f
			}(),
			contains: []string{
				"-int age",
			},
		},
		{
			name: "Protected field",
			field: func() *Field {
				f := NewField("data", "object")
				f.Visibility = FieldVisibilityProtected
				return f
			}(),
			contains: []string{
				"#object data",
			},
		},
		{
			name: "Internal field",
			field: func() *Field {
				f := NewField("internal", "bool")
				f.Visibility = FieldVisibilityInternal
				return f
			}(),
			contains: []string{
				"~bool internal",
			},
		},
		{
			name: "Static field",
			field: func() *Field {
				f := NewField("count", "int")
				f.Classifier = FieldClassifierStatic
				return f
			}(),
			contains: []string{
				"+int count$",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := tt.field.String()

			for _, expectedContent := range tt.contains {
				if !strings.Contains(output, expectedContent) {
					t.Errorf("String() output missing expected content: %q", expectedContent)
				}
			}
		})
	}
}
