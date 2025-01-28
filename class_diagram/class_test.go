package classdiagram

import (
	"strings"
	"testing"
)

func TestNewClass(t *testing.T) {
	tests := []struct {
		name string
		want *Class
	}{
		{
			name: "Create new class",
			want: &Class{
				Name:       "TestClass",
				Label:      "",
				Annotation: ClassAnnotationNone,
				methods:    []*Method{},
				fields:     []*Field{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewClass(tt.want.Name)

			if got.Name != tt.want.Name {
				t.Errorf("NewClass() Name = %v, want %v", got.Name, tt.want.Name)
			}

			// Check that methods and fields are initialized as empty slices
			if len(got.methods) != 0 || len(got.fields) != 0 {
				t.Errorf("NewClass() did not initialize empty methods and fields slices")
			}
		})
	}
}

func TestClass_AddMethod(t *testing.T) {
	class := NewClass("TestClass")

	method := class.AddMethod("testMethod")

	if method == nil {
		t.Errorf("AddMethod() returned nil")
	}

	if method.Name != "testMethod" {
		t.Errorf("AddMethod() method name = %v, want %v", method.Name, "testMethod")
	}

	if len(class.methods) != 1 {
		t.Errorf("AddMethod() did not add method to class methods")
	}
}

func TestClass_AddField(t *testing.T) {
	class := NewClass("TestClass")

	field := class.AddField("testField", "string")

	if field == nil {
		t.Errorf("AddField() returned nil")
	}

	if field.Name != "testField" || field.Type != "string" {
		t.Errorf("AddField() field = %v, type = %v, want name %v, type %v",
			field.Name, field.Type, "testField", "string")
	}

	if len(class.fields) != 1 {
		t.Errorf("AddField() did not add field to class fields")
	}
}

func TestClass_String(t *testing.T) {
	tests := []struct {
		name        string
		class       *Class
		indentation string
		contains    []string
	}{
		{
			name:        "Empty class",
			class:       NewClass("EmptyClass"),
			indentation: "%s",
			contains: []string{
				"class EmptyClass{",
				"}",
			},
		},
		{
			name: "Class with annotation",
			class: func() *Class {
				c := NewClass("ServiceClass")
				c.Annotation = ClassAnnotationService
				return c
			}(),
			indentation: "%s",
			contains: []string{
				"class ServiceClass{",
				"<<Service>>",
				"}",
			},
		},
		{
			name: "Class with label",
			class: func() *Class {
				c := NewClass("LabeledClass")
				c.Label = "Custom Label"
				return c
			}(),
			indentation: "%s",
			contains: []string{
				"class LabeledClass[\"Custom Label\"]{",
				"}",
			},
		},
		{
			name: "Class with fields and methods",
			class: func() *Class {
				c := NewClass("ComplexClass")

				// Add fields
				field1 := c.AddField("name", "string")
				field1.Visibility = FieldVisibilityPrivate

				field2 := c.AddField("age", "int")
				field2.Classifier = FieldClassifierStatic

				// Add methods
				method1 := c.AddMethod("getName")
				method1.ReturnType = "string"
				method1.Visibility = MethodVisibilityPublic

				method2 := c.AddMethod("setName")
				method2.AddParameter("newName", "string")
				method2.Visibility = MethodVisibilityPrivate
				method2.Classifier = MethodClassifierStatic

				return c
			}(),
			indentation: "%s",
			contains: []string{
				"class ComplexClass{",
				"-string name",
				"+int age$",
				"+getName() string",
				"-setName(newName:string)$",
				"}",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := tt.class.String(tt.indentation)

			for _, expectedContent := range tt.contains {
				if !strings.Contains(output, expectedContent) {
					t.Errorf("String() output missing expected content: %q, got %v", expectedContent, output)
				}
			}
		})
	}
}
