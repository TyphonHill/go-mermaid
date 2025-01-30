package timeline

import (
	"fmt"
	"strings"
)

// Section represents a section in the timeline diagram
type Section struct {
	Title  string
	Events []*Event
}

// NewSection creates a new timeline section
func NewSection(title string) *Section {
	return &Section{
		Title:  title,
		Events: make([]*Event, 0),
	}
}

// AddEvent adds a new event to the section
func (s *Section) AddEvent(timePeriod string, text string) *Event {
	event := NewEvent(timePeriod, text)
	s.Events = append(s.Events, event)
	return event
}

// AddSubEvent adds a new sub-event (without time period) to the section
func (s *Section) AddSubEvent(text string) *Event {
	return s.AddEvent("", text)
}

// String generates the Mermaid syntax for the section
func (s *Section) String() string {
	var sb strings.Builder

	if s.Title != "" {
		sb.WriteString(fmt.Sprintf("\tsection %s\n", s.Title))
	}

	for _, event := range s.Events {
		sb.WriteString(event.String())
	}

	return sb.String()
}
