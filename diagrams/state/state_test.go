package state

import (
	"reflect"
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
			name:        "Create empty state",
			id:          "",
			description: "",
			stateType:   StateNormal,
			want: &State{
				ID:          "",
				Description: "",
				Type:        StateNormal,
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

func TestState_AddNestedState(t *testing.T) {
	tests := []struct {
		name        string
		parentState *State
		nestedID    string
		description string
		stateType   StateType
		wantNested  *State
	}{
		{
			name: "Add nested state to normal state",
			parentState: &State{
				ID:          "parent",
				Description: "Parent State",
				Type:        StateComposite,
				Nested:      make([]*State, 0),
			},
			nestedID:    "child",
			description: "Child State",
			stateType:   StateNormal,
			wantNested: &State{
				ID:          "child",
				Description: "Child State",
				Type:        StateNormal,
				Nested:      make([]*State, 0),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.parentState.AddNestedState(tt.nestedID, tt.description, tt.stateType)
			if !reflect.DeepEqual(got, tt.wantNested) {
				t.Errorf("AddNestedState() = %v, want %v", got, tt.wantNested)
			}
			if len(tt.parentState.Nested) != 1 {
				t.Errorf("Parent state nested count = %d, want 1", len(tt.parentState.Nested))
			}
			if !reflect.DeepEqual(tt.parentState.Nested[0], tt.wantNested) {
				t.Errorf("Parent's nested state = %v, want %v", tt.parentState.Nested[0], tt.wantNested)
			}
		})
	}
}
