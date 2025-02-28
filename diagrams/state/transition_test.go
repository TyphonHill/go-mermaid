package state

import (
	"reflect"
	"strings"
	"testing"
)

func TestNewTransition(t *testing.T) {
	tests := []struct {
		name        string
		from        *State
		to          *State
		description string
		want        *Transition
	}{
		{
			name:        "Create transition between normal states",
			from:        &State{ID: "S1", Type: StateNormal},
			to:          &State{ID: "S2", Type: StateNormal},
			description: "Next",
			want: &Transition{
				From:        &State{ID: "S1", Type: StateNormal},
				To:          &State{ID: "S2", Type: StateNormal},
				Description: "Next",
				Type:        TransitionSolid,
			},
		},
		{
			name:        "Create transition without description",
			from:        &State{ID: "S1", Type: StateNormal},
			to:          &State{ID: "S2", Type: StateNormal},
			description: "",
			want: &Transition{
				From:        &State{ID: "S1", Type: StateNormal},
				To:          &State{ID: "S2", Type: StateNormal},
				Description: "",
				Type:        TransitionSolid,
			},
		},
		{
			name:        "Create transition from start state",
			from:        nil,
			to:          &State{ID: "S1", Type: StateNormal},
			description: "Begin",
			want: &Transition{
				From:        nil,
				To:          &State{ID: "S1", Type: StateNormal},
				Description: "Begin",
				Type:        TransitionSolid,
			},
		},
		{
			name:        "Create transition to end state",
			from:        &State{ID: "S1", Type: StateNormal},
			to:          nil,
			description: "End",
			want: &Transition{
				From:        &State{ID: "S1", Type: StateNormal},
				To:          nil,
				Description: "End",
				Type:        TransitionSolid,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewTransition(tt.from, tt.to, tt.description)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTransition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransition_String(t *testing.T) {
	tests := []struct {
		name        string
		transition  *Transition
		setup       func(*Transition)
		indentation string
		contains    []string
	}{
		{
			name: "Normal transition with description",
			transition: NewTransition(
				&State{ID: "S1", Type: StateNormal},
				&State{ID: "S2", Type: StateNormal},
				"Next state",
			),
			indentation: "",
			contains: []string{
				"S1 --> S2: Next state",
			},
		},
		{
			name: "Transition without description",
			transition: NewTransition(
				&State{ID: "S1", Type: StateNormal},
				&State{ID: "S2", Type: StateNormal},
				"",
			),
			indentation: "",
			contains: []string{
				"S1 --> S2",
			},
		},
		{
			name: "Start transition",
			transition: NewTransition(
				nil,
				&State{ID: "S1", Type: StateNormal},
				"Begin",
			),
			indentation: "",
			contains: []string{
				"[*] --> S1: Begin",
			},
		},
		{
			name: "End transition",
			transition: NewTransition(
				&State{ID: "S1", Type: StateNormal},
				nil,
				"End",
			),
			indentation: "",
			contains: []string{
				"S1 --> [*]: End",
			},
		},
		{
			name: "Dashed transition",
			transition: NewTransition(
				&State{ID: "S1", Type: StateNormal},
				&State{ID: "S2", Type: StateNormal},
				"Dashed",
			),
			setup: func(t *Transition) {
				t.SetType(TransitionDashed)
			},
			contains: []string{
				"S1 --> S2: Dashed",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup(tt.transition)
			}

			got := tt.transition.String(tt.indentation)
			for _, want := range tt.contains {
				if !strings.Contains(got, want) {
					t.Errorf("String() missing expected content %q in:\n%s", want, got)
				}
			}
		})
	}
}

func TestTransition_SetType(t *testing.T) {
	tests := []struct {
		name           string
		transitionType TransitionType
		want           TransitionType
	}{
		{
			name:           "Set solid type",
			transitionType: TransitionSolid,
			want:           TransitionSolid,
		},
		{
			name:           "Set dashed type",
			transitionType: TransitionDashed,
			want:           TransitionDashed,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transition := NewTransition(
				&State{ID: "S1"},
				&State{ID: "S2"},
				"",
			)
			result := transition.SetType(tt.transitionType)

			if transition.Type != tt.want {
				t.Errorf("SetType() = %v, want %v", transition.Type, tt.want)
			}

			if result != transition {
				t.Error("SetType() should return transition for chaining")
			}
		})
	}
}
