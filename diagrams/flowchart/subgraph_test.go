package flowchart

import (
	"reflect"
	"strings"
	"testing"
)

func TestNewSubgraph(t *testing.T) {
	tests := []struct {
		name         string
		id           uint64
		title        string
		wantSubgraph *Subgraph
	}{
		{
			name:  "Create new subgraph with basic title",
			id:    1,
			title: "Test Subgraph",
			wantSubgraph: &Subgraph{
				ID:        1,
				Title:     "Test Subgraph",
				Direction: SubgraphDirectionNone,
			},
		},
		{
			name:  "Create new subgraph with empty title",
			id:    2,
			title: "",
			wantSubgraph: &Subgraph{
				ID:        2,
				Title:     "",
				Direction: SubgraphDirectionNone,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSubgraph(tt.id, tt.title)
			if !reflect.DeepEqual(got, tt.wantSubgraph) {
				t.Errorf("NewSubgraph() = %v, want %v", got, tt.wantSubgraph)
			}
		})
	}
}

func TestSubgraph_AddSubgraph(t *testing.T) {
	tests := []struct {
		name       string
		subgraph   *Subgraph
		title      string
		wantLength int
	}{
		{
			name: "Add subgraph to empty subgraph",
			subgraph: func() *Subgraph {
				sg := NewSubgraph(1, "Parent")
				sg.idGenerator = &MockIDGenerator{}
				return sg
			}(),
			title:      "Child",
			wantLength: 1,
		},
		{
			name: "Add multiple subgraphs",
			subgraph: func() *Subgraph {
				sg := NewSubgraph(1, "Parent")
				sg.idGenerator = &MockIDGenerator{}
				return sg
			}(),
			title:      "Child",
			wantLength: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.subgraph.AddSubgraph(tt.title)

			if got == nil {
				t.Error("AddSubgraph() returned nil")
			}

			if len(tt.subgraph.subgraphs) != tt.wantLength {
				t.Errorf("AddSubgraph() resulted in length = %v, want %v",
					len(tt.subgraph.subgraphs), tt.wantLength)
			}

			if got.Title != tt.title {
				t.Errorf("AddSubgraph() title = %v, want %v", got.Title, tt.title)
			}
		})
	}
}

func TestSubgraph_AddLink(t *testing.T) {
	node1 := NewNode(1, "Start")
	node2 := NewNode(2, "End")

	tests := []struct {
		name       string
		subgraph   *Subgraph
		from       *Node
		to         *Node
		wantLength int
	}{
		{
			name:       "Add link to empty subgraph",
			subgraph:   NewSubgraph(1, "Test"),
			from:       node1,
			to:         node2,
			wantLength: 1,
		},
		{
			name:       "Add multiple links",
			subgraph:   NewSubgraph(1, "Test"),
			from:       node1,
			to:         node2,
			wantLength: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.subgraph.AddLink(tt.from, tt.to)

			if got == nil {
				t.Error("AddLink() returned nil")
			}

			if len(tt.subgraph.links) != tt.wantLength {
				t.Errorf("AddLink() resulted in length = %v, want %v",
					len(tt.subgraph.links), tt.wantLength)
			}

			if !reflect.DeepEqual(got.From, tt.from) || !reflect.DeepEqual(got.To, tt.to) {
				t.Errorf("AddLink() nodes do not match expected values")
			}
		})
	}
}

func TestSubgraph_String(t *testing.T) {
	tests := []struct {
		name        string
		subgraph    *Subgraph
		setup       func(*Subgraph)
		indentation string
		contains    []string
	}{
		{
			name:        "Empty subgraph",
			subgraph:    NewSubgraph(1, "Test"),
			indentation: "%s",
			contains: []string{
				"\tsubgraph 1 [Test]",
				"\tend",
			},
		},
		{
			name:     "Subgraph with direction",
			subgraph: NewSubgraph(1, "Test"),
			setup: func(s *Subgraph) {
				s.Direction = SubgraphDirectionLeftRight
			},
			indentation: "%s",
			contains: []string{
				"\tsubgraph 1 [Test]",
				"\t\tdirection LR",
				"\tend",
			},
		},
		{
			name: "Subgraph with nested subgraph",
			subgraph: func() *Subgraph {
				sg := NewSubgraph(1, "Parent")
				sg.idGenerator = &MockIDGenerator{}
				return sg
			}(),
			setup: func(s *Subgraph) {
				s.AddSubgraph("Child")
			},
			indentation: "%s",
			contains: []string{
				"\tsubgraph 1 [Parent]",
				"\tsubgraph 0 [Child]",
				"\tend",
			},
		},
		{
			name:     "Subgraph with links",
			subgraph: NewSubgraph(1, "Test"),
			setup: func(s *Subgraph) {
				node1 := NewNode(1, "Start")
				node2 := NewNode(2, "End")
				s.AddLink(node1, node2)
			},
			indentation: "\t",
			contains: []string{
				"\tsubgraph 1 [Test]",
				"\t\t1 --> 2\n",
				"\tend",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup(tt.subgraph)
			}

			got := tt.subgraph.String(tt.indentation)
			t.Logf("Generated output:\n%s", got)

			for _, want := range tt.contains {
				if !strings.Contains(got, want) {
					t.Errorf("Subgraph.String() missing expected content: %v\nFull output:\n%v", want, got)
				}
			}
		})
	}
}
