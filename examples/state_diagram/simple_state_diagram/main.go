package main

import (
	"fmt"
	"path/filepath"
	"runtime"

	stateDiagram "github.com/TyphonHill/go-mermaid/diagrams/state"
)

func main() {
	diagram := stateDiagram.NewDiagram()
	diagram.EnableMarkdownFence()

	diagram.Title = "Simple State Machine"

	// Add states
	idle := diagram.AddState("idle", "Idle State", stateDiagram.StateNormal)
	processing := diagram.AddState("processing", "Processing", stateDiagram.StateNormal)
	error := diagram.AddState("error", "Error", stateDiagram.StateNormal)

	// Add transitions
	diagram.AddTransition(idle, processing, "Start Process")
	diagram.AddTransition(processing, idle, "Complete")
	diagram.AddTransition(processing, error, "Error Occurred")
	diagram.AddTransition(error, idle, "Reset")

	// Write the diagram to README.md in the same directory as this source file
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	readmePath := filepath.Join(dir, "README.md")
	if err := diagram.RenderToFile(readmePath); err != nil {
		fmt.Printf("Error writing diagram to README.md: %v\n", err)
		return
	}
}
