package flowchart

import (
	"reflect"
	"testing"
)

func TestNewFlowchart(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)

	tests := []struct {
		name             string
		wantNewFlowchart *Flowchart
	}{
		{
			name: "Nominal test",
			wantNewFlowchart: &Flowchart{
				Direction:  FlowchartDirectionTopToBottom,
				CurveStyle: CurveStyleNone,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewFlowchart := NewFlowchart(); !reflect.DeepEqual(gotNewFlowchart, tt.wantNewFlowchart) {
				t.Errorf("NewFlowchart() = %v, want %v", gotNewFlowchart, tt.wantNewFlowchart)
			}
		})
	}
}

func TestFlowchart_String(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)

	testClass := NewClass("TestClass")
	testClass.Style = NewNodeStyle()

	testNode := NewNode(123, "Test Node")
	testNode.Shape = NodeShapeCircle
	testNode.Class = testClass

	testNode2 := NewNode(456, "Test Node 2")
	testNode2.Shape = NodeShapeParallelogramAlt
	testNode2.Style = NewNodeStyle()

	testLink := NewLink(testNode, testNode2)
	testLink.Length = 5
	testLink.Shape = LinkShapeThick
	testLink.Text = "Link Text"

	testSubgraph := NewSubgraph(999, "Test")
	testSubgraph.Title = "Title"
	testSubgraph.AddLink(testNode, testNode2)

	tests := []struct {
		name      string
		flowchart *Flowchart
		want      string
	}{
		{
			name: "Nominal test",
			flowchart: &Flowchart{
				Title:      "Test flowchart",
				Direction:  FlowchartDirectionLeftRight,
				CurveStyle: CurveStyleLinear,
				nodes: []*Node{
					testNode,
					testNode2,
				},
				links: []*Link{
					testLink,
				},
				classes: []*Class{
					testClass,
				},
				subgraphs: []*Subgraph{
					testSubgraph,
				},
			},
			want: `---
title: Test flowchart
---

%%{ init: { 'flowchart': { 'curve': 'linear' } } }%%
flowchart LR
	classDef TestClass stroke-width:1,stroke-dasharray:0
	123(("Test Node")):::TestClass
	456[\"Test Node 2"\]
	style 456 stroke-width:1,stroke-dasharray:0
	subgraph 999 [Title]
		123 --> 456
	end
	123 =======>|Link Text| 456`,
		},
		{
			name: "No title",
			flowchart: &Flowchart{
				Direction:  FlowchartDirectionLeftRight,
				CurveStyle: CurveStyleLinear,
				nodes: []*Node{
					testNode,
					testNode2,
				},
				links: []*Link{
					testLink,
				},
				classes: []*Class{
					testClass,
				},
			},
			want: `%%{ init: { 'flowchart': { 'curve': 'linear' } } }%%
flowchart LR
	classDef TestClass stroke-width:1,stroke-dasharray:0
	123(("Test Node")):::TestClass
	456[\"Test Node 2"\]
	style 456 stroke-width:1,stroke-dasharray:0
	123 =======>|Link Text| 456`,
		},
		{
			name: "CurveStyle is CurveStyleNone",
			flowchart: &Flowchart{
				Title:      "Test flowchart",
				Direction:  FlowchartDirectionLeftRight,
				CurveStyle: CurveStyleNone,
				nodes: []*Node{
					testNode,
					testNode2,
				},
				links: []*Link{
					testLink,
				},
				classes: []*Class{
					testClass,
				},
			},
			want: `---
title: Test flowchart
---

flowchart LR
	classDef TestClass stroke-width:1,stroke-dasharray:0
	123(("Test Node")):::TestClass
	456[\"Test Node 2"\]
	style 456 stroke-width:1,stroke-dasharray:0
	123 =======>|Link Text| 456`,
		},
		{
			name: "No classes",
			flowchart: &Flowchart{
				Title:      "Test flowchart",
				Direction:  FlowchartDirectionLeftRight,
				CurveStyle: CurveStyleLinear,
				nodes: []*Node{
					testNode,
					testNode2,
				},
				links: []*Link{
					testLink,
				},
				classes: []*Class{},
			},
			want: `---
title: Test flowchart
---

%%{ init: { 'flowchart': { 'curve': 'linear' } } }%%
flowchart LR
	123(("Test Node")):::TestClass
	456[\"Test Node 2"\]
	style 456 stroke-width:1,stroke-dasharray:0
	123 =======>|Link Text| 456`,
		},
		{
			name: "No Nodes",
			flowchart: &Flowchart{
				Title:      "Test flowchart",
				Direction:  FlowchartDirectionLeftRight,
				CurveStyle: CurveStyleLinear,
				nodes:      []*Node{},
				links: []*Link{
					testLink,
				},
				classes: []*Class{},
			},
			want: `---
title: Test flowchart
---

%%{ init: { 'flowchart': { 'curve': 'linear' } } }%%
flowchart LR
	123 =======>|Link Text| 456`,
		},
		{
			name: "No Links",
			flowchart: &Flowchart{
				Title:      "Test flowchart",
				Direction:  FlowchartDirectionLeftRight,
				CurveStyle: CurveStyleLinear,
				nodes: []*Node{
					testNode,
					testNode2,
				},
				links: []*Link{},
				classes: []*Class{
					testClass,
				},
			},
			want: `---
title: Test flowchart
---

%%{ init: { 'flowchart': { 'curve': 'linear' } } }%%
flowchart LR
	classDef TestClass stroke-width:1,stroke-dasharray:0
	123(("Test Node")):::TestClass
	456[\"Test Node 2"\]
	style 456 stroke-width:1,stroke-dasharray:0
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.flowchart.String(); got != tt.want {
				t.Errorf("Flowchart.String() = \n%v, want %v", got, tt.want)
			}

		})
	}
}

func TestFlowchart_AddNode(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)

	type args struct {
		text string
	}
	tests := []struct {
		name        string
		flowchart   *Flowchart
		args        args
		wantNewNode *Node
	}{
		{
			name: "Nominal test",
			flowchart: &Flowchart{
				Title:      "Flowchartt title",
				Direction:  FlowchartDirectionTopDown,
				CurveStyle: CurveStyleNone,
			},
			args: args{
				text: "Test",
			},
			wantNewNode: &Node{
				Text:  "Test",
				Shape: NodeShapeRoundEdges,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewNode := tt.flowchart.AddNode(tt.args.text); !reflect.DeepEqual(gotNewNode, tt.wantNewNode) {
				t.Errorf("Flowchart.AddNode() = %v, want %v", gotNewNode, tt.wantNewNode)
			}
		})
	}
}

func TestFlowchart_AddLink(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)

	type args struct {
		from *Node
		to   *Node
	}
	tests := []struct {
		name        string
		flowchart   *Flowchart
		args        args
		wantNewLink *Link
	}{
		{
			name: "Nominal test",
			flowchart: &Flowchart{
				Title:      "Flowchartt title",
				Direction:  FlowchartDirectionTopDown,
				CurveStyle: CurveStyleNone,
			},
			args: args{
				from: &Node{ID: 123},
				to:   &Node{ID: 456},
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
			if gotNewLink := tt.flowchart.AddLink(tt.args.from, tt.args.to); !reflect.DeepEqual(gotNewLink, tt.wantNewLink) {
				t.Errorf("Flowchart.AddLink() = %v, want %v", gotNewLink, tt.wantNewLink)
			}
		})
	}
}

func TestFlowchart_AddClass(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)

	type args struct {
		name string
	}
	tests := []struct {
		name         string
		flowchart    *Flowchart
		args         args
		wantNewClass *Class
	}{
		{
			name: "Nominal test",
			flowchart: &Flowchart{
				Title:      "Flowchartt title",
				Direction:  FlowchartDirectionTopDown,
				CurveStyle: CurveStyleNone,
			},
			args: args{
				name: "TestClass",
			},
			wantNewClass: &Class{
				Name: "TestClass",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewClass := tt.flowchart.AddClass(tt.args.name); !reflect.DeepEqual(gotNewClass, tt.wantNewClass) {
				t.Errorf("Flowchart.AddClass() = %v, want %v", gotNewClass, tt.wantNewClass)
			}
		})
	}
}

func TestFlowchart_AddSubgraph(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)

	type args struct {
		title string
	}
	tests := []struct {
		name            string
		flowchart       *Flowchart
		args            args
		wantNewSubgraph *Subgraph
	}{
		{
			name: "Nominal test",
			flowchart: &Flowchart{
				Title:      "Flowchartt title",
				Direction:  FlowchartDirectionTopDown,
				CurveStyle: CurveStyleNone,
			},
			args: args{
				title: "Test",
			},
			wantNewSubgraph: &Subgraph{
				ID:        0,
				Title:     "Test",
				Direction: SubgraphDirectionNone,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewSubgraph := tt.flowchart.AddSubgraph(tt.args.title); !reflect.DeepEqual(gotNewSubgraph, tt.wantNewSubgraph) {
				t.Errorf("Flowchart.AddSubgraph() = %v, want %v", gotNewSubgraph, tt.wantNewSubgraph)
			}
		})
	}
}
