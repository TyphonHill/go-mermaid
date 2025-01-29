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
	diagram.Title = "Order Processing State Machine"

	// Create main states
	idle := diagram.AddState("idle", "Order Idle", stateDiagram.StateNormal)
	validating := diagram.AddState("validating", "Validating Order", stateDiagram.StateNormal)
	processing := diagram.AddState("processing", "Processing Order", stateDiagram.StateComposite)
	error := diagram.AddState("error", "Error State", stateDiagram.StateNormal)
	decision := diagram.AddState("decision", "Payment Method", stateDiagram.StateChoice)

	// Add nested states to processing
	paymentPending := processing.AddNestedState("payment_pending", "Payment Pending", stateDiagram.StateNormal)
	fulfillment := processing.AddNestedState("fulfillment", "Order Fulfillment", stateDiagram.StateNormal)

	// Add fork and join states for parallel processing
	fork := diagram.AddState("fork", "Fork", stateDiagram.StateFork)
	join := diagram.AddState("join", "Join", stateDiagram.StateJoin)

	// Add notes to states
	idle.AddNote("System waiting for new orders", stateDiagram.NoteLeft)
	processing.AddNote("Payment and fulfillment processing", stateDiagram.NoteRight)
	error.AddNote("Error handling and recovery", stateDiagram.NoteRight)

	// Add transitions from start state
	diagram.AddTransition(nil, idle, "New Order")

	// Add main flow transitions
	diagram.AddTransition(idle, validating, "Submit Order")
	diagram.AddTransition(validating, decision, "Validation Complete")

	// Add choice transitions
	diagram.AddTransition(decision, processing, "Credit Card")
	diagram.AddTransition(decision, error, "Invalid Payment")

	// Add processing flow with nested states
	diagram.AddTransition(processing, paymentPending, "Process Payment")
	diagram.AddTransition(paymentPending, fork, "Payment Confirmed")

	// Add parallel processing paths
	diagram.AddTransition(fork, fulfillment, "Start Fulfillment")
	diagram.AddTransition(fork, validating, "Revalidate Stock")

	// Join parallel paths
	diagram.AddTransition(fulfillment, join, "Fulfillment Complete")
	diagram.AddTransition(validating, join, "Stock Confirmed")

	// Add completion and error transitions
	diagram.AddTransition(join, idle, "Order Complete")
	diagram.AddTransition(error, idle, "Reset Order")

	// Add transition to end state
	diagram.AddTransition(error, nil, "Cancel Order")

	// Write the diagram to README.md in the same directory as this source file
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	readmePath := filepath.Join(dir, "README.md")
	if err := diagram.RenderToFile(readmePath); err != nil {
		fmt.Printf("Error writing diagram to README.md: %v\n", err)
		return
	}
}
