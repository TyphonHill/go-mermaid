package timeline

import (
	"strings"
	"testing"
)

func TestNewEvent(t *testing.T) {
	timePeriod := "2024-01"
	text := "Test Event"
	event := NewEvent(timePeriod, text)

	if event.Title != timePeriod {
		t.Errorf("NewEvent().TimePeriod = %v, want %v", event.Title, timePeriod)
	}
	if event.Text != text {
		t.Errorf("NewEvent().Text = %v, want %v", event.Text, text)
	}
}

func TestEvent_String(t *testing.T) {
	tests := []struct {
		name       string
		timePeriod string
		text       string
		isFirst    bool
		subEvents  []string
		contains   []string
	}{
		{
			name:       "Event with time period",
			timePeriod: "2024-01",
			text:       "Test Event",
			contains: []string{
				"2024-01",
				"Test Event",
			},
		},
		{
			name:       "Regular sub-event without time period",
			timePeriod: "",
			text:       "Sub Event",
			contains: []string{
				"Sub Event",
			},
		},
		{
			name:       "First sub-event without time period",
			timePeriod: "",
			text:       "First Sub Event",
			contains: []string{
				"First Sub Event",
			},
		},
		{
			name:       "Event with sub-events",
			timePeriod: "2024-01",
			text:       "Main Event",
			subEvents:  []string{"Sub Event 1", "Sub Event 2"},
			contains: []string{
				"2024-01",
				"Main Event",
				"Sub Event 1",
				"Sub Event 2",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			event := NewEvent(tt.timePeriod, tt.text)
			for _, subText := range tt.subEvents {
				event.AddSubEvent(subText)
			}
			result := event.String()

			for _, want := range tt.contains {
				if !strings.Contains(result, want) {
					t.Errorf("String() missing expected content %q in:\n%s", want, result)
				}
			}
		})
	}
}
