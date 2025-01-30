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
	section := NewSection("Test")
	timePeriod := "2024-01"
	eventText := "Test Event"

	event := section.AddEvent(timePeriod, eventText)

	if event.TimePeriod != timePeriod {
		t.Errorf("AddEvent().TimePeriod = %v, want %v", event.TimePeriod, timePeriod)
	}
	if event.Text != eventText {
		t.Errorf("AddEvent().Text = %v, want %v", event.Text, eventText)
	}
	if len(section.Events) != 1 {
		t.Error("AddEvent() should add event to section")
	}
}

func TestSection_AddSubEvent(t *testing.T) {
	section := NewSection("Test")
	eventText := "Sub Event"

	event := section.AddSubEvent(eventText)

	if event.TimePeriod != "" {
		t.Errorf("AddSubEvent().TimePeriod = %v, want empty string", event.TimePeriod)
	}
	if event.Text != eventText {
		t.Errorf("AddSubEvent().Text = %v, want %v", event.Text, eventText)
	}
	if len(section.Events) != 1 {
		t.Error("AddSubEvent() should add event to section")
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
				s.AddSubEvent("Sub Event")
				return s
			},
			contains: []string{
				"section Test Section",
				"2024-01 : First Event",
				": Sub Event",
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
