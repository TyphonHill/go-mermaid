package main

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/TyphonHill/go-mermaid/diagrams/flowchart"
)

func main() {
	diagram := flowchart.NewFlowchart()
	diagram.SetTitle("Software Development Process")

	// Define nodes with various shapes
	start := diagram.AddNode("Start Project")
	start.SetShape(flowchart.NodeShapeTerminal)

	requirements := diagram.AddNode("Gather Requirements")
	requirements.SetShape(flowchart.NodeShapeDocument)

	design := diagram.AddNode("System Design")
	design.SetShape(flowchart.NodeShapeProcess)

	database := diagram.AddNode("Database Design")
	database.SetShape(flowchart.NodeShapeDatabase)

	coding := diagram.AddNode("Implementation")
	coding.SetShape(flowchart.NodeShapeProcess)

	testing := diagram.AddNode("Testing")
	testing.SetShape(flowchart.NodeShapePrepare)

	bugs := diagram.AddNode("Bugs Found?")
	bugs.SetShape(flowchart.NodeShapeDecision)

	deploy := diagram.AddNode("Deployment")
	deploy.SetShape(flowchart.NodeShapeInternalStorage)

	monitor := diagram.AddNode("Monitoring")
	monitor.SetShape(flowchart.NodeShapeDisplay)

	end := diagram.AddNode("End")
	end.SetShape(flowchart.NodeShapeTerminal)

	// Add links
	diagram.AddLink(start, requirements)
	diagram.AddLink(requirements, design)
	diagram.AddLink(design, database)
	diagram.AddLink(database, coding)
	diagram.AddLink(coding, testing)
	diagram.AddLink(testing, bugs)
	diagram.AddLink(bugs, coding)
	diagram.AddLink(bugs, deploy)
	diagram.AddLink(deploy, monitor)
	diagram.AddLink(monitor, end)

	// Write the diagram to README.md
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	readmePath := filepath.Join(dir, "README.md")
	if err := diagram.RenderToFile(readmePath); err != nil {
		fmt.Printf("Error writing diagram to README.md: %v\n", err)
		return
	}
}
