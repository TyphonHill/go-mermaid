package class

import (
	"strings"
	"testing"
)

func TestNewNamespace(t *testing.T) {
	tests := []struct {
		name string
		want *Namespace
	}{
		{
			name: "Create new namespace",
			want: &Namespace{
				Name:    "TestNamespace",
				Classes: []*Class{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewNamespace(tt.want.Name)

			if got.Name != tt.want.Name {
				t.Errorf("NewNamespace() Name = %v, want %v", got.Name, tt.want.Name)
			}

			// Check that classes are initialized as empty slice
			if len(got.Classes) != 0 {
				t.Errorf("NewNamespace() did not initialize empty classes slice")
			}
		})
	}
}

func TestNamespace_AddClass(t *testing.T) {
	namespace := NewNamespace("TestNamespace")

	class := NewClass("TestClass")
	namespace.AddClass(class)

	if class == nil {
		t.Errorf("AddClass() returned nil")
	}

	if class.Name != "TestClass" {
		t.Errorf("AddClass() class name = %v, want %v", class.Name, "TestClass")
	}

	if len(namespace.Classes) != 1 {
		t.Errorf("AddClass() did not add class to namespace classes")
	}

	// Add multiple classes
	namespace.AddClass(NewClass("AnotherClass"))
	namespace.AddClass(NewClass("ThirdClass"))

	if len(namespace.Classes) != 3 {
		t.Errorf("AddClass() did not correctly add multiple classes")
	}
}

func TestNamespace_String(t *testing.T) {
	tests := []struct {
		name        string
		namespace   *Namespace
		contains    []string
		notContains []string
	}{
		{
			name:      "Empty namespace",
			namespace: NewNamespace("EmptyNamespace"),
			contains:  []string{}, // Empty namespace should return empty string
			notContains: []string{
				"namespace EmptyNamespace",
				"}",
			},
		},
		{
			name: "Namespace with single class",
			namespace: func() *Namespace {
				ns := NewNamespace("TestNamespace")
				class := NewClass("TestClass")
				ns.AddClass(class)
				class.Annotation = ClassAnnotationService
				class.AddField("testField", "string")
				return ns
			}(),
			contains: []string{
				"namespace TestNamespace{",
				"class TestClass{",
				"<<Service>>",
				"+string testField",
				"}",
				"}",
			},
		},
		{
			name: "Namespace with multiple classes",
			namespace: func() *Namespace {
				ns := NewNamespace("MultiClassNamespace")
				class1 := NewClass("FirstClass")
				ns.AddClass(class1)
				class1.AddField("id", "int")

				class2 := NewClass("SecondClass")
				ns.AddClass(class2)
				class2.Annotation = ClassAnnotationInterface
				class2.AddMethod("doSomething")
				return ns
			}(),
			contains: []string{
				"namespace MultiClassNamespace{\n",
				"class FirstClass{",
				"+int id",
				"class SecondClass{",
				"<<Interface>>",
				"+doSomething()",
				"}",
				"}",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := tt.namespace.String("")

			// Check contains
			for _, expectedContent := range tt.contains {
				if !strings.Contains(output, expectedContent) {
					t.Errorf("String() output missing expected content: %q", expectedContent)
				}
			}

			// Check not contains
			for _, unexpectedContent := range tt.notContains {
				if strings.Contains(output, unexpectedContent) {
					t.Errorf("String() output should not contain: %q", unexpectedContent)
				}
			}
		})
	}
}
