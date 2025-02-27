package basediagram

import "testing"

func TestNewBaseDiagram(t *testing.T) {
	diagram := NewBaseDiagram()

	if diagram.title != "" {
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
	if diagram.title != title {
		t.Errorf("SetTitle() = %v, want %v", diagram.title, title)
	}
}
