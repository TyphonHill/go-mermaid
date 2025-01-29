package flowchart

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

// MockIDGenerator is a simple ID generator for testing
type MockIDGenerator struct {
	currentID uint64
}

func (m *MockIDGenerator) NextID() uint64 {
	current := m.currentID
	m.currentID++
	return current
}

func TestNewFlowchart(t *testing.T) {
	tests := []struct {
		name string
		want *Flowchart
	}{
		{
			name: "Create new flowchart with default settings",
			want: &Flowchart{
				Direction:     FlowchartDirectionTopToBottom,
				CurveStyle:    CurveStyleNone,
				classes:       make([]*Class, 0),
				nodes:         make([]*Node, 0),
				subgraphs:     make([]*Subgraph, 0),
				links:         make([]*Link, 0),
				markdownFence: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewFlowchart()

			// Remove the comparison of idGenerator as it's an interface
			tt.want.idGenerator = got.idGenerator

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFlowchart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlowchart_EnableMarkdownFence(t *testing.T) {
	tests := []struct {
		name      string
		flowchart *Flowchart
	}{
		{
			name:      "Enable markdown fence",
			flowchart: NewFlowchart(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.flowchart.EnableMarkdownFence()
			if !tt.flowchart.markdownFence {
				t.Error("EnableMarkdownFence() did not set markdownFence to true")
			}
		})
	}
}

func TestFlowchart_DisableMarkdownFence(t *testing.T) {
	tests := []struct {
		name      string
		flowchart *Flowchart
	}{
		{
			name:      "Disable markdown fence",
			flowchart: NewFlowchart(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// First enable it
			tt.flowchart.EnableMarkdownFence()
			// Then disable it
			tt.flowchart.DisableMarkdownFence()
			if tt.flowchart.markdownFence {
				t.Error("DisableMarkdownFence() did not set markdownFence to false")
			}
		})
	}
}

func TestFlowchart_String(t *testing.T) {
	tests := []struct {
		name      string
		flowchart *Flowchart
		setup     func(*Flowchart)
		wantStr   string
		contains  []string
	}{
		{
			name:      "Empty flowchart without fence",
			flowchart: NewFlowchart(),
			wantStr:   "flowchart TB\n",
		},
		{
			name:      "Empty flowchart with fence",
			flowchart: NewFlowchart(),
			setup: func(f *Flowchart) {
				f.EnableMarkdownFence()
			},
			wantStr: "```mermaid\nflowchart TB\n```\n",
		},
		{
			name:      "Flowchart with title and fence",
			flowchart: NewFlowchart(),
			setup: func(f *Flowchart) {
				f.EnableMarkdownFence()
				f.Title = "Test Flowchart"
			},
			contains: []string{
				"```mermaid\n",
				"---\ntitle: Test Flowchart\n---\n",
				"flowchart TB\n",
				"```\n",
			},
		},
		{
			name:      "Flowchart with nodes and links",
			flowchart: NewFlowchart(),
			setup: func(f *Flowchart) {
				f.EnableMarkdownFence()
				f.Title = "Test Flow"
				node1 := f.AddNode("Start")
				node2 := f.AddNode("End")
				f.AddLink(node1, node2)
			},
			contains: []string{
				"```mermaid\n",
				"flowchart TB\n",
				`0("Start")`,
				`1("End")`,
				"0 -->",
				"```\n",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup(tt.flowchart)
			}

			got := tt.flowchart.String()

			if tt.wantStr != "" && got != tt.wantStr {
				t.Errorf("Flowchart.String() = %v, want %v", got, tt.wantStr)
			}

			for _, want := range tt.contains {
				if !strings.Contains(got, want) {
					t.Errorf("Flowchart.String() output missing expected content: %v", want)
				}
			}
		})
	}
}

func TestFlowchart_RenderToFile(t *testing.T) {
	// Create temp directory for test files
	tempDir, err := os.MkdirTemp("", "flowchart_test_*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a sample flowchart
	flowchart := NewFlowchart()
	flowchart.Title = "Test Flowchart"
	node1 := flowchart.AddNode("Start")
	node2 := flowchart.AddNode("End")
	flowchart.AddLink(node1, node2)

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
			filename:    "flowchart.md",
			setupFence:  false,
			expectFence: true,
			validateOutput: func(content string) bool {
				return strings.HasPrefix(content, "```mermaid\n") &&
					strings.HasSuffix(content, "```\n")
			},
		},
		{
			name:        "Save as text file with fencing enabled",
			filename:    "flowchart.txt",
			setupFence:  true,
			expectFence: true,
			validateOutput: func(content string) bool {
				return strings.HasPrefix(content, "```mermaid\n") &&
					strings.HasSuffix(content, "```\n")
			},
		},
		{
			name:        "Save to nested directory",
			filename:    "nested/dir/flowchart.txt",
			setupFence:  false,
			expectFence: false,
			validateOutput: func(content string) bool {
				return strings.Contains(content, "Test Flowchart")
			},
		},
		{
			name:        "Save with invalid path",
			filename:    string([]byte{0}),
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setupFence {
				flowchart.EnableMarkdownFence()
			} else {
				flowchart.DisableMarkdownFence()
			}

			path := filepath.Join(tempDir, tt.filename)
			err := flowchart.RenderToFile(path)

			if tt.expectError {
				if err == nil {
					t.Error("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			content, err := os.ReadFile(path)
			if err != nil {
				t.Fatalf("Failed to read output file: %v", err)
			}

			if tt.validateOutput != nil {
				if !tt.validateOutput(string(content)) {
					t.Error("Output validation failed")
				}
			}

			if flowchart.markdownFence != tt.setupFence {
				t.Error("Flowchart fence state was permanently modified")
			}
		})
	}
}

func TestFlowchart_AddNode(t *testing.T) {
	tests := []struct {
		name      string
		flowchart *Flowchart
		text      string
		wantNode  *Node
	}{
		{
			name:      "Add simple node",
			flowchart: NewFlowchart(),
			text:      "Test Node",
			wantNode: &Node{
				ID:    0,
				Text:  "Test Node",
				Shape: NodeShapeRoundEdges,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.flowchart.AddNode(tt.text)

			if !reflect.DeepEqual(got, tt.wantNode) {
				t.Errorf("AddNode() = %v, want %v", got, tt.wantNode)
			}

			if len(tt.flowchart.nodes) != 1 || !reflect.DeepEqual(tt.flowchart.nodes[0], got) {
				t.Errorf("Node not added to flowchart correctly")
			}
		})
	}
}

func TestFlowchart_AddLink(t *testing.T) {
	tests := []struct {
		name      string
		flowchart *Flowchart
		setup     func(*Flowchart) (*Node, *Node)
		wantLink  *Link
	}{
		{
			name:      "Add simple link",
			flowchart: NewFlowchart(),
			setup: func(f *Flowchart) (*Node, *Node) {
				from := f.AddNode("Start")
				to := f.AddNode("End")
				return from, to
			},
			wantLink: &Link{
				Shape:  LinkShapeOpen,
				Head:   LinkArrowTypeArrow,
				Tail:   LinkArrowTypeNone,
				Length: 0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			from, to := tt.setup(tt.flowchart)
			got := tt.flowchart.AddLink(from, to)

			// Update expected link with actual nodes
			tt.wantLink.From = from
			tt.wantLink.To = to

			if !reflect.DeepEqual(got, tt.wantLink) {
				t.Errorf("AddLink() = %v, want %v", got, tt.wantLink)
			}

			if len(tt.flowchart.links) != 1 || !reflect.DeepEqual(tt.flowchart.links[0], got) {
				t.Errorf("Link not added to flowchart correctly")
			}
		})
	}
}
