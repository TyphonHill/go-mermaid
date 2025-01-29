package state

import (
	"reflect"
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
			name: "Create basic transition",
			from: &State{
				ID:          "start",
				Description: "Start State",
				Type:        StateNormal,
			},
			to: &State{
				ID:          "end",
				Description: "End State",
				Type:        StateNormal,
			},
			description: "Transition Description",
			want: &Transition{
				From: &State{
					ID:          "start",
					Description: "Start State",
					Type:        StateNormal,
				},
				To: &State{
					ID:          "end",
					Description: "End State",
					Type:        StateNormal,
				},
				Description: "Transition Description",
				Type:        TransitionSolid,
			},
		},
		{
			name: "Create transition with empty description",
			from: &State{
				ID:   "state1",
				Type: StateNormal,
			},
			to: &State{
				ID:   "state2",
				Type: StateNormal,
			},
			description: "",
			want: &Transition{
				From: &State{
					ID:   "state1",
					Type: StateNormal,
				},
				To: &State{
					ID:   "state2",
					Type: StateNormal,
				},
				Description: "",
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

func TestTransition_SetType(t *testing.T) {
	tests := []struct {
		name           string
		transition     *Transition
		transitionType TransitionType
		want           TransitionType
	}{
		{
			name: "Change to dashed transition",
			transition: &Transition{
				From:        &State{ID: "start"},
				To:          &State{ID: "end"},
				Type:        TransitionSolid,
				Description: "Test transition",
			},
			transitionType: TransitionDashed,
			want:           TransitionDashed,
		},
		{
			name: "Change to solid transition",
			transition: &Transition{
				From:        &State{ID: "start"},
				To:          &State{ID: "end"},
				Type:        TransitionDashed,
				Description: "Test transition",
			},
			transitionType: TransitionSolid,
			want:           TransitionSolid,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.transition.SetType(tt.transitionType)
			if tt.transition.Type != tt.want {
				t.Errorf("Transition.Type = %v, want %v", tt.transition.Type, tt.want)
			}
		})
	}
}
