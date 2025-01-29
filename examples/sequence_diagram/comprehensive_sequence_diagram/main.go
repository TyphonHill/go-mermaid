package main

import (
	"fmt"
	"path/filepath"
	"runtime"

	sequenceDiagram "github.com/TyphonHill/go-mermaid/diagrams/sequence"
)

func main() {
	// Create a new diagram with all the bells and whistles
	diagram := sequenceDiagram.NewDiagram()

	// Enable markdown fencing for compatibility with markdown documents
	diagram.EnableMarkdownFence()

	// Set title and enable autonumbering
	diagram.Title = "E-commerce Order Processing Flow"
	diagram.EnableAutoNumber()

	// Create initial actors
	browser := diagram.AddActor("browser", "Web Browser", sequenceDiagram.ActorParticipant)
	frontend := diagram.AddActor("frontend", "Frontend Server", sequenceDiagram.ActorParticipant)
	orderSvc := diagram.AddActor("orders", "Order Service", sequenceDiagram.ActorActor)

	// Add a note explaining the flow
	diagram.AddNote(sequenceDiagram.NoteOver, "Customer places a new order", browser, frontend)

	// Initial order submission
	orderReq := diagram.AddMessage(browser, frontend, sequenceDiagram.MessageSolid, "Submit Order")

	// Dynamically create payment service actor
	paymentSvc := diagram.CreateActor(frontend, "payment", "Payment Service", sequenceDiagram.ActorParticipant)
	diagram.AddNote(sequenceDiagram.NoteRight, "Payment service initialized on demand", paymentSvc)

	// Payment processing flow with nested messages
	paymentFlow := orderReq.AddNestedMessage(frontend, paymentSvc, sequenceDiagram.MessageSolid, "Process Payment")
	paymentFlow.AddNestedMessage(paymentSvc, frontend, sequenceDiagram.MessageAsync, "Payment Processing Started")

	// Add inventory service
	inventory := diagram.AddActor("inventory", "Inventory Service", sequenceDiagram.ActorParticipant)

	// Check inventory in parallel with payment
	inventoryCheck := orderReq.AddNestedMessage(frontend, inventory, sequenceDiagram.MessageSolid, "Check Stock")
	diagram.AddNote(sequenceDiagram.NoteLeft, "Verify item availability", inventory)
	inventoryCheck.AddNestedMessage(inventory, frontend, sequenceDiagram.MessageResponse, "Items Available")

	// Order service processing with activation
	processOrder := diagram.AddMessage(frontend, orderSvc, sequenceDiagram.MessageActivate, "Create Order")
	processOrder.AddNestedMessage(orderSvc, inventory, sequenceDiagram.MessageSolid, "Reserve Items")
	processOrder.AddNestedMessage(orderSvc, paymentSvc, sequenceDiagram.MessageSolid, "Confirm Payment")

	// Complete order processing
	diagram.AddMessage(orderSvc, frontend, sequenceDiagram.MessageResponse, "Order Created")
	diagram.AddMessage(orderSvc, orderSvc, sequenceDiagram.MessageDeactivate, "") // Deactivate orderSvc

	// Success response to browser
	diagram.AddMessage(frontend, browser, sequenceDiagram.MessageResponse, "Order Confirmation")

	// Cleanup: destroy payment service actor
	diagram.DestroyActor(paymentSvc)

	// Final note
	diagram.AddNote(sequenceDiagram.NoteOver, "Order processing complete", browser, orderSvc)

	// Write the diagram to README.md in the same directory as this source file
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	readmePath := filepath.Join(dir, "README.md")
	if err := diagram.RenderToFile(readmePath); err != nil {
		fmt.Printf("Error writing diagram to README.md: %v\n", err)
		return
	}
}
