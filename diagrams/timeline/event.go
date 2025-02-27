package timeline

import (
	"fmt"
	"strings"

	"github.com/TyphonHill/go-mermaid/diagrams/utils/basediagram"
)

// Base string formats for timeline events
const (
	eventTitle string = basediagram.Indentation + basediagram.Indentation + "%s\n"
	eventText  string = basediagram.Indentation + basediagram.Indentation + ": %s\n"
)

// Event represents a single event in the timeline
type Event struct {
	Title     string
	Text      string
	SubEvents []*Event
}

// NewEvent creates a new timeline event
func NewEvent(title string, text string) *Event {
	return &Event{
		Title: title,
		Text:  text,
	}
}

func (e *Event) AddSubEvent(text string) *Event {
	subEvent := NewEvent("", text)
	e.SubEvents = append(e.SubEvents, subEvent)
	return e
}

// String generates the Mermaid syntax for the event
func (e *Event) String() string {
	var sb strings.Builder

	if e.Title != "" {
		sb.WriteString(fmt.Sprintf(eventTitle, e.Title))
	}

	if e.Text != "" {
		sb.WriteString(fmt.Sprintf(eventText, e.Text))
	}

	for _, subEvent := range e.SubEvents {
		sb.WriteString(subEvent.String())
	}

	return sb.String()
}
