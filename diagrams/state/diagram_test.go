package state

import (
	"reflect"
	"strings"
	"testing"

	"github.com/TyphonHill/go-mermaid/diagrams/utils/basediagram"
)

func TestNewDiagram(t *testing.T) {
	tests := []struct {
		name string
		want *Diagram
	}{
		{
			name: "Create new diagram with default settings",
			want: &Diagram{
				BaseDiagram: basediagram.NewBaseDiagram(NewStateConfigurationProperties()),
				States:      make([]*State, 0),
				Transitions: make([]*Transition, 0),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDiagram()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDiagram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiagram_String(t *testing.T) {
	tests := []struct {
		name     string
		diagram  *Diagram
		setup    func(*Diagram)
		contains []string
	}{
		{
			name:    "Empty diagram",
			diagram: NewDiagram(),
			contains: []string{
				"stateDiagram-v2",
			},
		},
		{
			name:    "Diagram with single state",
			diagram: NewDiagram(),
			setup: func(d *Diagram) {
				d.AddState("S1", "State 1", StateNormal)
			},
			contains: []string{
				"stateDiagram-v2",
				`state "State 1" as S1`,
			},
		},
		{
			name:    "Diagram with multiple states and transition",
			diagram: NewDiagram(),
			setup: func(d *Diagram) {
				s1 := d.AddState("S1", "State 1", StateNormal)
				s2 := d.AddState("S2", "State 2", StateNormal)
				d.AddTransition(s1, s2, "Next")
			},
			contains: []string{
				"stateDiagram-v2",
				`state "State 1" as S1`,
				`state "State 2" as S2`,
				"S1 --> S2: Next",
			},
		},
		{
			name:    "Diagram with choice state",
			diagram: NewDiagram(),
			setup: func(d *Diagram) {
				s1 := d.AddState("S1", "State 1", StateNormal)
				c1 := d.AddState("C1", "", StateChoice)
				s2 := d.AddState("S2", "State 2", StateNormal)
				d.AddTransition(s1, c1, "Check")
				d.AddTransition(c1, s2, "Yes")
			},
			contains: []string{
				"stateDiagram-v2",
				`state "State 1" as S1`,
				"state C1 <<choice>>",
				`state "State 2" as S2`,
				"S1 --> C1: Check",
				"C1 --> S2: Yes",
			},
		},
		{
			name:    "Diagram with composite state",
			diagram: NewDiagram(),
			setup: func(d *Diagram) {
				composite := d.AddState("CS", "Composite", StateComposite)
				nested := composite.AddNestedState("NS", "Nested", StateNormal)
				s2 := d.AddState("S2", "State 2", StateNormal)
				d.AddTransition(nested, s2, "Exit")
			},
			contains: []string{
				"stateDiagram-v2",
				"state CS {",
				`    state "Nested" as NS`,
				"}",
				`state "State 2" as S2`,
				"NS --> S2: Exit",
			},
		},
		{
			name:    "Diagram with start and end states",
			diagram: NewDiagram(),
			setup: func(d *Diagram) {
				start := d.AddState("[*]", "", StateStart)
				s1 := d.AddState("S1", "Process", StateNormal)
				end := d.AddState("[*]", "", StateEnd)
				d.AddTransition(start, s1, "Begin")
				d.AddTransition(s1, end, "Finish")
			},
			contains: []string{
				"stateDiagram-v2",
				"[*] --> S1: Begin",
				`state "Process" as S1`,
				"S1 --> [*]: Finish",
			},
		},
		{
			name:    "Diagram with fork and join",
			diagram: NewDiagram(),
			setup: func(d *Diagram) {
				s1 := d.AddState("S1", "Start", StateNormal)
				fork := d.AddState("F1", "", StateFork)
				p1 := d.AddState("P1", "Path 1", StateNormal)
				p2 := d.AddState("P2", "Path 2", StateNormal)
				join := d.AddState("J1", "", StateJoin)
				s2 := d.AddState("S2", "End", StateNormal)

				d.AddTransition(s1, fork, "Split")
				d.AddTransition(fork, p1, "")
				d.AddTransition(fork, p2, "")
				d.AddTransition(p1, join, "")
				d.AddTransition(p2, join, "")
				d.AddTransition(join, s2, "Merge")
			},
			contains: []string{
				"stateDiagram-v2",
				`state "Start" as S1`,
				"state F1 <<fork>>",
				`state "Path 1" as P1`,
				`state "Path 2" as P2`,
				"state J1 <<join>>",
				`state "End" as S2`,
				"S1 --> F1: Split",
				"F1 --> P1",
				"F1 --> P2",
				"P1 --> J1",
				"P2 --> J1",
				"J1 --> S2: Merge",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup(tt.diagram)
			}

			got := tt.diagram.String()
			for _, want := range tt.contains {
				if !strings.Contains(got, want) {
					t.Errorf("String() missing expected content %q in:\n%s", want, got)
				}
			}
		})
	}
}

func TestDiagram_AddState(t *testing.T) {
	tests := []struct {
		name        string
		id          string
		description string
		stateType   StateType
		want        *State
	}{
		{
			name:        "Add normal state",
			id:          "S1",
			description: "State 1",
			stateType:   StateNormal,
			want: &State{
				ID:          "S1",
				Description: "State 1",
				Type:        StateNormal,
				Nested:      make([]*State, 0),
			},
		},
		{
			name:        "Add choice state",
			id:          "C1",
			description: "",
			stateType:   StateChoice,
			want: &State{
				ID:          "C1",
				Description: "",
				Type:        StateChoice,
				Nested:      make([]*State, 0),
			},
		},
		{
			name:        "Add composite state",
			id:          "CS1",
			description: "Composite",
			stateType:   StateComposite,
			want: &State{
				ID:          "CS1",
				Description: "Composite",
				Type:        StateComposite,
				Nested:      make([]*State, 0),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diagram := NewDiagram()
			got := diagram.AddState(tt.id, tt.description, tt.stateType)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddState() = %v, want %v", got, tt.want)
			}

			if len(diagram.States) != 1 || !reflect.DeepEqual(diagram.States[0], got) {
				t.Error("State not added to diagram correctly")
			}
		})
	}
}

func TestDiagram_AddTransition(t *testing.T) {
	tests := []struct {
		name        string
		fromState   *State
		toState     *State
		description string
		want        *Transition
	}{
		{
			name: "Add simple transition",
			fromState: &State{
				ID:   "S1",
				Type: StateNormal,
			},
			toState: &State{
				ID:   "S2",
				Type: StateNormal,
			},
			description: "Next",
			want: &Transition{
				From:        &State{ID: "S1", Type: StateNormal},
				To:          &State{ID: "S2", Type: StateNormal},
				Description: "Next",
			},
		},
		{
			name: "Add transition without description",
			fromState: &State{
				ID:   "S1",
				Type: StateNormal,
			},
			toState: &State{
				ID:   "S2",
				Type: StateNormal,
			},
			description: "",
			want: &Transition{
				From:        &State{ID: "S1", Type: StateNormal},
				To:          &State{ID: "S2", Type: StateNormal},
				Description: "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diagram := NewDiagram()
			got := diagram.AddTransition(tt.fromState, tt.toState, tt.description)

			if got.From.ID != tt.want.From.ID || got.To.ID != tt.want.To.ID || got.Description != tt.want.Description {
				t.Errorf("AddTransition() = %v, want %v", got, tt.want)
			}

			if len(diagram.Transitions) != 1 || !reflect.DeepEqual(diagram.Transitions[0], got) {
				t.Error("Transition not added to diagram correctly")
			}
		})
	}
}
