package flowchart

import (
	"reflect"
	"testing"
)

func TestNewNode(t *testing.T) {
	tests := []struct {
		name     string
		id       uint64
		text     string
		wantNode *Node
	}{
		{
			name: "Create new node with basic text",
			id:   1,
			text: "Test Node",
			wantNode: &Node{
				ID:    1,
				Text:  "Test Node",
				Shape: NodeShapeRoundEdges,
			},
		},
		{
			name: "Create new node with empty text",
			id:   2,
			text: "",
			wantNode: &Node{
				ID:    2,
				Text:  "",
				Shape: NodeShapeRoundEdges,
			},
		},
		{
			name: "Create new node with special characters",
			id:   3,
			text: "Test: Node & Special < > Characters",
			wantNode: &Node{
				ID:    3,
				Text:  "Test: Node & Special < > Characters",
				Shape: NodeShapeRoundEdges,
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
		name    string
		node    *Node
		setup   func(*Node)
		wantStr string
	}{
		{
			name: "Basic node with round edges",
			node: &Node{
				ID:    1,
				Text:  "Test",
				Shape: NodeShapeRoundEdges,
			},
			wantStr: "\t1(\"Test\")\n",
		},
		{
			name: "Node with stadium shape",
			node: &Node{
				ID:    2,
				Text:  "Test",
				Shape: NodeShapeStadium,
			},
			wantStr: "\t2([\"Test\"])\n",
		},
		{
			name: "Node with class",
			node: &Node{
				ID:    3,
				Text:  "Test",
				Shape: NodeShapeRoundEdges,
			},
			setup: func(n *Node) {
				class := NewClass("testClass")
				n.Class = class
			},
			wantStr: "\t3(\"Test\"):::testClass\n",
		},
		{
			name: "Node with style",
			node: &Node{
				ID:    4,
				Text:  "Test",
				Shape: NodeShapeRoundEdges,
			},
			setup: func(n *Node) {
				style := NewNodeStyle()
				style.Fill = "#f9f9f9"
				n.Style = style
			},
			wantStr: "\t4(\"Test\")\n\tstyle 4 fill:#f9f9f9,stroke-width:1,stroke-dasharray:0\n",
		},
		{
			name: "Node with all shapes",
			node: &Node{
				ID:    5,
				Text:  "Test",
				Shape: NodeShapeRoundEdges,
			},
			setup: func(n *Node) {
				shapes := []nodeShape{
					NodeShapeRoundEdges,
					NodeShapeStadium,
					NodeShapeSubRoutine,
					NodeShapeCylindrical,
					NodeShapeCircle,
					NodeShapeAsymmetric,
					NodeShapeRhombus,
					NodeShapeHexagon,
					NodeShapeParallelogram,
					NodeShapeParallelogramAlt,
					NodeShapeTrapezoid,
					NodeShapeTrapezoidAlt,
					NodeShapeDoubleCircle,
				}
				for _, shape := range shapes {
					n.Shape = shape
					got := n.String()
					if got == "" {
						t.Errorf("Node.String() with shape %v returned empty string", shape)
					}
				}
			},
			wantStr: "\t5(((\"Test\")))\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup(tt.node)
			}

			got := tt.node.String()
			if got != tt.wantStr {
				t.Errorf("Node.String() = %v, want %v", got, tt.wantStr)
			}
		})
	}
}

func TestNode_SetText(t *testing.T) {
	node := NewNode(1, "Initial")
	result := node.SetText("Updated Text")

	if node.Text != "Updated Text" {
		t.Errorf("SetText() = %v, want %v", node.Text, "Updated Text")
	}

	if result != node {
		t.Error("SetText() should return node for chaining")
	}
}

func TestNode_SetClass(t *testing.T) {
	node := NewNode(1, "Test")
	class := NewClass("TestClass")
	result := node.SetClass(class)

	if node.Class != class {
		t.Errorf("SetClass() = %v, want %v", node.Class, class)
	}

	if result != node {
		t.Error("SetClass() should return node for chaining")
	}
}
