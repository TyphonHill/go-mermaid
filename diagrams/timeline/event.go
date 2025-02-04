package timeline

import "fmt"

// Base string formats for timeline events
const (
	baseEventWithTitle    string = "\t\t%s : %s\n"
	baseEventWithoutTitle string = "\t\t: %s\n"
)

// Event represents a single event in the timeline
type Event struct {
	Title string
	Text  string
}

// NewEvent creates a new timeline event
func NewEvent(title string, text string) *Event {
	return &Event{
		Title: title,
		Text:  text,
	}
}

// String generates the Mermaid syntax for the event
func (e *Event) String() string {
	if e.Title != "" {
		return fmt.Sprintf(baseEventWithTitle, e.Title, e.Text)
	}
	return fmt.Sprintf(baseEventWithoutTitle, e.Text)
}
