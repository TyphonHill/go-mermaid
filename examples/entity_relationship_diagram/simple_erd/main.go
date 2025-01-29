package main

import (
	"fmt"
	"path/filepath"
	"runtime"

	erd "github.com/TyphonHill/go-mermaid/diagrams/entityrelationship"
)

func main() {
	diagram := erd.NewDiagram()
	diagram.EnableMarkdownFence()
	diagram.SetTitle("Blog System ERD")

	// Create entities
	user := diagram.AddEntity("USER_TABLE").
		SetAlias("User")
	user.AddAttribute("id", erd.TypeInteger).SetPrimaryKey()
	user.AddAttribute("username", erd.TypeString).SetRequired()
	user.AddAttribute("email", erd.TypeString).SetRequired()
	user.AddAttribute("created_at", erd.TypeDateTime)

	post := diagram.AddEntity("POSTS").
		SetAlias("Post")
	post.AddAttribute("id", erd.TypeInteger).SetPrimaryKey()
	post.AddAttribute("title", erd.TypeString).SetRequired()
	post.AddAttribute("content", erd.TypeString).SetRequired()
	post.AddAttribute("published", erd.TypeBoolean)
	post.AddAttribute("created_at", erd.TypeDateTime)

	comment := diagram.AddEntity("COMMENTS").
		SetAlias("Comment")
	comment.AddAttribute("id", erd.TypeInteger).SetPrimaryKey()
	comment.AddAttribute("user_id", erd.TypeInteger).SetForeignKey()
	comment.AddAttribute("post_id", erd.TypeInteger).SetForeignKey()
	comment.AddAttribute("content", erd.TypeString).SetRequired()
	comment.AddAttribute("created_at", erd.TypeDateTime)

	// Create relationships
	diagram.AddRelationship(user, post).
		SetLabel("writes").
		SetCardinality(erd.OneToZeroOrMore)

	diagram.AddRelationship(post, comment).
		SetLabel("has").
		SetCardinality(erd.OneToZeroOrMore)

	diagram.AddRelationship(user, comment).
		SetLabel("writes").
		SetCardinality(erd.OneToZeroOrMore)

	// Write the diagram to README.md in the same directory as this source file
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	readmePath := filepath.Join(dir, "README.md")
	if err := diagram.RenderToFile(readmePath); err != nil {
		fmt.Printf("Error writing diagram to README.md: %v\n", err)
		return
	}
}
