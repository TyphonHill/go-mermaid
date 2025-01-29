package entityrelationship

import (
	"os"
	"strings"
	"testing"
)

func TestNewDiagram(t *testing.T) {
	diagram := NewDiagram()

	if len(diagram.Entities) != 0 {
		t.Error("NewDiagram() should create empty entities slice")
	}
	if len(diagram.Relationships) != 0 {
		t.Error("NewDiagram() should create empty relationships slice")
	}
	if diagram.Title != "" {
		t.Error("NewDiagram() should not set a title")
	}
	if diagram.markdownFence {
		t.Error("NewDiagram() should not enable markdown fence by default")
	}
}

func TestDiagram_EnableMarkdownFence(t *testing.T) {
	diagram := NewDiagram()

	result := diagram.EnableMarkdownFence()

	if result != diagram {
		t.Error("EnableMarkdownFence() should return the diagram for chaining")
	}
	if !diagram.markdownFence {
		t.Error("EnableMarkdownFence() should enable markdown fence")
	}
}

func TestDiagram_DisableMarkdownFence(t *testing.T) {
	diagram := NewDiagram()
	diagram.EnableMarkdownFence()

	diagram.DisableMarkdownFence()

	if diagram.markdownFence {
		t.Error("DisableMarkdownFence() should disable markdown fence")
	}
}

func TestDiagram_SetTitle(t *testing.T) {
	diagram := NewDiagram()
	title := "Test Title"

	result := diagram.SetTitle(title)

	if result != diagram {
		t.Error("SetTitle() should return the diagram for chaining")
	}
	if diagram.Title != title {
		t.Errorf("SetTitle() = %v, want %v", diagram.Title, title)
	}
}

func TestDiagram_AddEntity(t *testing.T) {
	diagram := NewDiagram()
	name := "TEST"

	entity := diagram.AddEntity(name)

	if len(diagram.Entities) != 1 {
		t.Error("AddEntity() should add entity to diagram")
	}
	if entity.Name != name {
		t.Errorf("AddEntity().Name = %v, want %v", entity.Name, name)
	}
}

func TestDiagram_AddRelationship(t *testing.T) {
	diagram := NewDiagram()
	from := diagram.AddEntity("A")
	to := diagram.AddEntity("B")

	rel := diagram.AddRelationship(from, to)

	if len(diagram.Relationships) != 1 {
		t.Error("AddRelationship() should add relationship to diagram")
	}
	if rel.From != from {
		t.Errorf("AddRelationship().From = %v, want %v", rel.From, from)
	}
	if rel.To != to {
		t.Errorf("AddRelationship().To = %v, want %v", rel.To, to)
	}
}

func TestDiagram_String(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() *Diagram
		contains []string
	}{
		{
			name: "Empty diagram",
			setup: func() *Diagram {
				return NewDiagram()
			},
			contains: []string{
				"erDiagram",
			},
		},
		{
			name: "Diagram with title",
			setup: func() *Diagram {
				d := NewDiagram()
				d.SetTitle("Test ERD")
				return d
			},
			contains: []string{
				"title: Test ERD",
				"erDiagram",
			},
		},
		{
			name: "Diagram with markdown fence",
			setup: func() *Diagram {
				d := NewDiagram()
				d.EnableMarkdownFence()
				return d
			},
			contains: []string{
				"```mermaid",
				"erDiagram",
				"```",
			},
		},
		{
			name: "Complete diagram",
			setup: func() *Diagram {
				d := NewDiagram()
				user := d.AddEntity("USER")
				post := d.AddEntity("POST")
				user.AddAttribute("id", TypeInteger).SetPrimaryKey()
				post.AddAttribute("id", TypeInteger).SetPrimaryKey()
				d.AddRelationship(user, post).
					SetLabel("writes").
					SetCardinality(OneToZeroOrMore)
				return d
			},
			contains: []string{
				"erDiagram",
				"USER {",
				"POST {",
				"writes",
				string(OneToZeroOrMore),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diagram := tt.setup()
			result := diagram.String()

			for _, want := range tt.contains {
				if !strings.Contains(result, want) {
					t.Errorf("String() missing expected content %q in:\n%s", want, result)
				}
			}
		})
	}
}

func TestDiagram_RenderToFile(t *testing.T) {
	// Create temp test file
	f, err := os.CreateTemp("", "test-*.md")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	testFile := f.Name()
	f.Close()

	// Clean up after test
	defer os.Remove(testFile)

	tests := []struct {
		name        string
		diagram     *Diagram
		wantErr     bool
		wantContent []string
	}{
		{
			name: "Write markdown with fence",
			diagram: func() *Diagram {
				d := NewDiagram()
				d.SetTitle("Test ERD")
				d.EnableMarkdownFence()
				user := d.AddEntity("USER")
				user.AddAttribute("id", TypeInteger).SetPrimaryKey()
				return d
			}(),
			wantErr: false,
			wantContent: []string{
				"```mermaid",
				"title: Test ERD",
				"erDiagram",
				"USER {",
				"int id PK",
				"}",
				"```",
			},
		},
		{
			name: "Write without fence",
			diagram: func() *Diagram {
				d := NewDiagram()
				d.SetTitle("Test ERD")
				user := d.AddEntity("USER")
				user.AddAttribute("id", TypeInteger).SetPrimaryKey()
				return d
			}(),
			wantErr: false,
			wantContent: []string{
				"title: Test ERD",
				"erDiagram",
				"USER {",
				"int id PK",
				"}",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.diagram.RenderToFile(testFile)

			if (err != nil) != tt.wantErr {
				t.Errorf("RenderToFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				content, err := os.ReadFile(testFile)
				if err != nil {
					t.Fatalf("Failed to read rendered file: %v", err)
				}

				gotContent := string(content)
				for _, want := range tt.wantContent {
					if !strings.Contains(gotContent, want) {
						t.Errorf("Content missing %q in:\n%s", want, gotContent)
					}
				}
			}
		})
	}
}
