package main

import (
	"fmt"

	"github.com/TyphonHill/go-mermaid/flowchart"
)

func main() {
	fc := flowchart.NewFlowchart()
	fc.Title = "Subgraphs"
	fc.Direction = flowchart.FlowchartDirectionTopToBottom

	node1 := fc.AddNode("a1")
	node2 := fc.AddNode("a2")
	node3 := fc.AddNode("b1")
	node4 := fc.AddNode("b2")
	node5 := fc.AddNode("c1")
	node6 := fc.AddNode("c2")

	subgraph := fc.AddSubgraph("one")
	subgraph.AddLink(node1, node2)

	subgraph = fc.AddSubgraph("two")
	subgraph.AddLink(node3, node4)

	subgraph = fc.AddSubgraph("three")
	subgraph.AddLink(node5, node6)
	subgraph.AddLink(node5, node2)

	fmt.Println(fc.String())
}
