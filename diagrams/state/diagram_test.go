package state

import (
	"os"
	"path/filepath"
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
	if diagram.markdownFence {
		t.Error("NewDiagram() markdownFence = true, want false")
	}
}

func TestDiagram_EnableDisableMarkdownFence(t *testing.T) {
	diagram := NewDiagram()

	diagram.EnableMarkdownFence()
	if !diagram.markdownFence {
		t.Error("EnableMarkdownFence() did not enable markdown fence")
	}

	diagram.DisableMarkdownFence()
	if diagram.markdownFence {
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

func TestDiagram_AddTransition(t *testing.T) {
	diagram := NewDiagram()
	state1 := diagram.AddState("state1", "State 1", StateNormal)
	state2 := diagram.AddState("state2", "State 2", StateNormal)

	transition := diagram.AddTransition(state1, state2, "Test Transition")

	if len(diagram.Transitions) != 1 {
		t.Errorf("AddTransition() resulted in %d transitions, want 1", len(diagram.Transitions))
	}

	if transition.From != state1 || transition.To != state2 || transition.Description != "Test Transition" {
		t.Errorf("AddTransition() created incorrect transition: %v", transition)
	}
}

func TestDiagram_String(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(*Diagram)
		want    []string
		notWant []string
	}{
		{
			name: "Basic diagram with title",
			setup: func(d *Diagram) {
				d.Title = "Test Diagram"
				d.AddState("s1", "State 1", StateNormal)
			},
			want: []string{
				"title: Test Diagram",
				"stateDiagram-v2",
				`state "State 1" as s1`,
			},
		},
		{
			name: "Diagram with transition",
			setup: func(d *Diagram) {
				s1 := d.AddState("s1", "State 1", StateNormal)
				s2 := d.AddState("s2", "State 2", StateNormal)
				d.AddTransition(s1, s2, "goes to")
			},
			want: []string{
				"stateDiagram-v2",
				`state "State 1" as s1`,
				`state "State 2" as s2`,
				"s1 --> s2: goes to",
			},
		},
		{
			name: "Diagram with markdown fence",
			setup: func(d *Diagram) {
				d.EnableMarkdownFence()
				d.AddState("s1", "State 1", StateNormal)
			},
			want: []string{
				"```mermaid",
				"stateDiagram-v2",
				"```",
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
					t.Errorf("String() result missing expected content %q", want)
				}
			}

			for _, notWant := range tt.notWant {
				if strings.Contains(result, notWant) {
					t.Errorf("String() result contains unexpected content %q", notWant)
				}
			}
		})
	}
}

func TestDiagram_RenderToFile(t *testing.T) {
	// Create temp directory for test files
	tempDir, err := os.MkdirTemp("", "state_test_*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a sample diagram
	diagram := NewDiagram()
	diagram.Title = "Test Diagram"
	state1 := diagram.AddState("s1", "State 1", StateNormal)
	state2 := diagram.AddState("s2", "State 2", StateNormal)
	diagram.AddTransition(state1, state2, "Test Transition")

	tests := []struct {
		name           string
		filename       string
		setupFence     bool
		expectFence    bool
		expectError    bool
		validateOutput func(string) bool
	}{
		{
			name:        "Save as markdown file",
			filename:    "diagram.md",
			setupFence:  false, // Even with fencing disabled, .md should enable it
			expectFence: true,
			validateOutput: func(content string) bool {
				return strings.HasPrefix(content, "```mermaid\n") &&
					strings.HasSuffix(content, "```\n")
			},
		},
		{
			name:        "Save as text file with fencing enabled",
			filename:    "diagram.txt",
			setupFence:  true,
			expectFence: true,
			validateOutput: func(content string) bool {
				return strings.HasPrefix(content, "```mermaid\n") &&
					strings.HasSuffix(content, "```\n")
			},
		},
		{
			name:        "Save as text file without fencing",
			filename:    "diagram.txt",
			setupFence:  false,
			expectFence: false,
			validateOutput: func(content string) bool {
				return !strings.Contains(content, "```mermaid")
			},
		},
		{
			name:        "Save to nested directory",
			filename:    "nested/dir/diagram.txt",
			setupFence:  false,
			expectFence: false,
			validateOutput: func(content string) bool {
				return strings.Contains(content, "Test Diagram")
			},
		},
		{
			name:        "Save with invalid path",
			filename:    string([]byte{0}), // Invalid filename
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up diagram fencing
			if tt.setupFence {
				diagram.EnableMarkdownFence()
			} else {
				diagram.DisableMarkdownFence()
			}

			// Create full path
			path := filepath.Join(tempDir, tt.filename)

			// Attempt to render
			err := diagram.RenderToFile(path)

			// Check error expectation
			if tt.expectError {
				if err == nil {
					t.Error("Expected error but got none")
				}
				return
			}

			// If we don't expect an error but got one
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			// Read the file content
			content, err := os.ReadFile(path)
			if err != nil {
				t.Fatalf("Failed to read output file: %v", err)
			}

			// Validate content
			if tt.validateOutput != nil {
				if !tt.validateOutput(string(content)) {
					t.Error("Output validation failed")
				}
			}

			// Verify fence state wasn't changed permanently
			if diagram.markdownFence != tt.setupFence {
				t.Error("Diagram fence state was permanently modified")
			}
		})
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
