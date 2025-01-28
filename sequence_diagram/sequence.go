package sequence_diagram

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Diagram struct {
	Title         string
	Actors        []*Actor
	Messages      []*Message
	autonumber    bool
	markdownFence bool
}

// Creates a new sequence diagram with default settings
func NewDiagram() *Diagram {
	return &Diagram{
		Actors:     make([]*Actor, 0),
		Messages:   make([]*Message, 0),
		autonumber: false,
	}
}

// EnableMarkdownFence enables markdown fencing in the output
func (d *Diagram) EnableMarkdownFence() {
	d.markdownFence = true
}

// DisableMarkdownFence disables markdown fencing in the output
func (d *Diagram) DisableMarkdownFence() {
	d.markdownFence = false
}

func (d *Diagram) EnableAutoNumber() {
	d.autonumber = true
}

func (d *Diagram) AddActor(id, name string, actorType ActorType) *Actor {
	actor := NewActor(id, name, actorType)
	d.Actors = append(d.Actors, actor)
	return actor
}

// Creates a new actor and adds a creation message
func (d *Diagram) CreateActor(creator *Actor, id, name string, actorType ActorType) *Actor {
	newActor := NewActor(id, name, actorType)
	d.Actors = append(d.Actors, newActor)

	// Add creation message
	d.Messages = append(d.Messages, &Message{
		From: creator,
		To:   newActor,
		Type: messageCreate,
	})

	return newActor
}

// Adds a destroy message for an actor
func (d *Diagram) DestroyActor(actor *Actor) {
	d.Messages = append(d.Messages, &Message{
		From: nil, // Destruction doesn't need a from actor
		To:   actor,
		Type: messageDestroy,
	})
}

func (d *Diagram) AddMessage(from, to *Actor, msgType MessageType, text string) *Message {
	msg := NewMessage(from, to, msgType, text)
	d.Messages = append(d.Messages, msg)
	return msg
}

func (d *Diagram) String() string {
	var sb strings.Builder

	// Add markdown fence if enabled
	if d.markdownFence {
		sb.WriteString("```mermaid\n")
	}

	// Add title if present
	if d.Title != "" {
		sb.WriteString(fmt.Sprintf("---\ntitle: %s\n---\n\n", d.Title))
	}

	// Start sequence diagram
	sb.WriteString("sequenceDiagram\n")

	// Add autonumber if enabled
	if d.autonumber {
		sb.WriteString("    autonumber\n")
	}

	// First declare all actors that will appear in the diagram
	// This includes both initial actors and those created during the sequence
	for _, actor := range d.Actors {
		sb.WriteString(fmt.Sprintf("    %s %s as %s\n", actor.Type, actor.ID, actor.Name))
	}

	// Add messages
	for _, msg := range d.Messages {
		d.renderMessage(&sb, msg, "    ")
	}

	// Close markdown fence if enabled
	if d.markdownFence {
		sb.WriteString("```\n")
	}

	return sb.String()
}

// RenderToFile saves the diagram to a file at the specified path
// If the file extension is .md, markdown fencing is automatically enabled
func (d *Diagram) RenderToFile(path string) error {
	// Create directory if it doesn't exist
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// If file has .md extension, enable markdown fencing
	originalFenceState := d.markdownFence
	if strings.ToLower(filepath.Ext(path)) == ".md" {
		d.EnableMarkdownFence()
	}

	// Generate diagram content
	content := d.String()

	// Restore original fence state
	d.markdownFence = originalFenceState

	// Write to file
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func (d *Diagram) AddNote(position NotePosition, text string, actors ...*Actor) *Note {
	if len(actors) == 0 {
		return nil
	}
	note := NewNote(position, text, actors...)
	// Convert note to a special message type to maintain ordering
	msg := &Message{
		From: actors[0],             // Use first actor as reference
		To:   actors[0],             // Note doesn't really send a message
		Type: MessageType(position), // Store note position as message type
		Text: text,
		Note: note, // Store note information
	}
	d.Messages = append(d.Messages, msg)
	return note
}

func (d *Diagram) renderNote(sb *strings.Builder, note *Note, indent string) {
	switch len(note.Actors) {
	case 0:
		// Invalid note, skip it
		return
	case 1:
		// Single actor note
		sb.WriteString(fmt.Sprintf("%sNote %s %s: %s\n",
			indent, note.Position, note.Actors[0].ID, note.Text))
	default:
		// Multi-actor note (over)
		if note.Position != NoteOver {
			// Only "over" is valid for multiple actors
			return
		}
		actorIDs := make([]string, len(note.Actors))
		for i, actor := range note.Actors {
			actorIDs[i] = actor.ID
		}
		sb.WriteString(fmt.Sprintf("%sNote over %s: %s\n",
			indent, strings.Join(actorIDs, ","), note.Text))
	}
}

func (d *Diagram) renderMessage(sb *strings.Builder, msg *Message, indent string) {
	// If this message is actually a note, render it as such
	if msg.Note != nil {
		d.renderNote(sb, msg.Note, indent)
		return
	}

	switch msg.Type {
	case messageCreate:
		// Only render the creation message if there is text
		if msg.Text != "" {
			sb.WriteString(fmt.Sprintf("%s%s%s%s: %s\n",
				indent, msg.From.ID, MessageSolid, msg.To.ID, msg.Text))
		}

	case messageDestroy:
		// Render destruction message
		sb.WriteString(fmt.Sprintf("%sdestroy %s\n",
			indent, msg.To.ID))

	default:
		// Render regular message
		sb.WriteString(fmt.Sprintf("%s%s%s%s: %s\n",
			indent, msg.From.ID, msg.Type, msg.To.ID, msg.Text))
	}

	// Render nested messages
	for _, nested := range msg.Nested {
		d.renderMessage(sb, nested, indent+"    ")
	}
}
