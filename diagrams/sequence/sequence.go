package sequence

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Diagram represents a sequence diagram with actors, messages, and rendering options.
type Diagram struct {
	Title         string
	Actors        []*Actor
	Messages      []*Message
	autonumber    bool
	markdownFence bool
}

// NewDiagram creates a new sequence diagram with default settings.
func NewDiagram() *Diagram {
	return &Diagram{
		Actors:     make([]*Actor, 0),
		Messages:   make([]*Message, 0),
		autonumber: false,
	}
}

// EnableMarkdownFence enables markdown fencing for the diagram output.
func (d *Diagram) EnableMarkdownFence() {
	d.markdownFence = true
}

// DisableMarkdownFence disables markdown fencing for the diagram output.
func (d *Diagram) DisableMarkdownFence() {
	d.markdownFence = false
}

// EnableAutoNumber enables automatic numbering of messages in the sequence diagram.
func (d *Diagram) EnableAutoNumber() {
	d.autonumber = true
}

// AddActor creates and adds a new actor to the diagram.
func (d *Diagram) AddActor(id, name string, actorType ActorType) *Actor {
	actor := NewActor(id, name, actorType)
	d.Actors = append(d.Actors, actor)
	return actor
}

// CreateActor adds a new actor to the diagram with a creation message.
func (d *Diagram) CreateActor(creator *Actor, id, name string, actorType ActorType) *Actor {
	newActor := NewActor(id, name, actorType)
	d.Actors = append(d.Actors, newActor)

	d.Messages = append(d.Messages, &Message{
		From: creator,
		To:   newActor,
		Type: MessageCreate,
	})

	return newActor
}

// DestroyActor adds a destroy message for the specified actor.
func (d *Diagram) DestroyActor(actor *Actor) {
	d.Messages = append(d.Messages, &Message{
		From: nil,
		To:   actor,
		Type: MessageDestroy,
	})
}

// AddMessage creates and adds a new message to the diagram.
func (d *Diagram) AddMessage(from, to *Actor, msgType MessageType, text string) *Message {
	msg := NewMessage(from, to, msgType, text)
	d.Messages = append(d.Messages, msg)
	return msg
}

// String generates a Mermaid-formatted string representation of the sequence diagram.
func (d *Diagram) String() string {
	var sb strings.Builder

	if d.markdownFence {
		sb.WriteString("```mermaid\n")
	}

	if d.Title != "" {
		sb.WriteString(fmt.Sprintf("---\ntitle: %s\n---\n\n", d.Title))
	}

	sb.WriteString("sequenceDiagram\n")

	if d.autonumber {
		sb.WriteString("autonumber\n")
	}

	// Render actors
	for _, actor := range d.Actors {
		sb.WriteString(fmt.Sprintf("\t%s %s as %s\n",
			actor.Type, actor.ID, actor.Name))
	}

	// Render messages (which will include notes)
	for _, message := range d.Messages {
		sb.WriteString(message.String(""))
	}

	if d.markdownFence {
		sb.WriteString("```\n")
	}

	return sb.String()
}

// RenderToFile saves the diagram to a file, automatically enabling markdown fencing for .md files.
func (d *Diagram) RenderToFile(path string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	originalFenceState := d.markdownFence
	if strings.ToLower(filepath.Ext(path)) == ".md" {
		d.EnableMarkdownFence()
	}

	content := d.String()
	d.markdownFence = originalFenceState

	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// AddNote adds a note to the diagram positioned relative to one or more actors.
func (d *Diagram) AddNote(position NotePosition, text string, actors ...*Actor) *Note {
	if len(actors) == 0 {
		return nil
	}
	note := NewNote(position, text, actors...)
	msg := &Message{
		From: actors[0],
		To:   actors[0],
		Type: MessageType(position),
		Text: text,
		Note: note,
	}
	d.Messages = append(d.Messages, msg)
	return note
}
