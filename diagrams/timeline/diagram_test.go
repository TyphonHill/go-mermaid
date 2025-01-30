package timeline

import (
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
	if diagram.IsMarkdownFenceEnabled() {
		t.Error("NewDiagram() should not enable markdown fence by default")
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
				"timeline",
			},
		},
		{
			name: "Diagram with title",
			setup: func() *Diagram {
				d := NewDiagram()
				d.SetTitle("Test Timeline")
				return d
			},
			contains: []string{
				"timeline",
				"title Test Timeline",
			},
		},
		{
			name: "Complete diagram",
			setup: func() *Diagram {
				d := NewDiagram()
				d.SetTitle("Project Timeline")

				planning := d.AddSection("Planning")
				planning.AddEvent("2024-01", "Project kickoff")
				planning.AddSubEvent("Requirements gathering")

				dev := d.AddSection("Development")
				dev.AddEvent("2024-02", "Implementation")
				dev.AddSubEvent("Testing")

				return d
			},
			contains: []string{
				"timeline",
				"title Project Timeline",
				"section Planning",
				"2024-01 : Project kickoff",
				": Requirements gathering",
				"section Development",
				"2024-02 : Implementation",
				": Testing",
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
