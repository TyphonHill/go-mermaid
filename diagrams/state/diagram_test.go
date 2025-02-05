package state

import (
	"strings"
	"testing"
)

func TestNewDiagram(t *testing.T) {
	diagram := NewDiagram()

	if diagram.States == nil {
		t.Error("NewDiagram() States is nil, want empty slice")
	}
	if diagram.Transitions == nil {
		t.Error("NewDiagram() Transitions is nil, want empty slice")
	}
	if diagram.IsMarkdownFenceEnabled() {
		t.Error("NewDiagram() markdown fence enabled by default, want disabled")
	}
}

func TestDiagram_EnableDisableMarkdownFence(t *testing.T) {
	diagram := NewDiagram()

	diagram.EnableMarkdownFence()
	if !diagram.IsMarkdownFenceEnabled() {
		t.Error("EnableMarkdownFence() did not enable markdown fence")
	}

	diagram.DisableMarkdownFence()
	if diagram.IsMarkdownFenceEnabled() {
		t.Error("DisableMarkdownFence() did not disable markdown fence")
	}
}

func TestDiagram_AddState(t *testing.T) {
	diagram := NewDiagram()
	state := diagram.AddState("test", "Test State", StateNormal)

	if len(diagram.States) != 1 {
		t.Errorf("AddState() resulted in %d states, want 1", len(diagram.States))
	}

	if state.ID != "test" || state.Description != "Test State" || state.Type != StateNormal {
		t.Errorf("AddState() = %v, want {ID: test, Description: Test State, Type: normal}", state)
	}
}

func TestAddTransition(t *testing.T) {
	diagram := NewDiagram()
	state1 := diagram.AddState("state1", "State 1", StateNormal)
	state2 := diagram.AddState("state2", "State 2", StateNormal)

	transition := diagram.AddTransition(state1, state2, "test")
	if len(diagram.Transitions) != 1 {
		t.Errorf("Expected 1 transition, got %d", len(diagram.Transitions))
	}

	if transition.From != state1 || transition.To != state2 {
		t.Error("Transition has incorrect states")
	}

	transition.SetType(TransitionDashed)
	if transition.Type != TransitionDashed {
		t.Error("SetType did not update transition type")
	}
}

func TestDiagramString(t *testing.T) {
	diagram := NewDiagram()
	diagram.Title = "Test Diagram"

	state1 := diagram.AddState("state1", "State 1", StateNormal)
	state2 := diagram.AddState("state2", "State 2", StateNormal)

	diagram.AddTransition(state1, state2, "forward")
	diagram.AddTransition(state2, state1, "back")

	result := diagram.String()
	expected := `---
title: Test Diagram
---

stateDiagram-v2
	state "State 1" as state1
	state "State 2" as state2
	state1 --> state2: forward
	state2 --> state1: back
`

	if result != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, result)
	}
}

func TestDiagram_RenderState(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(*Diagram)
		want    []string
		notWant []string
	}{
		{
			name: "Render start state",
			setup: func(d *Diagram) {
				d.AddState("start", "Start State", StateStart)
			},
			want: []string{
				"\t[*] --> start",
			},
		},
		{
			name: "Render end state",
			setup: func(d *Diagram) {
				d.AddState("end", "End State", StateEnd)
			},
			want: []string{
				"\tend --> [*]",
			},
		},
		{
			name: "Render choice state",
			setup: func(d *Diagram) {
				d.AddState("choice", "Choice State", StateChoice)
			},
			want: []string{
				"\tstate choice <<choice>>",
			},
		},
		{
			name: "Render fork state",
			setup: func(d *Diagram) {
				d.AddState("fork", "Fork State", StateFork)
			},
			want: []string{
				"\tstate fork <<fork>>",
			},
		},
		{
			name: "Render join state",
			setup: func(d *Diagram) {
				d.AddState("join", "Join State", StateJoin)
			},
			want: []string{
				"\tstate join <<join>>",
			},
		},
		{
			name: "Render normal state with description",
			setup: func(d *Diagram) {
				d.AddState("normal", "Normal State", StateNormal)
			},
			want: []string{
				"\tstate \"Normal State\" as normal",
			},
		},
		{
			name: "Render composite state with nested states",
			setup: func(d *Diagram) {
				parent := d.AddState("parent", "Parent State", StateComposite)
				nested := NewState("child", "Child State", StateNormal)
				parent.Nested = append(parent.Nested, nested)
			},
			want: []string{
				"state parent {",
				"state \"Child State\" as child",
				"}",
			},
		},
		{
			name: "Render multiple nested levels",
			setup: func(d *Diagram) {
				parent := d.AddState("parent", "Parent State", StateComposite)
				child := NewState("child", "Child State", StateComposite)
				grandchild := NewState("grandchild", "Grandchild State", StateNormal)
				child.Nested = append(child.Nested, grandchild)
				parent.Nested = append(parent.Nested, child)
			},
			want: []string{
				"state parent {",
				"state child {",
				"state \"Grandchild State\" as grandchild",
				"}",
			},
		},
		{
			name: "Render state with special characters in description",
			setup: func(d *Diagram) {
				d.AddState("special", "State: with \"quotes\" and {braces}", StateNormal)
			},
			want: []string{
				`state "State: with \"quotes\" and {braces}" as special`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diagram := NewDiagram()
			tt.setup(diagram)
			result := diagram.String()

			for _, want := range tt.want {
				if !strings.Contains(result, want) {
					t.Errorf("String() result missing expected content %q in output:\n%s", want, result)
				}
			}

			for _, notWant := range tt.notWant {
				if strings.Contains(result, notWant) {
					t.Errorf("String() result contains unexpected content %q in output:\n%s", notWant, result)
				}
			}
		})
	}
}

func TestDiagram_RenderTransition(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(*Diagram)
		want    []string
		notWant []string
	}{
		{
			name: "Render transition with description",
			setup: func(d *Diagram) {
				s1 := d.AddState("s1", "State 1", StateNormal)
				s2 := d.AddState("s2", "State 2", StateNormal)
				d.AddTransition(s1, s2, "test transition")
			},
			want: []string{
				"s1 --> s2: test transition",
			},
		},
		{
			name: "Render transition without description",
			setup: func(d *Diagram) {
				s1 := d.AddState("s1", "State 1", StateNormal)
				s2 := d.AddState("s2", "State 2", StateNormal)
				d.AddTransition(s1, s2, "")
			},
			want: []string{
				"s1 --> s2",
			},
			notWant: []string{
				"s1 --> s2:",
			},
		},
		{
			name: "Render dashed transition",
			setup: func(d *Diagram) {
				s1 := d.AddState("s1", "State 1", StateNormal)
				s2 := d.AddState("s2", "State 2", StateNormal)
				t := d.AddTransition(s1, s2, "dashed")
				t.SetType(TransitionDashed)
			},
			want: []string{
				"s1 --> s2: dashed",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diagram := NewDiagram()
			tt.setup(diagram)
			result := diagram.String()

			for _, want := range tt.want {
				if !strings.Contains(result, want) {
					t.Errorf("String() result missing expected content %q in output:\n%s", want, result)
				}
			}

			for _, notWant := range tt.notWant {
				if strings.Contains(result, notWant) {
					t.Errorf("String() result contains unexpected content %q in output:\n%s", notWant, result)
				}
			}
		})
	}
}
