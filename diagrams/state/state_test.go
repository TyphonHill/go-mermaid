package state

import (
	"reflect"
	"strings"
	"testing"
)

func TestNewState(t *testing.T) {
	tests := []struct {
		name        string
		id          string
		description string
		stateType   StateType
		want        *State
	}{
		{
			name:        "Create normal state",
			id:          "idle",
			description: "Idle State",
			stateType:   StateNormal,
			want: &State{
				ID:          "idle",
				Description: "Idle State",
				Type:        StateNormal,
				Nested:      make([]*State, 0),
			},
		},
		{
			name:        "Create choice state",
			id:          "decision",
			description: "Decision Point",
			stateType:   StateChoice,
			want: &State{
				ID:          "decision",
				Description: "Decision Point",
				Type:        StateChoice,
				Nested:      make([]*State, 0),
			},
		},
		{
			name:        "Create fork state",
			id:          "fork1",
			description: "",
			stateType:   StateFork,
			want: &State{
				ID:          "fork1",
				Description: "",
				Type:        StateFork,
				Nested:      make([]*State, 0),
			},
		},
		{
			name:        "Create join state",
			id:          "join1",
			description: "",
			stateType:   StateJoin,
			want: &State{
				ID:          "join1",
				Description: "",
				Type:        StateJoin,
				Nested:      make([]*State, 0),
			},
		},
		{
			name:        "Create start state",
			id:          "[*]",
			description: "",
			stateType:   StateStart,
			want: &State{
				ID:          "[*]",
				Description: "",
				Type:        StateStart,
				Nested:      make([]*State, 0),
			},
		},
		{
			name:        "Create end state",
			id:          "[*]",
			description: "",
			stateType:   StateEnd,
			want: &State{
				ID:          "[*]",
				Description: "",
				Type:        StateEnd,
				Nested:      make([]*State, 0),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewState(tt.id, tt.description, tt.stateType)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewState() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestState_String(t *testing.T) {
	tests := []struct {
		name        string
		state       *State
		setup       func(*State)
		indentation string
		contains    []string
	}{
		{
			name:        "Normal state",
			state:       NewState("S1", "State 1", StateNormal),
			indentation: "",
			contains: []string{
				`state "State 1" as S1`,
			},
		},
		{
			name:        "Choice state",
			state:       NewState("C1", "", StateChoice),
			indentation: "",
			contains: []string{
				"state C1 <<choice>>",
			},
		},
		{
			name:        "Fork state",
			state:       NewState("F1", "", StateFork),
			indentation: "",
			contains: []string{
				"state F1 <<fork>>",
			},
		},
		{
			name:        "Join state",
			state:       NewState("J1", "", StateJoin),
			indentation: "",
			contains: []string{
				"state J1 <<join>>",
			},
		},
		{
			name:        "Start state",
			state:       NewState("[*]", "", StateStart),
			indentation: "",
			contains: []string{
				"[*] --> ",
			},
		},
		{
			name:        "End state",
			state:       NewState("[*]", "", StateEnd),
			indentation: "",
			contains: []string{
				" --> [*]",
			},
		},
		{
			name:        "State with indentation",
			state:       NewState("S1", "State 1", StateNormal),
			indentation: "    ",
			contains: []string{
				`    state "State 1" as S1`,
			},
		},
		{
			name:  "State with note",
			state: NewState("S1", "State 1", StateNormal),
			setup: func(s *State) {
				s.AddNote("This is a note", NoteLeft)
			},
			contains: []string{
				`state "State 1" as S1`,
				"note left of S1: This is a note",
			},
		},
		{
			name:  "Composite state with nested states",
			state: NewState("CS1", "Composite", StateComposite),
			setup: func(s *State) {
				s.AddNestedState("N1", "Nested 1", StateNormal)
				s.AddNestedState("N2", "Nested 2", StateNormal)
			},
			contains: []string{
				"state CS1 {",
				`    state "Nested 1" as N1`,
				`    state "Nested 2" as N2`,
				"}",
			},
		},
		{
			name:  "Composite state with note and nested states",
			state: NewState("CS1", "Composite", StateComposite),
			setup: func(s *State) {
				s.AddNestedState("N1", "Nested 1", StateNormal)
				s.AddNote("Composite note", NoteRight)
			},
			contains: []string{
				"state CS1 {",
				`    state "Nested 1" as N1`,
				"}",
				"note right of CS1: Composite note",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup(tt.state)
			}

			got := tt.state.String(tt.indentation)
			for _, want := range tt.contains {
				if !strings.Contains(got, want) {
					t.Errorf("String() missing expected content %q in:\n%s", want, got)
				}
			}
		})
	}
}

func TestState_AddNestedState(t *testing.T) {
	tests := []struct {
		name        string
		parentState *State
		nestedID    string
		description string
		stateType   StateType
		want        *State
	}{
		{
			name: "Add normal nested state",
			parentState: &State{
				ID:          "parent",
				Description: "Parent State",
				Type:        StateComposite,
				Nested:      make([]*State, 0),
			},
			nestedID:    "child",
			description: "Child State",
			stateType:   StateNormal,
			want: &State{
				ID:          "child",
				Description: "Child State",
				Type:        StateNormal,
				Nested:      make([]*State, 0),
			},
		},
		{
			name: "Add choice nested state",
			parentState: &State{
				ID:          "parent",
				Description: "Parent State",
				Type:        StateComposite,
				Nested:      make([]*State, 0),
			},
			nestedID:    "choice",
			description: "",
			stateType:   StateChoice,
			want: &State{
				ID:          "choice",
				Description: "",
				Type:        StateChoice,
				Nested:      make([]*State, 0),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.parentState.AddNestedState(tt.nestedID, tt.description, tt.stateType)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddNestedState() = %v, want %v", got, tt.want)
			}
			if len(tt.parentState.Nested) != 1 {
				t.Errorf("Parent state nested count = %d, want 1", len(tt.parentState.Nested))
			}
			if !reflect.DeepEqual(tt.parentState.Nested[0], tt.want) {
				t.Errorf("Parent's nested state = %v, want %v", tt.parentState.Nested[0], tt.want)
			}
		})
	}
}

func TestState_AddNote(t *testing.T) {
	tests := []struct {
		name     string
		state    *State
		text     string
		position NotePosition
		want     *Note
	}{
		{
			name:     "Add left note",
			state:    NewState("S1", "State 1", StateNormal),
			text:     "Left note",
			position: NoteLeft,
			want: &Note{
				Text:     "Left note",
				Position: NoteLeft,
			},
		},
		{
			name:     "Add right note",
			state:    NewState("S1", "State 1", StateNormal),
			text:     "Right note",
			position: NoteRight,
			want: &Note{
				Text:     "Right note",
				Position: NoteRight,
			},
		},
		{
			name:     "Replace existing note",
			state:    NewState("S1", "State 1", StateNormal),
			text:     "New note",
			position: NoteRight,
			want: &Note{
				Text:     "New note",
				Position: NoteRight,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// For replace test, add initial note
			if tt.name == "Replace existing note" {
				tt.state.AddNote("Old note", NoteLeft)
			}

			result := tt.state.AddNote(tt.text, tt.position)

			// Test method chaining
			if result != tt.state {
				t.Error("AddNote() should return state for chaining")
			}

			// Test note was set correctly
			if tt.state.Note == nil {
				t.Fatal("Note was not set")
			}
			if tt.state.Note.Text != tt.want.Text {
				t.Errorf("Note.Text = %v, want %v", tt.state.Note.Text, tt.want.Text)
			}
			if tt.state.Note.Position != tt.want.Position {
				t.Errorf("Note.Position = %v, want %v", tt.state.Note.Position, tt.want.Position)
			}
		})
	}
}
