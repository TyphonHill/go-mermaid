package main

import (
	"fmt"
	"path/filepath"
	"runtime"

	sequenceDiagram "github.com/TyphonHill/go-mermaid/diagrams/sequence"
)

func main() {
	// Create a new diagram
	diagram := sequenceDiagram.NewDiagram()

	// Enable markdown fencing for compatibility with markdown documents
	diagram.EnableMarkdownFence()

	// Set title
	diagram.SetTitle("Login Flow")

	// Create actors
	client := diagram.AddActor("client", "Client", sequenceDiagram.ActorParticipant)
	server := diagram.AddActor("server", "Auth Server", sequenceDiagram.ActorParticipant)
	db := diagram.AddActor("db", "Database", sequenceDiagram.ActorParticipant)

	// Add login request flow
	loginReq := diagram.AddMessage(client, server, sequenceDiagram.MessageSolid, "POST /login")

	// Server validates credentials with database
	validateReq := loginReq.AddNestedMessage(server, db, sequenceDiagram.MessageSolid, "Validate credentials")
	validateReq.AddNestedMessage(db, server, sequenceDiagram.MessageResponse, "User validated")

	// Server sends back token
	diagram.AddMessage(server, client, sequenceDiagram.MessageResponse, "Return JWT token")

	// Write the diagram to README.md in the same directory as this source file
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	readmePath := filepath.Join(dir, "README.md")
	if err := diagram.RenderToFile(readmePath); err != nil {
		fmt.Printf("Error writing diagram to README.md: %v\n", err)
		return
	}
}
