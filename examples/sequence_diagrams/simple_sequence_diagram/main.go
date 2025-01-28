package main

import (
	"fmt"

	sequence "github.com/TyphonHill/go-mermaid/sequence_diagram"
)

func main() {
	// Create a new diagram
	diagram := sequence.NewDiagram()

	// Enable markdown fencing for compatibility with markdown documents
	diagram.EnableMarkdownFence()

	// Set title
	diagram.Title = "Login Flow"

	// Create actors
	client := diagram.AddActor("client", "Client", sequence.ActorParticipant)
	server := diagram.AddActor("server", "Auth Server", sequence.ActorParticipant)
	db := diagram.AddActor("db", "Database", sequence.ActorParticipant)

	// Add login request flow
	loginReq := diagram.AddMessage(client, server, sequence.MessageSolid, "POST /login")

	// Server validates credentials with database
	validateReq := loginReq.AddNestedMessage(server, db, sequence.MessageSolid, "Validate credentials")
	validateReq.AddNestedMessage(db, server, sequence.MessageResponse, "User validated")

	// Server sends back token
	diagram.AddMessage(server, client, sequence.MessageResponse, "Return JWT token")

	// Print the diagram
	fmt.Println(diagram.String())
}
