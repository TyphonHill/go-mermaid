package main

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/TyphonHill/go-mermaid/diagrams/block"
)

func main() {
	diagram := block.NewDiagram()
	diagram.SetColumns(4)

	// Different block styles and widths
	header := diagram.AddBlock("Dashboard").
		SetWidth(4).
		SetShape(block.BlockShapeRoundEdges).
		SetStyle("fill:#6BE,stroke:#333,stroke-width:2px")

	// Left side - Data Sources
	sources := diagram.AddBlock("")
	sources.AddBlock("Database").SetShape(block.BlockShapeCylindrical)
	sources.AddBlock("API").SetShape(block.BlockShapeHexagon).SetWidth(2)
	sources.AddBlock("Files").SetShape(block.BlockShapeParallelogram)
	sources.SetStyle("fill:#EBF,stroke:#333")

	// Right side - Processing
	process := diagram.AddBlock("")
	process.AddBlock("ETL Pipeline").SetShape(block.BlockShapeSubroutine).SetWidth(2)
	process.AddBlock("Transform").
		SetArrow(block.BlockArrowDirectionRight).
		SetWidth(2)
	process.AddBlock("ML Model").SetShape(block.BlockShapeRhombus)
	process.AddBlock("Aggregate").
		SetArrow(block.BlockArrowDirectionUp, block.BlockArrowDirectionDown)
	process.AddBlock("Cache").SetShape(block.BlockShapeDoubleCircle)
	process.SetStyle("fill:#BFE,stroke:#333")

	diagram.AddSpaceWithWidth(2)

	// Output blocks with different shapes and arrows
	viz := diagram.AddBlock("Visualizations").
		SetWidth(2).
		SetShape(block.BlockShapeStadium).
		SetStyle("fill:#FBE,stroke:#333")
	alerts := diagram.AddBlock("Alerts").
		SetArrow(block.BlockArrowDirectionX, block.BlockArrowDirectionY).
		SetStyle("fill:#FBE,stroke:#333")
	export := diagram.AddBlock("Export").
		SetArrow(block.BlockArrowDirectionRight).
		SetStyle("fill:#FBE,stroke:#333")

	// Add links between blocks
	diagram.AddLink(header, sources)
	diagram.AddLink(header, process)

	diagram.AddLink(sources.Children[0], process.Children[0])
	diagram.AddLink(sources.Children[1], process.Children[0])
	diagram.AddLink(sources.Children[2], process.Children[0])

	diagram.AddLink(process.Children[0], process.Children[1]).SetText("Data")
	diagram.AddLink(process.Children[1], process.Children[2]).SetText("Results")

	diagram.AddLink(process.Children[2], viz)
	diagram.AddLink(process.Children[2], alerts)
	diagram.AddLink(process.Children[2], export)

	// Write the diagram to README.md
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	readmePath := filepath.Join(dir, "README.md")
	if err := diagram.RenderToFile(readmePath); err != nil {
		fmt.Printf("Error writing diagram to README.md: %v\n", err)
		return
	}
}
