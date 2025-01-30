package userjourney

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestNewDiagram(t *testing.T) {
	diagram := NewDiagram()

	if len(diagram.Sections) != 0 {
		t.Error("NewDiagram() should create empty sections slice")
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

func TestDiagram_AddSection(t *testing.T) {
	diagram := NewDiagram()
	title := "Test Section"

	section := diagram.AddSection(title)

	if len(diagram.Sections) != 1 {
		t.Error("AddSection() should add section to diagram")
	}
	if section.Title != title {
		t.Errorf("AddSection().Title = %v, want %v", section.Title, title)
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
				"journey",
			},
		},
		{
			name: "Diagram with title",
			setup: func() *Diagram {
				d := NewDiagram()
				d.SetTitle("User Shopping Journey")
				return d
			},
			contains: []string{
				"journey",
				"title User Shopping Journey",
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
				"journey",
				"```",
			},
		},
		{
			name: "Complete diagram",
			setup: func() *Diagram {
				d := NewDiagram()
				shopping := d.AddSection("Shopping")
				shopping.AddTask("Visit Store", 3)
				shopping.AddTask("Find Item", 2)
				checkout := d.AddSection("Checkout")
				checkout.AddTask("Pay", 5)
				return d
			},
			contains: []string{
				"journey",
				"section Shopping",
				"Visit Store: 3",
				"Find Item: 2",
				"section Checkout",
				"Pay: 5",
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
	// Create temp directory for test files
	tempDir, err := os.MkdirTemp("", "userjourney_test_*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	tests := []struct {
		name        string
		diagram     *Diagram
		filename    string
		wantContent string
		wantErr     bool
	}{
		{
			name: "Write markdown file",
			diagram: func() *Diagram {
				d := NewDiagram()
				d.SetTitle("Test Journey")
				section := d.AddSection("Test")
				section.AddTask("Task 1", 3)
				return d
			}(),
			filename: "test.md",
			wantContent: `journey
	title Test Journey
	section Test
		Task 1: 3
`,
			wantErr: false,
		},
		{
			name: "Write non-markdown file",
			diagram: func() *Diagram {
				d := NewDiagram()
				d.SetTitle("Test Journey")
				section := d.AddSection("Test")
				section.AddTask("Task 1", 3)
				return d
			}(),
			filename: "test.txt",
			wantContent: `journey
	title Test Journey
	section Test
		Task 1: 3
`,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := filepath.Join(tempDir, tt.filename)

			err := tt.diagram.RenderToFile(path)
			if (err != nil) != tt.wantErr {
				t.Errorf("RenderToFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				content, err := os.ReadFile(path)
				if err != nil {
					t.Fatalf("Failed to read rendered file: %v", err)
				}

				// For .md files, content should be wrapped in markdown fence
				wantContent := tt.wantContent
				if strings.HasSuffix(path, ".md") {
					wantContent = "```mermaid\n" + tt.wantContent + "```\n"
				}

				if got := string(content); got != wantContent {
					t.Errorf("RenderToFile() content mismatch:\nwant:\n%s\ngot:\n%s", wantContent, got)
				}
			}
		})
	}
}
