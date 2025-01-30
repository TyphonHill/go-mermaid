package utils

import "testing"

func TestNewBaseDiagram(t *testing.T) {
	diagram := NewBaseDiagram()

	if diagram.Title != "" {
		t.Error("NewBaseDiagram() should not set a title")
	}
	if diagram.IsMarkdownFenceEnabled() {
		t.Error("NewBaseDiagram() should not enable markdown fence by default")
	}
}

func TestBaseDiagram_SetTitle(t *testing.T) {
	diagram := NewBaseDiagram()
	title := "Test Title"

	result := diagram.SetTitle(title)

	if result != &diagram {
		t.Error("SetTitle() should return the diagram for chaining")
	}
	if diagram.Title != title {
		t.Errorf("SetTitle() = %v, want %v", diagram.Title, title)
	}
}

func TestBaseDiagram_GetTitle(t *testing.T) {
	diagram := NewBaseDiagram()
	title := "Test Title"
	diagram.SetTitle(title)

	if got := diagram.GetTitle(); got != title {
		t.Errorf("GetTitle() = %v, want %v", got, title)
	}
}
