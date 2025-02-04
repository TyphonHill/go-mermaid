package timeline

import (
	"fmt"
	"strings"
)

// Base string formats for timeline sections
const (
	baseSectionTitle string = "\tsection %s\n"
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
func (s *Section) AddEvent(title string, text string) *Section {
	s.Events = append(s.Events, NewEvent(title, text))
	return s
}

// AddSubEvent adds a new sub-event (without time period) to the section
func (s *Section) AddSubEvent(text string) *Section {
	s.AddEvent("", text)
	return s
}

// String generates the Mermaid syntax for the section
func (s *Section) String() string {
	var sb strings.Builder

	if s.Title != "" {
		sb.WriteString(fmt.Sprintf(baseSectionTitle, s.Title))
	}

	for _, event := range s.Events {
		sb.WriteString(event.String())
	}

	return sb.String()
}
