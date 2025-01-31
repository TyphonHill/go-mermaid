package timeline

import "fmt"

// Base string formats for timeline events
const (
	baseEventWithTime    string = "\t\t%s : %s\n"
	baseEventWithoutTime string = "\t\t: %s\n"
)

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
		return fmt.Sprintf(baseEventWithTime, e.TimePeriod, e.Text)
	}
	return fmt.Sprintf(baseEventWithoutTime, e.Text)
}
