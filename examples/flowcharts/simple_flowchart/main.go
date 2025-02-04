package main

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/TyphonHill/go-mermaid/diagrams/flowchart"
)

func main() {
	// Create a new flowchart
	diagram := flowchart.NewFlowchart()
	diagram.SetTitle("Simple Process Flow")

	// Add nodes with different shapes
	start := diagram.AddNode("Start")
	start.SetShape(flowchart.NodeShapeTerminal)

	input := diagram.AddNode("Get User Input")
	input.SetShape(flowchart.NodeShapeManualInput)

	process := diagram.AddNode("Process Data")
	process.SetShape(flowchart.NodeShapeProcess)

	decision := diagram.AddNode("Valid?")
	decision.SetShape(flowchart.NodeShapeDecision)

	output := diagram.AddNode("Display Result")
	output.SetShape(flowchart.NodeShapeDisplay)

	end := diagram.AddNode("End")
	end.SetShape(flowchart.NodeShapeTerminal)

	// Add links between nodes
	diagram.AddLink(start, input)
	diagram.AddLink(input, process)
	diagram.AddLink(process, decision)
	diagram.AddLink(decision, output)
	diagram.AddLink(decision, input)
	diagram.AddLink(output, end)

	// Write the diagram to README.md
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	readmePath := filepath.Join(dir, "README.md")
	if err := diagram.RenderToFile(readmePath); err != nil {
		fmt.Printf("Error writing diagram to README.md: %v\n", err)
		return
	}
}
