package flowchart

import (
	"reflect"
	"testing"
)

func TestNewSubgraph(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)

	type args struct {
		id    uint64
		title string
	}
	tests := []struct {
		name            string
		args            args
		wantNewSubgraph *Subgraph
	}{
		{
			name: "Nominal test",
			args: args{
				id:    123,
				title: "Test",
			},
			wantNewSubgraph: &Subgraph{
				ID:        123,
				Title:     "Test",
				Direction: SubgraphDirectionNone,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewSubgraph := NewSubgraph(tt.args.id, tt.args.title); !reflect.DeepEqual(gotNewSubgraph, tt.wantNewSubgraph) {
				t.Errorf("NewSubgraph() = %v, want %v", gotNewSubgraph, tt.wantNewSubgraph)
			}
		})
	}
}

func TestSubgraph_AddSubgraph(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)

	type args struct {
		title string
	}
	tests := []struct {
		name            string
		subgraph        *Subgraph
		args            args
		wantNewSubgraph *Subgraph
	}{
		{
			name: "Nominal test",
			args: args{
				title: "Test",
			},
			subgraph: &Subgraph{},
			wantNewSubgraph: &Subgraph{
				ID:    0,
				Title: "Test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewSubgraph := tt.subgraph.AddSubgraph(tt.args.title); !reflect.DeepEqual(gotNewSubgraph, tt.wantNewSubgraph) {
				t.Errorf("Subgraph.AddSubgraph() = %v, want %v", gotNewSubgraph, tt.wantNewSubgraph)
			}
		})
	}
}

func TestSubgraph_AddLink(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)

	type args struct {
		from *Node
		to   *Node
	}
	tests := []struct {
		name        string
		subgraph    *Subgraph
		args        args
		wantNewLink *Link
	}{
		{
			name: "Nominal test",
			args: args{
				from: &Node{ID: 123},
				to:   &Node{ID: 456},
			},
			subgraph: &Subgraph{
				links: []*Link{
					{
						From:   &Node{ID: 123},
						To:     &Node{ID: 456},
						Shape:  LinkShapeOpen,
						Head:   LinkArrowTypeArrow,
						Tail:   LinkArrowTypeNone,
						Length: 0,
					},
				},
			},
			wantNewLink: &Link{
				From:   &Node{ID: 123},
				To:     &Node{ID: 456},
				Shape:  LinkShapeOpen,
				Head:   LinkArrowTypeArrow,
				Tail:   LinkArrowTypeNone,
				Length: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewLink := tt.subgraph.AddLink(tt.args.from, tt.args.to); !reflect.DeepEqual(gotNewLink, tt.wantNewLink) {
				t.Errorf("Subgraph.AddLink() = %v, want %v", gotNewLink, tt.wantNewLink)
			}
		})
	}
}

func TestSubgraph_String(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)

	type args struct {
		curIndentation string
	}
	tests := []struct {
		name     string
		subgraph *Subgraph
		args     args
		want     string
	}{
		{
			name: "Nominal test",
			args: args{
				curIndentation: "%s",
			},
			subgraph: &Subgraph{
				ID:        123,
				Title:     "TestSubgraph",
				Direction: SubgraphDirectionTopToBottom,
				links: []*Link{
					{
						From:   &Node{ID: 111},
						To:     &Node{ID: 222},
						Shape:  LinkShapeOpen,
						Head:   LinkArrowTypeArrow,
						Tail:   LinkArrowTypeNone,
						Length: 0,
					},
				},
				subgraphs: []*Subgraph{
					{
						ID:        456,
						Title:     "TestSubSubgraph",
						Direction: SubgraphDirectionTopToBottom,
						links: []*Link{
							{
								From:   &Node{ID: 333},
								To:     &Node{ID: 444},
								Shape:  LinkShapeOpen,
								Head:   LinkArrowTypeArrow,
								Tail:   LinkArrowTypeNone,
								Length: 0,
							},
						},
					},
				},
			},
			want: `	subgraph 123 [TestSubgraph]
		direction TB
		subgraph 456 [TestSubSubgraph]
			direction TB
			333 --> 444
		end
		111 --> 222
	end
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.subgraph.String(tt.args.curIndentation); got != tt.want {
				t.Errorf("Subgraph.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
