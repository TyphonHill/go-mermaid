package timeline

import "fmt"

// Event represents a single event in the timeline
type Event struct {
	TimePeriod string
	Text       string
}

// NewEvent creates a new timeline event
func NewEvent(timePeriod string, text string) *Event {
	return &Event{
		TimePeriod: timePeriod,
		Text:       text,
	}
}

// String generates the Mermaid syntax for the event
func (e *Event) String() string {
	if e.TimePeriod != "" {
		return fmt.Sprintf("\t\t%s : %s\n", e.TimePeriod, e.Text)
	}
	return fmt.Sprintf("\t\t: %s\n", e.Text)
}
