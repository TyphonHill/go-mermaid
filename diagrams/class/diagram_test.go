package class

import (
	"strings"
	"testing"
)

func TestNewClassDiagram(t *testing.T) {
	tests := []struct {
		name string
		want *ClassDiagram
	}{
		{
			name: "Create new class diagram with default settings",
			want: &ClassDiagram{
				Direction:  ClassDiagramDirectionTopToBottom,
				namespaces: []*Namespace{},
				notes:      []*Note{},
				classes:    []*Class{},
				relations:  []*Relation{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewClassDiagram()

			// Compare Direction
			if got.Direction != tt.want.Direction {
				t.Errorf("NewClassDiagram() Direction = %v, want %v", got.Direction, tt.want.Direction)
			}

			// Compare slice lengths
			if len(got.namespaces) != len(tt.want.namespaces) ||
				len(got.notes) != len(tt.want.notes) ||
				len(got.classes) != len(tt.want.classes) ||
				len(got.relations) != len(tt.want.relations) {
				t.Errorf("NewClassDiagram() did not initialize empty slices correctly")
			}
		})
	}
}

func TestClassDiagram_AddComponents(t *testing.T) {
	diagram := NewClassDiagram()

	// Test AddNamespace
	namespace := diagram.AddNamespace("TestNamespace")
	if namespace == nil || namespace.Name != "TestNamespace" {
		t.Errorf("AddNamespace() failed to create namespace")
	}
	if len(diagram.namespaces) != 1 {
		t.Errorf("Namespace not added to diagram")
	}

	// Test AddNote
	class := diagram.AddClass("TestClass", namespace)
	diagram.AddNote("Test Note", class)
	if len(diagram.notes) != 1 {
		t.Errorf("Note not added to diagram")
	}

	// Test AddClass with namespace
	class2 := diagram.AddClass("AnotherClass", namespace)
	if class2 == nil || class2.Name != "AnotherClass" {
		t.Errorf("AddClass() with namespace failed")
	}

	if len(namespace.Classes) != 2 || len(diagram.classes) != 0 {
		t.Errorf("Class not added to namespace correctly")
	}

	// Test AddClass without namespace
	class3 := diagram.AddClass("GlobalClass", nil)
	if class3 == nil || class3.Name != "GlobalClass" {
		t.Errorf("AddClass() without namespace failed")
	}
	if len(diagram.classes) != 1 {
		t.Errorf("Global class not added to diagram")
	}

	// Test AddRelation
	relation := diagram.AddRelation(class, class2)
	if relation == nil {
		t.Errorf("AddRelation() failed")
	}
	if len(diagram.relations) != 1 {
		t.Errorf("Relation not added to diagram")
	}
}

func TestClassDiagram_String(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(*ClassDiagram)
		contains []string
	}{
		{
			name: "Empty diagram",
			setup: func(cd *ClassDiagram) {
			},
			contains: []string{
				"classDiagram\n",
				"direction TB\n",
			},
		},
		{
			name: "Diagram with full components",
			setup: func(cd *ClassDiagram) {
				// Add namespace
				ns := cd.AddNamespace("TestNamespace")

				// Add class in namespace
				class1 := cd.AddClass("TestClass1", ns)
				class1.Annotation = ClassAnnotationService
				class1.AddField("testField", "string")
				method := class1.AddMethod("testMethod")
				method.ReturnType = "bool"

				// Add global class
				class2 := cd.AddClass("TestClass2", nil)

				// Add relation
				relation := cd.AddRelation(class1, class2)
				relation.Label = "Test Relation"

				// Add note
				cd.AddNote("Test Note", class1)
			},
			contains: []string{
				"namespace TestNamespace",
				"<<Service>>",
				"testField",
				"testMethod()",
				"TestClass2",
				"TestClass1 -- TestClass2 : Test Relation",
				"note for TestClass1 \"Test Note\"",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diagram := NewClassDiagram()

			if tt.setup != nil {
				tt.setup(diagram)
			}

			output := diagram.String()

			for _, expectedContent := range tt.contains {
				if !strings.Contains(output, expectedContent) {
					t.Errorf("String() output missing expected content: %q", expectedContent)
				}
			}

			// Verify fence markers appear together or not at all
			hasFenceStart := strings.Contains(output, "```mermaid\n")
			hasFenceEnd := strings.Contains(output, "```\n")
			if hasFenceStart != hasFenceEnd {
				t.Error("Markdown fence markers are not properly paired")
			}
		})
	}
}

func TestClassDiagram_SetDirection(t *testing.T) {
	tests := []struct {
		name      string
		direction classDiagramDirection
	}{
		{"Top to Bottom", ClassDiagramDirectionTopToBottom},
		{"Bottom Up", ClassDiagramDirectionBottomUp},
		{"Right Left", ClassDiagramDirectionRightLeft},
		{"Left Right", ClassDiagramDirectionLeftRight},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diagram := NewClassDiagram()
			result := diagram.SetDirection(tt.direction)

			if diagram.Direction != tt.direction {
				t.Errorf("SetDirection() = %v, want %v", diagram.Direction, tt.direction)
			}

			if result != diagram {
				t.Error("SetDirection() should return diagram for chaining")
			}
		})
	}
}
