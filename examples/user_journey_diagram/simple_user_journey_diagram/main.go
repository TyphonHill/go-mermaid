package main

import (
	"fmt"
	"path/filepath"
	"runtime"

	journey "github.com/TyphonHill/go-mermaid/diagrams/userjourney"
)

func main() {
	diagram := journey.NewDiagram()
	diagram.EnableMarkdownFence()
	diagram.SetTitle("Online Shopping User Journey")

	// Browse section
	browse := diagram.AddSection("Browse")
	browse.AddTask("Visit Homepage", 5)
	browse.AddTask("View Categories", 3)
	browse.AddTask("Search Products", 2)
	browse.AddTask("Read Reviews", 4)

	// Purchase section
	purchase := diagram.AddSection("Purchase")
	purchase.AddTask("Add to Cart", 3)
	purchase.AddTask("Enter Address", 2)
	purchase.AddTask("Choose Payment", 3)
	purchase.AddTask("Confirm Order", 5)

	// Post-Purchase section
	postPurchase := diagram.AddSection("Post-Purchase")
	postPurchase.AddTask("Receive Order Confirmation", 5)
	postPurchase.AddTask("Track Package", 4)
	postPurchase.AddTask("Receive Delivery", 5)
	postPurchase.AddTask("Write Review", 2)

	// Write the diagram to README.md in the same directory as this source file
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	readmePath := filepath.Join(dir, "README.md")
	if err := diagram.RenderToFile(readmePath); err != nil {
		fmt.Printf("Error writing diagram to README.md: %v\n", err)
		return
	}
}
