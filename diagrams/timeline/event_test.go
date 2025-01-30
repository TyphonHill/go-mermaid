package timeline

import (
	"strings"
	"testing"
)

func TestNewEvent(t *testing.T) {
	timePeriod := "2024-01"
	text := "Test Event"
	event := NewEvent(timePeriod, text)

	if event.TimePeriod != timePeriod {
		t.Errorf("NewEvent().TimePeriod = %v, want %v", event.TimePeriod, timePeriod)
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
		want       string
	}{
		{
			name:       "Event with time period",
			timePeriod: "2024-01",
			text:       "Test Event",
			want:       "2024-01 : Test Event",
		},
		{
			name:       "Sub-event without time period",
			timePeriod: "",
			text:       "Sub Event",
			want:       ": Sub Event",
		},
		{
			name:       "Event with special characters",
			timePeriod: "Q1 2024",
			text:       "Event: with \"quotes\" and {braces}",
			want:       "Q1 2024 : Event: with \"quotes\" and {braces}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			event := NewEvent(tt.timePeriod, tt.text)
			result := event.String()

			if !strings.Contains(result, tt.want) {
				t.Errorf("String() = %v, want %v", result, tt.want)
			}
		})
	}
}
