package class

import (
	"os"
	"path/filepath"
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

func TestClassDiagram_EnableMarkdownFence(t *testing.T) {
	tests := []struct {
		name      string
		diagram   *ClassDiagram
		wantFence bool
	}{
		{
			name:      "Enable markdown fence",
			diagram:   NewClassDiagram(),
			wantFence: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.diagram.EnableMarkdownFence()
			if !tt.diagram.markdownFence {
				t.Error("EnableMarkdownFence() did not set markdownFence to true")
			}
		})
	}
}

func TestClassDiagram_DisableMarkdownFence(t *testing.T) {
	tests := []struct {
		name      string
		diagram   *ClassDiagram
		wantFence bool
	}{
		{
			name:      "Disable markdown fence",
			diagram:   NewClassDiagram(),
			wantFence: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// First enable it
			tt.diagram.EnableMarkdownFence()
			// Then disable it
			tt.diagram.DisableMarkdownFence()
			if tt.diagram.markdownFence {
				t.Error("DisableMarkdownFence() did not set markdownFence to false")
			}
		})
	}
}

func TestClassDiagram_RenderToFile(t *testing.T) {
	// Create temp directory for test files
	tempDir, err := os.MkdirTemp("", "class_diagram_test_*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a sample diagram
	diagram := NewClassDiagram()
	diagram.Title = "Test Class Diagram"
	namespace := diagram.AddNamespace("TestNamespace")
	class1 := diagram.AddClass("TestClass1", namespace)
	class2 := diagram.AddClass("TestClass2", namespace)
	relation := diagram.AddRelation(class1, class2)
	relation.Label = "Test Relation"

	tests := []struct {
		name           string
		filename       string
		setupFence     bool
		expectFence    bool
		expectError    bool
		validateOutput func(string) bool
	}{
		{
			name:        "Save as markdown file",
			filename:    "diagram.md",
			setupFence:  false, // Even with fencing disabled, .md should enable it
			expectFence: true,
			validateOutput: func(content string) bool {
				return strings.HasPrefix(content, "```mermaid\n") &&
					strings.HasSuffix(content, "```\n") &&
					strings.Contains(content, "Test Class Diagram") &&
					strings.Contains(content, "TestNamespace") &&
					strings.Contains(content, "TestClass1") &&
					strings.Contains(content, "Test Relation")
			},
		},
		{
			name:        "Save as text file with fencing enabled",
			filename:    "diagram.txt",
			setupFence:  true,
			expectFence: true,
			validateOutput: func(content string) bool {
				return strings.HasPrefix(content, "```mermaid\n") &&
					strings.HasSuffix(content, "```\n")
			},
		},
		{
			name:        "Save as text file without fencing",
			filename:    "diagram.txt",
			setupFence:  false,
			expectFence: false,
			validateOutput: func(content string) bool {
				return !strings.Contains(content, "```mermaid")
			},
		},
		{
			name:        "Save to nested directory",
			filename:    "nested/dir/diagram.txt",
			setupFence:  false,
			expectFence: false,
			validateOutput: func(content string) bool {
				return strings.Contains(content, "Test Class Diagram")
			},
		},
		{
			name:        "Save with invalid path",
			filename:    string([]byte{0}), // Invalid filename
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up diagram fencing
			if tt.setupFence {
				diagram.EnableMarkdownFence()
			} else {
				diagram.DisableMarkdownFence()
			}

			// Create full path
			path := filepath.Join(tempDir, tt.filename)

			// Attempt to render
			err := diagram.RenderToFile(path)

			// Check error expectation
			if tt.expectError {
				if err == nil {
					t.Error("Expected error but got none")
				}
				return
			}

			// If we don't expect an error but got one
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			// Read the file content
			content, err := os.ReadFile(path)
			if err != nil {
				t.Fatalf("Failed to read output file: %v", err)
			}

			// Validate content
			if tt.validateOutput != nil {
				if !tt.validateOutput(string(content)) {
					t.Error("Output validation failed")
				}
			}

			// Verify fence state wasn't changed permanently
			if diagram.markdownFence != tt.setupFence {
				t.Error("Diagram fence state was permanently modified")
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

	if len(namespace.classes) != 2 || len(diagram.classes) != 0 {
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
				cd.EnableMarkdownFence()
			},
			contains: []string{
				"```mermaid\n",
				"classDiagram\n",
				"direction TB\n",
				"```\n",
			},
		},
		{
			name: "Diagram with title",
			setup: func(cd *ClassDiagram) {
				cd.EnableMarkdownFence()
				cd.Title = "Test Class Diagram"
			},
			contains: []string{
				"---\ntitle: Test Class Diagram\n---\n",
				"classDiagram\n",
				"direction TB\n",
			},
		},
		{
			name: "Diagram with full components",
			setup: func(cd *ClassDiagram) {
				cd.EnableMarkdownFence()
				cd.Title = "Complex Diagram"

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
				"---\ntitle: Complex Diagram\n---\n",
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

func TestClassDiagram_SetTitle(t *testing.T) {
	diagram := NewClassDiagram()
	result := diagram.SetTitle("Test Title")

	if diagram.Title != "Test Title" {
		t.Errorf("SetTitle() = %v, want %v", diagram.Title, "Test Title")
	}

	if result != diagram {
		t.Error("SetTitle() should return diagram for chaining")
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
