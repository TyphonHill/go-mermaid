package state

import (
	"fmt"
)

// TransitionType represents the different types of transitions in a state diagram.
type TransitionType string

// Predefined transition types for state diagrams.
const (
	TransitionSolid  TransitionType = "solid"
	TransitionDashed TransitionType = "dashed"
)

// Base string formats for transition diagram elements
const (
	baseTransition         string = "%s --> %s"
	baseTransitionWithDesc string = ": %s"
	baseNewline            string = "\n"
)

// Transition represents a transition between states in a state diagram.
type Transition struct {
	From        *State
	To          *State
	Description string
	Type        TransitionType
}

// NewTransition creates a new Transition between two states.
func NewTransition(from, to *State, description string) *Transition {
	return &Transition{
		From:        from,
		To:          to,
		Description: description,
		Type:        TransitionSolid,
	}
}

// SetType sets the transition type and returns the transition for chaining
func (t *Transition) SetType(transitionType TransitionType) *Transition {
	t.Type = transitionType
	return t
}

// String generates a Mermaid-formatted string representation of the transition with custom indentation.
func (t *Transition) String(indentation string) string {
	var fromID, toID string

	// Handle nil states
	if t.From == nil {
		fromID = "[*]"
	} else {
		fromID = t.From.ID
	}

	if t.To == nil {
		toID = "[*]"
	} else {
		toID = t.To.ID
	}

	if t.Description == "" {
		return fmt.Sprintf("%s\t%s --> %s\n", indentation, fromID, toID)
	}
	return fmt.Sprintf("%s\t%s --> %s: %s\n", indentation, fromID, toID, t.Description)
}
