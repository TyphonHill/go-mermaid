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
	diagram.SetTitle("Software Development Lifecycle")
	diagram.SetTheme(timeline.TimelineThemeDark)

	// Planning Phase
	planning := diagram.AddSection("Planning")
	planning.AddEvent("Week 1", "Initial project kickoff meeting")
	planning.AddEvent("", "Stakeholder interviews and requirement gathering")
	planning.AddEvent("Week 2", "Market research and competitor analysis")
	planning.AddEvent("", "Project scope definition and documentation")
	planning.AddEvent("Week 3", "Resource planning and team allocation")
	planning.AddEvent("", "Risk assessment and mitigation strategies")

	// Design Phase
	design := diagram.AddSection("Design")
	design.AddEvent("Week 4", "High-level system architecture design")
	design.AddEvent("", "Database schema and data flow modeling")
	design.AddEvent("Week 5", "UI/UX wireframes and user journey mapping")
	design.AddEvent("", "Security architecture planning")
	design.AddEvent("Week 6", "API design and documentation")
	design.AddEvent("", "Technical specification review")

	// Development Phase
	development := diagram.AddSection("Development")
	development.AddEvent("Sprint 1", "Core infrastructure setup")
	development.AddEvent("", "Basic user authentication")
	development.AddEvent("Sprint 2", "Core feature implementation")
	development.AddEvent("", "API integration and testing")
	development.AddEvent("Sprint 3", "UI implementation")
	development.AddEvent("", "Code review and optimization")

	// Testing Phase
	testing := diagram.AddSection("Testing")
	testing.AddEvent("Week 12", "Unit testing implementation")
	testing.AddEvent("", "Integration testing setup")
	testing.AddEvent("Week 13", "Performance testing and optimization")
	testing.AddEvent("", "Security testing and vulnerability assessment")
	testing.AddEvent("Week 14", "User acceptance testing coordination")
	testing.AddEvent("", "Bug fixing and regression testing")

	// Deployment Phase
	deployment := diagram.AddSection("Deployment")
	deployment.AddEvent("Week 15", "Staging environment setup and configuration")
	deployment.AddEvent("", "Production environment preparation")
	deployment.AddEvent("Week 16", "Database migration planning")
	deployment.AddEvent("", "Deployment automation setup")
	deployment.AddEvent("Week 17", "Production deployment execution")
	deployment.AddEvent("", "Post-deployment health checks")

	// Maintenance Phase
	maintenance := diagram.AddSection("Maintenance")
	maintenance.AddEvent("Month 1", "24/7 system monitoring setup")
	maintenance.AddEvent("", "Performance metrics tracking")
	maintenance.AddEvent("Month 2", "Regular security patches and updates")
	maintenance.AddEvent("", "User feedback collection and analysis")
	maintenance.AddEvent("Month 3", "Feature enhancement planning")
	maintenance.AddEvent("", "Documentation updates and maintenance")

	// Write the diagram to README.md in the same directory as this source file
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	readmePath := filepath.Join(dir, "README.md")
	if err := diagram.RenderToFile(readmePath); err != nil {
		fmt.Printf("Error writing diagram to README.md: %v\n", err)
		return
	}
}
