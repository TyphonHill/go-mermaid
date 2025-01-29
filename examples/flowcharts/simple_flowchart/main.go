package main

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/TyphonHill/go-mermaid/diagrams/flowchart"
)

func main() {
	diagram := flowchart.NewFlowchart()
	diagram.EnableMarkdownFence()
	diagram.Title = "Simple Login Flow"

	// Create nodes
	start := diagram.AddNode("Start")
	login := diagram.AddNode("Login Form")
	validate := diagram.AddNode("Validate")
	success := diagram.AddNode("Success")
	failure := diagram.AddNode("Failure")
	end := diagram.AddNode("End")

	// Add connections
	diagram.AddLink(start, login)
	diagram.AddLink(login, validate)

	successLink := diagram.AddLink(validate, success)
	successLink.Text = "Valid"

	failureLink := diagram.AddLink(validate, failure)
	failureLink.Text = "Invalid"

	diagram.AddLink(success, end)
	diagram.AddLink(failure, login)

	// Write the diagram to README.md in the same directory as this source file
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	readmePath := filepath.Join(dir, "README.md")
	if err := diagram.RenderToFile(readmePath); err != nil {
		fmt.Printf("Error writing diagram to README.md: %v\n", err)
		return
	}
}
