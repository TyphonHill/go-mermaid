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
		want       string
	}{
		{
			name:       "Event with time period",
			timePeriod: "2024-01",
			text:       "Test Event",
			want:       "2024-01 \t\t: Test Event",
		},
		{
			name:       "Regular sub-event without time period",
			timePeriod: "",
			text:       "Sub Event",
			want:       ": Sub Event",
		},
		{
			name:       "First sub-event without time period",
			timePeriod: "",
			text:       "First Sub Event",
			want:       "First Sub Event",
		},
		{
			name:       "Event with sub-events",
			timePeriod: "2024-01",
			text:       "Main Event",
			subEvents:  []string{"Sub Event 1", "Sub Event 2"},
			want:       "2024-01 \t\t: Main Event\n\t\t: Sub Event 1\n\t\t: Sub Event 2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			event := NewEvent(tt.timePeriod, tt.text)
			for _, subText := range tt.subEvents {
				event.AddSubEvent(subText)
			}
			result := event.String()

			if !strings.Contains(result, tt.want) {
				t.Errorf("String() = %v, want %v", result, tt.want)
			}
		})
	}
}
