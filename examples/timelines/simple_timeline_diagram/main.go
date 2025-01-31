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
	diagram.SetTitle("Simple Project Timeline")

	// Add a Planning section with events
	planning := diagram.AddSection("Planning")
	planning.AddEvent("2024-01", "Project kickoff meeting with stakeholders")
	planning.AddSubEvent("Initial requirements gathering")
	planning.AddEvent("2024-02", "Budget and resource allocation")
	planning.AddSubEvent("Project plan finalization")

	// Add a Development section with events
	development := diagram.AddSection("Development")
	development.AddEvent("2024-03", "Setup development environment")
	development.AddSubEvent("Core feature implementation")
	development.AddEvent("2024-04", "Integration with external services")
	development.AddSubEvent("Code review and testing")

	// Add a Deployment section
	deployment := diagram.AddSection("Deployment")
	deployment.AddEvent("2024-05", "Staging environment deployment")
	deployment.AddSubEvent("User acceptance testing")
	deployment.AddEvent("2024-06", "Production deployment")
	deployment.AddSubEvent("Post-deployment monitoring")

	// Write the diagram to README.md in the same directory as this source file
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	readmePath := filepath.Join(dir, "README.md")
	if err := diagram.RenderToFile(readmePath); err != nil {
		fmt.Printf("Error writing diagram to README.md: %v\n", err)
		return
	}
}
