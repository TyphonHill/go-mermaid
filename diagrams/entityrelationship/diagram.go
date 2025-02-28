// Package entityrelationship provides functionality for creating Mermaid ER diagrams
package entityrelationship

import (
	"strings"

	"github.com/TyphonHill/go-mermaid/diagrams/utils"
	"github.com/TyphonHill/go-mermaid/diagrams/utils/basediagram"
)

// Base string formats for entity relationship diagrams
const (
	baseDiagramType string = "erDiagram\n"
)

// Diagram represents an entity relationship diagram
type Diagram struct {
	basediagram.BaseDiagram[ErConfigurationProperties]
	Entities      []*Entity
	Relationships []*Relationship
}

// NewDiagram creates a new ERD diagram
func NewDiagram() *Diagram {
	return &Diagram{
		BaseDiagram:   basediagram.NewBaseDiagram(NewErConfigurationProperties()),
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

	sb.WriteString(baseDiagramType)

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

	return d.BaseDiagram.String(sb.String())
}

// RenderToFile saves the diagram to a file at the specified path.
func (d *Diagram) RenderToFile(path string) error {
	return utils.RenderToFile(path, d.String())
}
