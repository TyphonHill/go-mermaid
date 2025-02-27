package timeline

import (
	"fmt"
	"strings"

	"github.com/TyphonHill/go-mermaid/diagrams/utils/basediagram"
)

// Base string formats for timeline sections
const (
	baseSectionTitle string = basediagram.Indentation + "section %s\n"
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
func (s *Section) AddEvent(title string, text string) *Event {
	event := NewEvent(title, text)
	s.Events = append(s.Events, event)
	return event
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
