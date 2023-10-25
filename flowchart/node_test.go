package flowchart

import (
	"reflect"
	"testing"
)

func TestNewNode(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)

	type args struct {
		id   uint64
		text string
	}
	tests := []struct {
		name        string
		args        args
		wantNewNode *Node
	}{
		{
			name: "Nominal test",
			args: args{
				id:   123,
				text: "Test",
			},
			wantNewNode: &Node{
				ID:    123,
				Text:  "Test",
				Shape: NodeShapeRoundEdges,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewNode := NewNode(tt.args.id, tt.args.text); !reflect.DeepEqual(gotNewNode, tt.wantNewNode) {
				t.Errorf("NewNode() = %v, want %v", gotNewNode, tt.wantNewNode)
			}
		})
	}
}

func TestNode_String(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)

	tests := []struct {
		name string
		node *Node
		want string
	}{
		{
			name: "Nominal test",
			node: &Node{
				ID:    123,
				Shape: NodeShapeRoundEdges,
				Text:  "Some text",
			},
			want: "\t123(\"Some text\")\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.node.String(); got != tt.want {
				t.Errorf("Node.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
