package entityrelationship

import (
	"fmt"

	"github.com/TyphonHill/go-mermaid/diagrams/utils/basediagram"
)

// Cardinality represents the relationship cardinality
type Cardinality string

// Common relationship patterns
const (
	OneToZeroOrMore Cardinality = "||--o{" // One to many (optional)
	OneToOneOrMore  Cardinality = "||--|{" // One to many (required)
	OneToExactlyOne Cardinality = "||--||" // One to one
	ZeroOrOneToMany Cardinality = "|o--o{" // Optional one to many
	ManyToMany      Cardinality = "}o--o{" // Many to many

	// Base cardinality symbols
	ZeroOrOne  Cardinality = "|o"
	ExactlyOne Cardinality = "||"
	ZeroOrMore Cardinality = "o{"
	OneOrMore  Cardinality = "|{"
)

// Relationship represents a relationship between two entities
type Relationship struct {
	From        *Entity
	To          *Entity
	Label       string
	Cardinality Cardinality
}

// NewRelationship creates a new relationship between entities
func NewRelationship(from, to *Entity) *Relationship {
	return &Relationship{
		From:        from,
		To:          to,
		Cardinality: ExactlyOne, // default cardinality
	}
}

// SetLabel sets the relationship label
func (r *Relationship) SetLabel(label string) *Relationship {
	r.Label = label
	return r
}

// SetCardinality sets the relationship cardinality
func (r *Relationship) SetCardinality(cardinality Cardinality) *Relationship {
	r.Cardinality = cardinality
	return r
}

// String generates the Mermaid syntax for the relationship
func (r *Relationship) String() string {
	label := r.Label
	if label == "" {
		label = "relates"
	}
	return fmt.Sprintf(basediagram.Indentation+"%s %s %s : %s\n", r.From.Name, string(r.Cardinality), r.To.Name, label)
}
