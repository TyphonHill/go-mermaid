package main

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/TyphonHill/go-mermaid/diagrams/block"
)

func main() {
	// Create a new block diagram
	diagram := block.NewDiagram()
	diagram.EnableMarkdownFence()
	diagram.SetColumns(3)

	// Add blocks with different shapes and arrows
	start := diagram.AddBlock("Start").SetShape(block.BlockShapeStadium)
	process := diagram.AddBlock("Process").SetShape(block.BlockShapeRoundEdges).SetWidth(2)
	transform := diagram.AddBlock("Transform").SetArrow(block.BlockArrowDirectionRight)
	diagram.AddBlock("End").SetShape(block.BlockShapeCircle)

	// Add links
	diagram.AddLink(start, process)
	diagram.AddLink(process, transform)

	// Write the diagram to README.md
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	readmePath := filepath.Join(dir, "README.md")
	if err := diagram.RenderToFile(readmePath); err != nil {
		fmt.Printf("Error writing diagram to README.md: %v\n", err)
		return
	}
}
