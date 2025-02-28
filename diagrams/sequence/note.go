package sequence

import (
	"fmt"
	"strings"
)

// NotePosition represents the positioning of a note in a sequence diagram.
type NotePosition string

// Predefined note positioning options.
const (
	NoteLeft  NotePosition = "left of"
	NoteRight NotePosition = "right of"
	NoteOver  NotePosition = "over"
)

// Base string formats for note elements
const (
	baseNoteLeft      string = "Note left of %s: %s\n"
	baseNoteRight     string = "Note right of %s: %s\n"
	baseNoteOver      string = "Note over %s: %s\n"
	baseNoteOverMulti string = "Note over %s,%s: %s\n"
)

// Note represents an annotation or comment in a sequence diagram.
type Note struct {
	Position NotePosition
	Text     string
	Actors   []*Actor
}

// newNote creates a new Note with the specified properties.
func newNote(position NotePosition, text string, actors ...*Actor) *Note {
	return &Note{
		Position: position,
		Text:     text,
		Actors:   actors,
	}
}

// String generates a Mermaid-formatted string representation of the note with custom indentation.
func (n *Note) String(curIndentation string) string {
	if len(n.Actors) == 0 {
		return ""
	}

	var sb strings.Builder

	switch n.Position {
	case NoteLeft:
		if len(n.Actors) == 1 {
			sb.WriteString(fmt.Sprintf("%s\t%s", curIndentation,
				fmt.Sprintf(baseNoteLeft, n.Actors[0].ID, n.Text)))
		}
	case NoteRight:
		if len(n.Actors) == 1 {
			sb.WriteString(fmt.Sprintf("%s\t%s", curIndentation,
				fmt.Sprintf(baseNoteRight, n.Actors[0].ID, n.Text)))
		}
	case NoteOver:
		if len(n.Actors) == 1 {
			sb.WriteString(fmt.Sprintf("%s\t%s", curIndentation,
				fmt.Sprintf(baseNoteOver, n.Actors[0].ID, n.Text)))
		} else if len(n.Actors) == 2 {
			sb.WriteString(fmt.Sprintf("%s\t%s", curIndentation,
				fmt.Sprintf(baseNoteOverMulti, n.Actors[0].ID, n.Actors[1].ID, n.Text)))
		}
	}

	return sb.String()
}
