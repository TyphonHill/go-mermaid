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
	diagram.Title = "Simple Class Diagram"

	// Create classes
	user := diagram.AddClass("User", nil)
	user.AddField("id", "string")
	user.AddField("name", "string")
	user.AddField("email", "string")

	loginMethod := user.AddMethod("login")
	loginMethod.AddParameter("password", "string")
	loginMethod.ReturnType = "bool"

	user.AddMethod("logout")

	account := diagram.AddClass("Account", nil)
	account.AddField("id", "string")
	account.AddField("balance", "float")

	depositMethod := account.AddMethod("deposit")
	depositMethod.AddParameter("amount", "float")

	withdrawMethod := account.AddMethod("withdraw")
	withdrawMethod.AddParameter("amount", "float")
	withdrawMethod.ReturnType = "error"

	// Add relationship
	rel := diagram.AddRelation(user, account)
	rel.Label = "has"

	// Write the diagram to README.md in the same directory as this source file
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	readmePath := filepath.Join(dir, "README.md")
	if err := diagram.RenderToFile(readmePath); err != nil {
		fmt.Printf("Error writing diagram to README.md: %v\n", err)
		return
	}
}
