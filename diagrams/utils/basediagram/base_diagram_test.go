package basediagram

import "testing"

type TestConfig struct{}

func (c TestConfig) String() string { return "" }

func TestNewBaseDiagram(t *testing.T) {
	diagram := NewBaseDiagram(TestConfig{})

	if diagram.Title != "" {
		t.Error("NewBaseDiagram() should not set a title")
	}
}

func TestBaseDiagram_SetTitle(t *testing.T) {
	diagram := NewBaseDiagram(TestConfig{})
	title := "Test Title"

	result := diagram.SetTitle(title)

	if result != &diagram {
		t.Error("SetTitle() should return the diagram for chaining")
	}
	if diagram.Title != title {
		t.Errorf("SetTitle() = %v, want %v", diagram.Title, title)
	}
}
