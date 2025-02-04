package flowchart

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestNewNode(t *testing.T) {
	tests := []struct {
		name     string
		id       string
		text     string
		wantNode *Node
	}{
		{
			name: "Create new node with text",
			id:   "b0",
			text: "Test Node",
			wantNode: &Node{
				ID:    "b0",
				Text:  "Test Node",
				Shape: NodeShapeProcess,
			},
		},
		{
			name: "Create new node without text",
			id:   "b1",
			text: "",
			wantNode: &Node{
				ID:    "b1",
				Text:  "",
				Shape: NodeShapeProcess,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewNode(tt.id, tt.text)
			if !reflect.DeepEqual(got, tt.wantNode) {
				t.Errorf("NewNode() = %v, want %v", got, tt.wantNode)
			}
		})
	}
}

func TestNode_String(t *testing.T) {
	tests := []struct {
		name     string
		node     *Node
		setup    func(*Node)
		contains []string
	}{
		{
			name: "Basic node with default shape",
			node: NewNode("1", "Test Node"),
			contains: []string{
				"1@{ shape: rect label: \"Test Node\"}",
			},
		},
		{
			name: "Node with custom shape",
			node: NewNode("2", "Diamond Node"),
			setup: func(n *Node) {
				n.SetShape(NodeShapeDecision)
			},
			contains: []string{
				"2@{ shape: diam label: \"Diamond Node\"}",
			},
		},
		{
			name: "Node with class",
			node: NewNode("3", "Classy Node"),
			setup: func(n *Node) {
				n.SetClass(NewClass("highlight"))
			},
			contains: []string{
				"3@{ shape: rect label: \"Classy Node\"}:::highlight",
			},
		},
		{
			name: "Node with style",
			node: NewNode("4", "Styled Node"),
			setup: func(n *Node) {
				n.SetStyle(NewNodeStyle())
			},
			contains: []string{
				"4@{ shape: rect label: \"Styled Node\"}",
				"style 4 fill:#f9f",
			},
		},
		{
			name: "Complex node with shape, class and style",
			node: NewNode("5", "Complex Node"),
			setup: func(n *Node) {
				n.SetShape(NodeShapeDatabase)
				n.SetClass(NewClass("db"))
				n.SetStyle(NewNodeStyle())
			},
			contains: []string{
				"5@{ shape: cyl label: \"Complex Node\"}:::db",
				"style 5 fill:#ccf,stroke:#333",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup(tt.node)
			}

			result := tt.node.String()
			for _, want := range tt.contains {
				if !strings.Contains(result, want) {
					t.Errorf("String() missing expected content %q in:\n%v", want, result)
				}
			}
		})
	}
}

func TestNode_SetText(t *testing.T) {
	node := NewNode("1", "Initial")
	result := node.SetText("Updated Text")

	if node.Text != "Updated Text" {
		t.Errorf("SetText() = %v, want %v", node.Text, "Updated Text")
	}

	if result != node {
		t.Error("SetText() should return node for chaining")
	}
}

func TestNode_SetClass(t *testing.T) {
	node := NewNode("1", "Test")
	class := NewClass("TestClass")
	result := node.SetClass(class)

	if node.Class != class {
		t.Errorf("SetClass() = %v, want %v", node.Class, class)
	}

	if result != node {
		t.Error("SetClass() should return node for chaining")
	}
}

func TestNode_SetStyle(t *testing.T) {
	node := NewNode("1", "Test Node")
	style := NewNodeStyle()
	style.Fill = "#f9f9f9"
	style.Stroke = "#333"

	result := node.SetStyle(style)

	if result != node {
		t.Error("SetStyle() should return node for chaining")
	}
	if node.Style != style {
		t.Errorf("SetStyle() = %v, want %v", node.Style, style)
	}
}

func TestNode_SetShape(t *testing.T) {
	tests := []struct {
		name      string
		shape     nodeShape
		wantShape nodeShape
	}{
		{
			name:      "Set to terminal shape",
			shape:     NodeShapeTerminal,
			wantShape: NodeShapeTerminal,
		},
		{
			name:      "Set to database shape",
			shape:     NodeShapeDatabase,
			wantShape: NodeShapeDatabase,
		},
		{
			name:      "Set to decision shape",
			shape:     NodeShapeDecision,
			wantShape: NodeShapeDecision,
		},
		{
			name:      "Set to document shape",
			shape:     NodeShapeDocument,
			wantShape: NodeShapeDocument,
		},
		{
			name:      "Set to process shape",
			shape:     NodeShapeProcess,
			wantShape: NodeShapeProcess,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := NewNode("1", "Test")
			result := node.SetShape(tt.shape)

			if result != node {
				t.Error("SetShape() should return node for chaining")
			}
			if node.Shape != tt.wantShape {
				t.Errorf("SetShape() = %v, want %v", node.Shape, tt.wantShape)
			}

			// Verify the shape appears correctly in the string output
			output := node.String()
			expectedShape := fmt.Sprintf("shape: %s", tt.wantShape)
			if !strings.Contains(output, expectedShape) {
				t.Errorf("String() = %v, should contain %v", output, expectedShape)
			}
		})
	}
}
