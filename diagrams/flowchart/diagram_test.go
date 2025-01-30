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
				Direction:  FlowchartDirectionTopToBottom,
				CurveStyle: CurveStyleNone,
				classes:    make([]*Class, 0),
				nodes:      make([]*Node, 0),
				subgraphs:  make([]*Subgraph, 0),
				links:      make([]*Link, 0),
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
			if !tt.flowchart.IsMarkdownFenceEnabled() {
				t.Error("EnableMarkdownFence() did not enable markdown fence")
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
			if tt.flowchart.IsMarkdownFenceEnabled() {
				t.Error("DisableMarkdownFence() did not disable markdown fence")
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
		{
			name:      "Flowchart with curve style",
			flowchart: NewFlowchart(),
			setup: func(f *Flowchart) {
				f.SetCurveStyle(CurveStyleBasis)
			},
			contains: []string{
				"%%{ init: { 'flowchart': { 'curve': 'basis' } } }%%\n",
				"flowchart TB\n",
			},
		},
		{
			name:      "Flowchart with class",
			flowchart: NewFlowchart(),
			setup: func(f *Flowchart) {
				class := f.AddClass("testClass")
				class.Style.Fill = "#f9f9f9"
			},
			contains: []string{
				"flowchart TB\n",
				"classDef testClass fill:#f9f9f9",
			},
		},
		{
			name:      "Flowchart with subgraph",
			flowchart: NewFlowchart(),
			setup: func(f *Flowchart) {
				subgraph := f.AddSubgraph("Test Subgraph")
				node1 := f.AddNode("Node1")
				node2 := f.AddNode("Node2")
				subgraph.AddLink(node1, node2)
			},
			contains: []string{
				"flowchart TB\n",
				"subgraph 0 [Test Subgraph]",
				"1(\"Node1\")",
				"2(\"Node2\")",
				"1 --> 2",
				"end",
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
					t.Errorf("Flowchart.String() output missing expected content: %v, got %v", want, got)
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

			if flowchart.IsMarkdownFenceEnabled() != tt.setupFence {
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

func TestFlowchart_AddSubgraph(t *testing.T) {
	flowchart := NewFlowchart()

	subgraph := flowchart.AddSubgraph("Test Subgraph")
	if subgraph == nil {
		t.Error("AddSubgraph() returned nil")
	}

	if len(flowchart.subgraphs) != 1 {
		t.Errorf("AddSubgraph() resulted in %d subgraphs, want 1", len(flowchart.subgraphs))
	}

	if subgraph.Title != "Test Subgraph" {
		t.Errorf("AddSubgraph() created subgraph with title %q, want %q", subgraph.Title, "Test Subgraph")
	}

	// Test ID generation
	subgraph2 := flowchart.AddSubgraph("Second Subgraph")
	if subgraph2.ID <= subgraph.ID {
		t.Errorf("Second subgraph ID %d should be greater than first subgraph ID %d", subgraph2.ID, subgraph.ID)
	}
}

func TestFlowchart_AddClass(t *testing.T) {
	flowchart := NewFlowchart()

	class := flowchart.AddClass("TestClass")
	if class == nil {
		t.Error("AddClass() returned nil")
	}

	if len(flowchart.classes) != 1 {
		t.Errorf("AddClass() resulted in %d classes, want 1", len(flowchart.classes))
	}

	if class.Name != "TestClass" {
		t.Errorf("AddClass() created class with name %q, want %q", class.Name, "TestClass")
	}

	// Test multiple classes
	class2 := flowchart.AddClass("SecondClass")
	if class2.Name != "SecondClass" {
		t.Errorf("AddClass() created second class with name %q, want %q", class2.Name, "SecondClass")
	}

	if len(flowchart.classes) != 2 {
		t.Errorf("After adding second class, got %d classes, want 2", len(flowchart.classes))
	}
}

func TestFlowchart_SetDirection(t *testing.T) {
	tests := []struct {
		name      string
		direction flowchartDirection
	}{
		{"Top to Bottom", FlowchartDirectionTopToBottom},
		{"Left to Right", FlowchartDirectionLeftRight},
		{"Right to Left", FlowchartDirectionRightLeft},
		{"Bottom Up", FlowchartDirectionBottomUp},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flowchart := NewFlowchart()
			result := flowchart.SetDirection(tt.direction)

			if flowchart.Direction != tt.direction {
				t.Errorf("SetDirection() = %v, want %v", flowchart.Direction, tt.direction)
			}

			if result != flowchart {
				t.Error("SetDirection() should return flowchart for chaining")
			}
		})
	}
}

func TestFlowchart_SetCurveStyle(t *testing.T) {
	tests := []struct {
		name  string
		style curveStyle
	}{
		{"Basis", CurveStyleBasis},
		{"Linear", CurveStyleLinear},
		{"Natural", CurveStyleNatural},
		{"Step", CurveStyleStep},
		{"None", CurveStyleNone},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flowchart := NewFlowchart()
			result := flowchart.SetCurveStyle(tt.style)

			if flowchart.CurveStyle != tt.style {
				t.Errorf("SetCurveStyle() = %v, want %v", flowchart.CurveStyle, tt.style)
			}

			if result != flowchart {
				t.Error("SetCurveStyle() should return flowchart for chaining")
			}
		})
	}
}

func TestLink_SetText(t *testing.T) {
	link := NewLink(nil, nil)
	result := link.SetText("Test Text")

	if link.Text != "Test Text" {
		t.Errorf("SetText() = %v, want %v", link.Text, "Test Text")
	}

	if result != link {
		t.Error("SetText() should return link for chaining")
	}
}

func TestLink_SetShape(t *testing.T) {
	tests := []struct {
		name  string
		shape linkShape
	}{
		{"Open", LinkShapeOpen},
		{"Dotted", LinkShapeDotted},
		{"Thick", LinkShapeThick},
		{"Invisible", LinkShapeInvisible},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			link := NewLink(nil, nil)
			result := link.SetShape(tt.shape)

			if link.Shape != tt.shape {
				t.Errorf("SetShape() = %v, want %v", link.Shape, tt.shape)
			}

			if result != link {
				t.Error("SetShape() should return link for chaining")
			}
		})
	}
}

func TestLink_SetLength(t *testing.T) {
	link := NewLink(nil, nil)
	result := link.SetLength(5)

	if link.Length != 5 {
		t.Errorf("SetLength() = %v, want %v", link.Length, 5)
	}

	if result != link {
		t.Error("SetLength() should return link for chaining")
	}
}

func TestLink_SetHead(t *testing.T) {
	tests := []struct {
		name      string
		arrowType linkArrowType
	}{
		{"None", LinkArrowTypeNone},
		{"Arrow", LinkArrowTypeArrow},
		{"Left Arrow", LinkArrowTypeLeftArrow},
		{"Bullet", LinkArrowTypeBullet},
		{"Cross", LinkArrowTypeCross},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			link := NewLink(nil, nil)
			result := link.SetHead(tt.arrowType)

			if link.Head != tt.arrowType {
				t.Errorf("SetHead() = %v, want %v", link.Head, tt.arrowType)
			}

			if result != link {
				t.Error("SetHead() should return link for chaining")
			}
		})
	}
}

func TestLink_SetTail(t *testing.T) {
	tests := []struct {
		name      string
		arrowType linkArrowType
	}{
		{"None", LinkArrowTypeNone},
		{"Arrow", LinkArrowTypeArrow},
		{"Left Arrow", LinkArrowTypeLeftArrow},
		{"Bullet", LinkArrowTypeBullet},
		{"Cross", LinkArrowTypeCross},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			link := NewLink(nil, nil)
			result := link.SetTail(tt.arrowType)

			if link.Tail != tt.arrowType {
				t.Errorf("SetTail() = %v, want %v", link.Tail, tt.arrowType)
			}

			if result != link {
				t.Error("SetTail() should return link for chaining")
			}
		})
	}
}
