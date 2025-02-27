package timeline

import (
	"strings"
	"testing"
)

func TestNewSection(t *testing.T) {
	title := "Test Section"
	section := NewSection(title)

	if section.Title != title {
		t.Errorf("NewSection().Title = %v, want %v", section.Title, title)
	}
	if len(section.Events) != 0 {
		t.Error("NewSection() should create empty events slice")
	}
}

func TestSection_AddEvent(t *testing.T) {
	tests := []struct {
		name       string
		title      string
		text       string
		isFirst    bool
		wantString string
	}{
		{
			name:       "First event with title",
			title:      "2024-01",
			text:       "Test Event",
			isFirst:    true,
			wantString: "\t\t2024-01\n\t\t: Test Event",
		},
		{
			name:       "First event without title",
			title:      "",
			text:       "Test Event",
			isFirst:    true,
			wantString: "Test Event", // No colon for first event without title
		},
		{
			name:       "Second event without title",
			title:      "",
			text:       "Test Event",
			isFirst:    false,
			wantString: ": Test Event", // Has colon for non-first events
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			section := NewSection("Test")
			if !tt.isFirst {
				section.AddEvent("2024-01", "Previous Event") // Add a first event
			}
			section.AddEvent(tt.title, tt.text)

			result := section.String()
			if !strings.Contains(result, tt.wantString) {
				t.Errorf("AddEvent() result = %v, want string containing %v", result, tt.wantString)
			}
		})
	}
}

func TestSection_String(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() *Section
		contains []string
	}{
		{
			name: "Section with events",
			setup: func() *Section {
				s := NewSection("Test Section")
				s.AddEvent("2024-01", "First Event")
				return s
			},
			contains: []string{
				"section Test Section",
				"\t\t2024-01\n\t\t: First Event",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			section := tt.setup()
			result := section.String()

			for _, want := range tt.contains {
				if !strings.Contains(result, want) {
					t.Errorf("String() missing expected content %q in:\n%s", want, result)
				}
			}
		})
	}
}
