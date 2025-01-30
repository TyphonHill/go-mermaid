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
	diagram.SetTitle("E-Commerce User Journey")

	// Discovery section - how users find and explore the platform
	discovery := diagram.AddSection("Discovery")
	discovery.AddTask("See Social Media Ad", 3, "Customer", "Marketing")
	discovery.AddTask("Visit Website", 5, "Customer")
	discovery.AddTask("Browse Categories", 4, "Customer")
	discovery.AddTask("Read Product Reviews", 4, "Customer", "Previous Buyers")

	// Account section - user registration and profile setup
	account := diagram.AddSection("Account")
	account.AddTask("Create Account", 2, "Customer", "System")
	account.AddTask("Verify Email", 3, "Customer", "System")
	account.AddTask("Complete Profile", 3, "Customer")
	account.AddTask("Save Payment Info", 4, "Customer", "Payment Provider")

	// Shopping section - product selection and cart management
	shopping := diagram.AddSection("Shopping")
	shopping.AddTask("Search Products", 4, "Customer")
	shopping.AddTask("Apply Filters", 3, "Customer")
	shopping.AddTask("Compare Items", 4, "Customer")
	shopping.AddTask("Add to Cart", 5, "Customer", "System")
	shopping.AddTask("Adjust Quantities", 4, "Customer", "System")

	// Checkout section - payment and order confirmation
	checkout := diagram.AddSection("Checkout")
	checkout.AddTask("Review Cart", 5, "Customer")
	checkout.AddTask("Apply Coupon", 4, "Customer", "System")
	checkout.AddTask("Select Shipping", 3, "Customer", "Shipping Provider")
	checkout.AddTask("Complete Payment", 4, "Customer", "Payment Provider")
	checkout.AddTask("Receive Confirmation", 5, "Customer", "System")

	// Post-Purchase section - order tracking and follow-up
	postPurchase := diagram.AddSection("Post-Purchase")
	postPurchase.AddTask("Track Order", 5, "Customer", "Shipping Provider")
	postPurchase.AddTask("Receive Updates", 4, "Customer", "System", "Shipping Provider")
	postPurchase.AddTask("Get Delivery", 5, "Customer", "Delivery Driver")
	postPurchase.AddTask("Rate Purchase", 3, "Customer")
	postPurchase.AddTask("Contact Support", 2, "Customer", "Support Team")

	// Support section - customer service interactions
	support := diagram.AddSection("Support")
	support.AddTask("Find FAQ", 3, "Customer")
	support.AddTask("Chat with Bot", 2, "Customer", "Chat Bot")
	support.AddTask("Email Support", 3, "Customer", "Support Team")
	support.AddTask("Live Chat", 4, "Customer", "Support Agent")
	support.AddTask("Resolution", 4, "Customer", "Support Agent")

	// Write the diagram to README.md in the same directory as this source file
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	readmePath := filepath.Join(dir, "README.md")
	if err := diagram.RenderToFile(readmePath); err != nil {
		fmt.Printf("Error writing diagram to README.md: %v\n", err)
		return
	}
}
