package main

import (
	"fmt"

	"github.com/BruceNocentini13/go-mermaid/flowchart"
)

func main() {
	fc := flowchart.NewFlowchart()
	fc.Title = "Simple Flowchart"

	node1 := fc.AddNode("Start")
	node2 := fc.AddNode("Is It?")
	node3 := fc.AddNode("OK")
	node4 := fc.AddNode("Rethink")
	node5 := fc.AddNode("End")

	link := fc.AddLink(node1, node2)
	link.Shape = flowchart.LinkShapeDotted

	link = fc.AddLink(node2, node3)
	link.Shape = flowchart.LinkShapeOpen
	link.Text = "Yes"

	link = fc.AddLink(node3, node4)
	link.Shape = flowchart.LinkShapeOpen

	link = fc.AddLink(node4, node2)
	link.Shape = flowchart.LinkShapeOpen

	link = fc.AddLink(node2, node5)
	link.Shape = flowchart.LinkShapeOpen
	link.Text = "No"
	link.Length = 1

	fmt.Println(fc.String())
}
