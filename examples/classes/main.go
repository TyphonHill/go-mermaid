package main

import (
	"fmt"

	"github.com/BruceNocentini13/go-mermaid/flowchart"
)

func main() {
	fc := flowchart.NewFlowchart()
	fc.Title = "Classes"
	class1 := fc.AddClass("Class1")
	class1.Style.Color = "red"
	class1.Style.Fill = "white"
	class1.Style.StrokeWidth = 2
	class1.Style.StrokeDash = "5"

	class2 := fc.AddClass("Class2")
	class2.Style.Color = "white"
	class2.Style.Stroke = "#333"
	class2.Style.Fill = "#13f"

	node1 := fc.AddNode("Start")
	node1.Class = class1

	node2 := fc.AddNode("Is It?")
	node3 := fc.AddNode("OK")
	node3.Class = class2

	node4 := fc.AddNode("Rethink")
	node5 := fc.AddNode("End")
	node5.Class = class1

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
