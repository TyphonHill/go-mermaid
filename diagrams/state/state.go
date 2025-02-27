package state

import (
	"fmt"
	"strings"

	"github.com/TyphonHill/go-mermaid/diagrams/utils/basediagram"
)

// StateType represents the different types of states in a state diagram.
type StateType string

// Predefined state types for state diagrams.
const (
	StateNormal    StateType = "normal"
	StateStart     StateType = "start"
	StateEnd       StateType = "end"
	StateChoice    StateType = "choice"
	StateFork      StateType = "fork"
	StateJoin      StateType = "join"
	StateComposite StateType = "composite"
)

// Base string formats for state diagram elements
const (
	baseStartState     string = basediagram.Indentation + "[*] --> %s\n"
	baseEndState       string = basediagram.Indentation + "%s --> [*]\n"
	baseChoiceState    string = basediagram.Indentation + "state %s <<choice>>\n"
	baseForkState      string = basediagram.Indentation + "state %s <<fork>>\n"
	baseJoinState      string = basediagram.Indentation + "state %s <<join>>\n"
	baseNormalState    string = basediagram.Indentation + "state %q as %s\n"
	baseCompositeStart string = basediagram.Indentation + "state %s {\n"
	baseCompositeEnd   string = basediagram.Indentation + "}\n"
	baseNote           string = basediagram.Indentation + "note %s of %s: %s\n"
)

// NotePosition represents the positioning of a note in a state diagram.
type NotePosition string

// Predefined note positions
const (
	NoteLeft  NotePosition = "left"
	NoteRight NotePosition = "right"
)

// Note represents an annotation attached to a state
type Note struct {
	Text     string
	Position NotePosition
}

// State represents a state in a state diagram.
type State struct {
	ID          string
	Description string
	Type        StateType
	Nested      []*State
	Note        *Note
}

// NewState creates a new State with the specified properties.
func NewState(id, description string, stateType StateType) *State {
	return &State{
		ID:          id,
		Description: description,
		Type:        stateType,
		Nested:      make([]*State, 0),
	}
}

// AddNestedState adds a nested state to the current state.
func (s *State) AddNestedState(id, description string, stateType StateType) *State {
	nested := NewState(id, description, stateType)
	s.Nested = append(s.Nested, nested)
	return nested
}

// AddNote adds a note to the state
func (s *State) AddNote(text string, position NotePosition) *State {
	s.Note = &Note{
		Text:     text,
		Position: position,
	}
	return s
}

// String generates a Mermaid-formatted string representation of the state with custom indentation.
func (s *State) String(curIndentation string) string {
	var sb strings.Builder

	switch s.Type {
	case StateStart:
		sb.WriteString(fmt.Sprintf("%s%s", curIndentation, fmt.Sprintf(baseStartState, s.ID)))
	case StateEnd:
		sb.WriteString(fmt.Sprintf("%s%s", curIndentation, fmt.Sprintf(baseEndState, s.ID)))
	case StateChoice:
		sb.WriteString(fmt.Sprintf("%s%s", curIndentation, fmt.Sprintf(baseChoiceState, s.ID)))
	case StateFork:
		sb.WriteString(fmt.Sprintf("%s%s", curIndentation, fmt.Sprintf(baseForkState, s.ID)))
	case StateJoin:
		sb.WriteString(fmt.Sprintf("%s%s", curIndentation, fmt.Sprintf(baseJoinState, s.ID)))
	default:
		if s.Description != "" {
			sb.WriteString(fmt.Sprintf("%s%s", curIndentation, fmt.Sprintf(baseNormalState, s.Description, s.ID)))
		}
	}

	if len(s.Nested) > 0 {
		sb.WriteString(fmt.Sprintf("%s%s", curIndentation, fmt.Sprintf(baseCompositeStart, s.ID)))
		nextIndentation := fmt.Sprintf("%s    ", curIndentation)
		for _, nested := range s.Nested {
			sb.WriteString(nested.String(nextIndentation))
		}
		sb.WriteString(fmt.Sprintf("%s%s", curIndentation, baseCompositeEnd))
	}

	if s.Note != nil {
		sb.WriteString(fmt.Sprintf("%s%s", curIndentation,
			fmt.Sprintf(baseNote, s.Note.Position, s.ID, s.Note.Text)))
	}

	return sb.String()
}
