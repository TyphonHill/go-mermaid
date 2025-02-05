// Package entityrelationship provides functionality for creating Mermaid ER diagrams
package entityrelationship

import (
	"fmt"
	"strings"

	"github.com/TyphonHill/go-mermaid/diagrams/utils"
)

// Diagram represents an entity relationship diagram
type Diagram struct {
	utils.BaseDiagram
	Entities      []*Entity
	Relationships []*Relationship
}

// NewDiagram creates a new ERD diagram
func NewDiagram() *Diagram {
	return &Diagram{
		Entities:      make([]*Entity, 0),
		Relationships: make([]*Relationship, 0),
	}
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

	if d.Title != "" {
		sb.WriteString(fmt.Sprintf("---\ntitle: %s\n---\n\n", d.Title))
	}

	sb.WriteString("erDiagram\n")

	// Add entities
	for _, entity := range d.Entities {
		sb.WriteString(entity.String())
	}

	// Add relationships
	if len(d.Relationships) > 0 {
		sb.WriteString("\n")
		for _, rel := range d.Relationships {
			sb.WriteString(rel.String())
		}
	}

	return d.WrapWithFence(sb.String())
}

// RenderToFile saves the diagram to a file at the specified path.
func (d *Diagram) RenderToFile(path string) error {
	return utils.RenderToFile(path, d.String())
}
