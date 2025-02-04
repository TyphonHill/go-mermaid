package timeline

import (
	"fmt"
	"strings"
)

// Base string formats for timeline events
const (
	baseEventWithTitle         string = "\t\t%s : %s\n"
	baseEventWithoutTitle      string = "\t\t: %s\n"
	baseFirstEventWithoutTitle string = "\t\t%s\n"
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
func (e *Event) String(isFirstEvent bool) string {
	var sb strings.Builder

	if e.Title != "" {
		sb.WriteString(fmt.Sprintf(baseEventWithTitle, e.Title, e.Text))
	} else {
		if isFirstEvent {
			sb.WriteString(fmt.Sprintf(baseFirstEventWithoutTitle, e.Text))
		} else {
			sb.WriteString(fmt.Sprintf(baseEventWithoutTitle, e.Text))
		}
	}

	for _, subEvent := range e.SubEvents {
		sb.WriteString(subEvent.String(false))
	}

	return sb.String()
}
