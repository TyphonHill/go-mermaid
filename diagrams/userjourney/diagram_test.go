package userjourney

import (
	"strings"
	"testing"
)

func TestNewDiagram(t *testing.T) {
	diagram := NewDiagram()

	if len(diagram.Sections) != 0 {
		t.Error("NewDiagram() should create empty sections slice")
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
