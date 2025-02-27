package main

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/TyphonHill/go-mermaid/diagrams/timeline"
)

func main() {
	// Create a new timeline diagram
	diagram := timeline.NewDiagram()
	diagram.EnableMarkdownFence()
	diagram.SetTitle("Simple Project Timeline")
	diagram.DisableMultiColot()

	// Add a Planning section with events
	planning := diagram.AddSection("Planning")
	e := planning.AddEvent("2024-01", "Project kickoff meeting with stakeholders")
	e.AddSubEvent("Initial requirements gathering")
	e = planning.AddEvent("2024-02", "Budget and resource allocation")
	e.AddSubEvent("Project plan finalization")

	// Add a Development section with events
	development := diagram.AddSection("Development")
	e = development.AddEvent("2024-03", "Setup development environment")
	e.AddSubEvent("Core feature implementation")
	development.AddEvent("2024-04", "")

	e = development.AddEvent("2024-05", "Staging environment deployment")
	e.AddSubEvent("User acceptance testing")
	e = development.AddEvent("2024-06", "Production deployment")
	e.AddSubEvent("Post-deployment monitoring")

	// Write the diagram to README.md in the same directory as this source file
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	readmePath := filepath.Join(dir, "README.md")
	if err := diagram.RenderToFile(readmePath); err != nil {
		fmt.Printf("Error writing diagram to README.md: %v\n", err)
		return
	}
}
