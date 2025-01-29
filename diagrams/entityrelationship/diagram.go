package entityrelationship

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Diagram represents an entity relationship diagram
type Diagram struct {
	Title         string
	Entities      []*Entity
	Relationships []*Relationship
	markdownFence bool
}

// NewDiagram creates a new ERD diagram
func NewDiagram() *Diagram {
	return &Diagram{
		Entities:      make([]*Entity, 0),
		Relationships: make([]*Relationship, 0),
	}
}

// EnableMarkdownFence enables markdown code fencing and returns the diagram for chaining
func (d *Diagram) EnableMarkdownFence() *Diagram {
	d.markdownFence = true
	return d
}

// DisableMarkdownFence disables markdown code fencing
func (d *Diagram) DisableMarkdownFence() {
	d.markdownFence = false
}

// SetTitle sets the diagram title and returns the diagram for chaining
func (d *Diagram) SetTitle(title string) *Diagram {
	d.Title = title
	return d
}

// AddEntity creates and adds a new entity to the diagram
func (d *Diagram) AddEntity(name string) *Entity {
	entity := NewEntity(name)
	d.Entities = append(d.Entities, entity)
	return entity
}

// AddRelationship creates a new relationship between two entities
func (d *Diagram) AddRelationship(from, to *Entity) *Relationship {
	rel := NewRelationship(from, to)
	d.Relationships = append(d.Relationships, rel)
	return rel
}

// String generates the Mermaid syntax for the diagram
func (d *Diagram) String() string {
	var sb strings.Builder

	if d.markdownFence {
		sb.WriteString("```mermaid\n")
	}

	if d.Title != "" {
		sb.WriteString(fmt.Sprintf("---\ntitle: %s\n---\n\n", d.Title))
	}

	sb.WriteString("erDiagram\n")

	// Add entities
	for _, entity := range d.Entities {
		sb.WriteString(entity.String())
	}

	// Add relationships with a newline between entities and relationships
	if len(d.Relationships) > 0 {
		sb.WriteString("\n")
		for _, rel := range d.Relationships {
			sb.WriteString(rel.String())
		}
	}

	if d.markdownFence {
		sb.WriteString("```\n")
	}

	return sb.String()
}

// RenderToFile saves the diagram to a file
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
