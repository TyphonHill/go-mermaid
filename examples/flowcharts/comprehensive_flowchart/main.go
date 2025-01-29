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
	diagram.Title = "Order Processing System"

	// Create all nodes first
	start := diagram.AddNode("Start")
	cart := diagram.AddNode("Shopping Cart")
	checkout := diagram.AddNode("Checkout")
	validatePayment := diagram.AddNode("Validate Payment")
	processPayment := diagram.AddNode("Process Payment")
	paymentFailed := diagram.AddNode("Payment Failed")
	checkInventory := diagram.AddNode("Check Inventory")
	createOrder := diagram.AddNode("Create Order")
	notifyWarehouse := diagram.AddNode("Notify Warehouse")
	outOfStock := diagram.AddNode("Out of Stock")
	end := diagram.AddNode("End")

	// Create subgraphs and add relevant links
	userFlow := diagram.AddSubgraph("User Flow")
	userFlow.AddLink(start, cart)
	userFlow.AddLink(cart, checkout)

	paymentFlow := diagram.AddSubgraph("Payment Processing")
	paymentFlow.AddLink(checkout, validatePayment)
	paymentFlow.AddLink(validatePayment, processPayment).Text = "Valid"
	paymentFlow.AddLink(validatePayment, paymentFailed).Text = "Invalid"
	paymentFlow.AddLink(paymentFailed, checkout).Text = "Retry"

	fulfillmentFlow := diagram.AddSubgraph("Order Fulfillment")
	fulfillmentFlow.AddLink(processPayment, checkInventory)
	fulfillmentFlow.AddLink(checkInventory, createOrder).Text = "In Stock"
	fulfillmentFlow.AddLink(checkInventory, outOfStock).Text = "No Stock"
	fulfillmentFlow.AddLink(createOrder, notifyWarehouse)
	fulfillmentFlow.AddLink(notifyWarehouse, end)
	fulfillmentFlow.AddLink(outOfStock, cart).Text = "Update Cart"

	// Write the diagram to README.md in the same directory as this source file
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	readmePath := filepath.Join(dir, "README.md")
	if err := diagram.RenderToFile(readmePath); err != nil {
		fmt.Printf("Error writing diagram to README.md: %v\n", err)
		return
	}
}
