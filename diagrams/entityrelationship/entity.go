package entityrelationship

import (
	"fmt"
	"strings"
)

// DataType represents the type of an attribute
type DataType string

const (
	TypeString   DataType = "string"
	TypeInteger  DataType = "int"
	TypeFloat    DataType = "float"
	TypeBoolean  DataType = "boolean"
	TypeDateTime DataType = "datetime"
)

// Entity represents a table or entity in the ERD
type Entity struct {
	Name       string
	Alias      string
	Attributes []*Attribute
}

// Attribute represents a column or field in an entity
type Attribute struct {
	Name     string
	Type     DataType
	PK       bool // Primary Key
	FK       bool // Foreign Key
	Required bool
}

// NewEntity creates a new Entity
func NewEntity(name string) *Entity {
	return &Entity{
		Name:       name,
		Attributes: make([]*Attribute, 0),
	}
}

// AddAttribute adds a new attribute to the entity
func (e *Entity) AddAttribute(name string, dataType DataType) *Attribute {
	attr := &Attribute{
		Name: name,
		Type: dataType,
	}
	e.Attributes = append(e.Attributes, attr)
	return attr
}

// SetPrimaryKey marks the attribute as a primary key and returns it for chaining
func (a *Attribute) SetPrimaryKey() *Attribute {
	a.PK = true
	return a
}

// SetForeignKey marks the attribute as a foreign key and returns it for chaining
func (a *Attribute) SetForeignKey() *Attribute {
	a.FK = true
	return a
}

// SetRequired marks the attribute as required and returns it for chaining
func (a *Attribute) SetRequired() *Attribute {
	a.Required = true
	return a
}

// SetAlias sets an alternative display name for the entity and returns it for chaining
func (e *Entity) SetAlias(alias string) *Entity {
	e.Alias = alias
	return e
}

// String generates the Mermaid syntax for the entity
func (e *Entity) String() string {
	var sb strings.Builder

	// Start entity definition with optional alias
	if e.Alias != "" {
		sb.WriteString(fmt.Sprintf("\t%s [%s] {\n", e.Name, e.Alias))
	} else {
		sb.WriteString(fmt.Sprintf("\t%s {\n", e.Name))
	}

	// Add attributes
	for _, attr := range e.Attributes {
		keys := ""
		if attr.PK && attr.FK {
			keys = " PK,FK"
		} else if attr.PK {
			keys = " PK"
		} else if attr.FK {
			keys = " FK"
		}
		sb.WriteString(fmt.Sprintf("\t\t%s %s%s\n", attr.Type, attr.Name, keys))
	}

	sb.WriteString("\t}\n")
	return sb.String()
}
