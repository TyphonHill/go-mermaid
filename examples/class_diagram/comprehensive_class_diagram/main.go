package main

import (
	"fmt"
	"path/filepath"
	"runtime"

	classDiagram "github.com/TyphonHill/go-mermaid/diagrams/class"
)

func main() {
	diagram := classDiagram.NewClassDiagram()
	diagram.EnableMarkdownFence()
	diagram.Title = "E-commerce System Class Diagram"

	// Create namespaces
	modelNS := diagram.AddNamespace("Models")
	serviceNS := diagram.AddNamespace("Services")

	// Create base interfaces
	entity := diagram.AddClass("Entity", modelNS)
	entity.Annotation = classDiagram.ClassAnnotationInterface
	entity.AddField("id", "string")
	validateMethod := entity.AddMethod("validate")
	validateMethod.ReturnType = "error"

	repository := diagram.AddClass("Repository", serviceNS)
	repository.Annotation = classDiagram.ClassAnnotationInterface

	findMethod := repository.AddMethod("find")
	findMethod.AddParameter("id", "string")
	findMethod.ReturnType = "Entity"

	saveMethod := repository.AddMethod("save")
	saveMethod.AddParameter("entity", "Entity")
	saveMethod.ReturnType = "error"

	// Create model classes
	user := diagram.AddClass("User", modelNS)
	user.AddField("name", "string")
	user.AddField("email", "string")
	user.AddField("passwordHash", "string")

	validatePasswordMethod := user.AddMethod("validatePassword")
	validatePasswordMethod.AddParameter("password", "string")
	validatePasswordMethod.ReturnType = "bool"

	order := diagram.AddClass("Order", modelNS)
	order.AddField("userId", "string")
	order.AddField("items", "List~OrderItem~")
	order.AddField("status", "OrderStatus")
	temp := order.AddMethod("calculateTotal")
	temp.ReturnType = "float"

	orderItem := diagram.AddClass("OrderItem", modelNS)
	orderItem.AddField("productId", "string")
	orderItem.AddField("quantity", "int")
	orderItem.AddField("price", "float")

	// Create service classes
	userService := diagram.AddClass("UserService", serviceNS)
	userService.AddField("repo", "UserRepository")

	registerMethod := userService.AddMethod("register")
	registerMethod.AddParameter("user", "User")
	registerMethod.ReturnType = "error"

	authMethod := userService.AddMethod("authenticate")
	authMethod.AddParameter("credentials", "LoginDTO")
	authMethod.ReturnType = "Token"

	orderService := diagram.AddClass("OrderService", serviceNS)
	orderService.AddField("repo", "OrderRepository")

	orderService.AddField("userService", "UserService")

	createMethod := orderService.AddMethod("createOrder")
	createMethod.AddParameter("order", "Order")
	createMethod.ReturnType = "error"

	processMethod := orderService.AddMethod("processOrder")
	processMethod.AddParameter("orderId", "string")
	processMethod.ReturnType = "error"

	// Add relationships
	diagram.AddRelation(user, entity)
	diagram.AddRelation(order, entity)
	diagram.AddRelation(orderItem, entity)
	diagram.AddRelation(order, orderItem)
	diagram.AddRelation(user, order)

	// Add notes
	diagram.AddNote("Base interface for all domain entities", entity)
	diagram.AddNote("Generic repository interface for data access", repository)

	// Write the diagram to README.md in the same directory as this source file
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	readmePath := filepath.Join(dir, "README.md")
	if err := diagram.RenderToFile(readmePath); err != nil {
		fmt.Printf("Error writing diagram to README.md: %v\n", err)
		return
	}
}
