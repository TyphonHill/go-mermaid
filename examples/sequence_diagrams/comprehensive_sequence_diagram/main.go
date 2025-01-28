package main

import (
	"fmt"

	sequence "github.com/TyphonHill/go-mermaid/sequence_diagram"
)

func main() {
	// Create a new diagram with all the bells and whistles
	diagram := sequence.NewDiagram()

	// Enable markdown fencing for compatibility with markdown documents
	diagram.EnableMarkdownFence()

	// Set title and enable autonumbering
	diagram.Title = "E-commerce Order Processing Flow"
	diagram.EnableAutoNumber()

	// Create initial actors
	browser := diagram.AddActor("browser", "Web Browser", sequence.ActorParticipant)
	frontend := diagram.AddActor("frontend", "Frontend Server", sequence.ActorParticipant)
	orderSvc := diagram.AddActor("orders", "Order Service", sequence.ActorActor)

	// Add a note explaining the flow
	diagram.AddNote(sequence.NoteOver, "Customer places a new order", browser, frontend)

	// Initial order submission
	orderReq := diagram.AddMessage(browser, frontend, sequence.MessageSolid, "Submit Order")

	// Dynamically create payment service actor
	paymentSvc := diagram.CreateActor(frontend, "payment", "Payment Service", sequence.ActorParticipant)
	diagram.AddNote(sequence.NoteRight, "Payment service initialized on demand", paymentSvc)

	// Payment processing flow with nested messages
	paymentFlow := orderReq.AddNestedMessage(frontend, paymentSvc, sequence.MessageSolid, "Process Payment")
	paymentFlow.AddNestedMessage(paymentSvc, frontend, sequence.MessageAsync, "Payment Processing Started")

	// Add inventory service
	inventory := diagram.AddActor("inventory", "Inventory Service", sequence.ActorParticipant)

	// Check inventory in parallel with payment
	inventoryCheck := orderReq.AddNestedMessage(frontend, inventory, sequence.MessageSolid, "Check Stock")
	diagram.AddNote(sequence.NoteLeft, "Verify item availability", inventory)
	inventoryCheck.AddNestedMessage(inventory, frontend, sequence.MessageResponse, "Items Available")

	// Order service processing with activation
	processOrder := diagram.AddMessage(frontend, orderSvc, sequence.MessageActivate, "Create Order")
	processOrder.AddNestedMessage(orderSvc, inventory, sequence.MessageSolid, "Reserve Items")
	processOrder.AddNestedMessage(orderSvc, paymentSvc, sequence.MessageSolid, "Confirm Payment")

	// Complete order processing
	diagram.AddMessage(orderSvc, frontend, sequence.MessageDeactivate, "Order Created")

	// Success response to browser
	diagram.AddMessage(frontend, browser, sequence.MessageResponse, "Order Confirmation")

	// Cleanup: destroy payment service actor
	diagram.DestroyActor(paymentSvc)

	// Final note
	diagram.AddNote(sequence.NoteOver, "Order processing complete", browser, orderSvc)

	// Print the diagram with markdown fence markers
	fmt.Println(diagram.String())
}
