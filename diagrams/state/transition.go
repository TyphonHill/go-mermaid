package state

import (
	"fmt"
	"strings"
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

// SetType sets the transition type (solid or dashed).
func (t *Transition) SetType(transitionType TransitionType) {
	t.Type = transitionType
}

// String generates a Mermaid-formatted string representation of the transition with custom indentation.
func (t *Transition) String(curIndentation string) string {
	var sb strings.Builder

	// Handle transitions involving start/end states ([*])
	if t.From == nil {
		// This is a start state transition: [*] --> someState
		transitionStr := fmt.Sprintf(baseTransition, "[*]", t.To.ID)
		if t.Description != "" {
			transitionStr += fmt.Sprintf(baseTransitionWithDesc, t.Description)
		}
		sb.WriteString(fmt.Sprintf("%s\t%s%s", curIndentation, transitionStr, baseNewline))
		return sb.String()
	}

	if t.To == nil {
		// This is an end state transition: someState --> [*]
		transitionStr := fmt.Sprintf(baseTransition, t.From.ID, "[*]")
		if t.Description != "" {
			transitionStr += fmt.Sprintf(baseTransitionWithDesc, t.Description)
		}
		sb.WriteString(fmt.Sprintf("%s\t%s%s", curIndentation, transitionStr, baseNewline))
		return sb.String()
	}

	// Handle normal transitions
	transitionStr := fmt.Sprintf(baseTransition, t.From.ID, t.To.ID)
	if t.Description != "" {
		transitionStr += fmt.Sprintf(baseTransitionWithDesc, t.Description)
	}
	sb.WriteString(fmt.Sprintf("%s\t%s%s", curIndentation, transitionStr, baseNewline))

	return sb.String()
}
