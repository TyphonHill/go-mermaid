package main

import (
	"fmt"

	"github.com/BruceNocentini13/go-mermaid/flowchart"
)

func main() {
	test := flowchart.NewFlowchart()
	test.Title = "Testing"
	test.CurveStyle = flowchart.CurveStyleLinear

	class := test.AddClass("NewClass")
	class.Style = flowchart.NewNodeStyle()
	class.Style.StrokeWidth = 1
	class.Style.Fill = "#f9f"
	class.Style.Color = "#fff"
	class.Style.StrokeDash = "5"

	node1 := test.AddNode("1")
	node1.Text = "Start"
	node1.Class = class

	node2 := test.AddNode("2")
	node2.Text = "Is it?"
	node2.Shape = flowchart.NodeShapeRhombus
	node2.Class = class
	node2.Style = flowchart.NewNodeStyle()

	node3 := test.AddNode("3")
	node3.Text = "OK"

	node4 := test.AddNode("4")
	node4.Text = "Rethink#9829;"

	node5 := test.AddNode("5")
	node5.Text = "End"

	subgraph := test.AddSubgraph("123")
	subgraph.Title = "Title"
	subgraph.Direction = flowchart.SubgraphDirectionTopToBottom
	subgraph.AddLink(node1, node2)

	subgraph = test.AddSubgraph("456")
	subgraph.Title = "Title2"
	subgraph.Direction = flowchart.SubgraphDirectionBottomUp
	subgraph.AddLink(node4, node2)

	subgraph = subgraph.AddSubgraph("098")
	subgraph.Title = "Inside"
	subgraph.Direction = flowchart.SubgraphDirectionLeftRight
	subgraph.AddLink(node3, node5)

	fmt.Println(test.String())

	// test := sequencediagram.NewSequenceDiagram()
	// test.Title = "Testing"

	// part := test.AddParticipant("Alice")
	// part.Alias = "A"

	// part = test.AddParticipant("Bob")
	// part.Alias = "B"

	// fmt.Println(test.String())
}
