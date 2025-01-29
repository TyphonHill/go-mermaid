package state

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Diagram represents a state diagram with states, transitions, and rendering options.
type Diagram struct {
	Title         string
	States        []*State
	Transitions   []*Transition
	markdownFence bool
}

// NewDiagram creates a new state diagram with default settings.
func NewDiagram() *Diagram {
	return &Diagram{
		States:      make([]*State, 0),
		Transitions: make([]*Transition, 0),
	}
}

// EnableMarkdownFence enables markdown fencing for the diagram output.
func (d *Diagram) EnableMarkdownFence() {
	d.markdownFence = true
}

// DisableMarkdownFence disables markdown fencing for the diagram output.
func (d *Diagram) DisableMarkdownFence() {
	d.markdownFence = false
}

// AddState creates and adds a new state to the diagram.
func (d *Diagram) AddState(id, description string, stateType StateType) *State {
	state := NewState(id, description, stateType)
	d.States = append(d.States, state)
	return state
}

// AddTransition creates and adds a new transition between states.
func (d *Diagram) AddTransition(from, to *State, description string) *Transition {
	transition := NewTransition(from, to, description)
	d.Transitions = append(d.Transitions, transition)
	return transition
}

// String generates a Mermaid-formatted string representation of the state diagram.
func (d *Diagram) String() string {
	var sb strings.Builder

	if d.markdownFence {
		sb.WriteString("```mermaid\n")
	}

	if d.Title != "" {
		sb.WriteString(fmt.Sprintf("---\ntitle: %s\n---\n\n", d.Title))
	}

	sb.WriteString("stateDiagram-v2\n")

	// Render states
	for _, state := range d.States {
		sb.WriteString(state.String(""))
	}

	// Render transitions
	for _, transition := range d.Transitions {
		sb.WriteString(transition.String(""))
	}

	if d.markdownFence {
		sb.WriteString("```\n")
	}

	return sb.String()
}

// RenderToFile saves the diagram to a file, automatically enabling markdown fencing for .md files.
func (d *Diagram) RenderToFile(path string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	originalFenceState := d.markdownFence
	if strings.ToLower(filepath.Ext(path)) == ".md" {
		d.EnableMarkdownFence()
	}

	content := d.String()
	d.markdownFence = originalFenceState

	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
