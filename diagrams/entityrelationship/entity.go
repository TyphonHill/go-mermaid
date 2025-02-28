package entityrelationship

import (
	"fmt"
	"strings"

	"github.com/TyphonHill/go-mermaid/diagrams/utils/basediagram"
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

const (
	baseEntityNoAliasString   = basediagram.Indentation + "%s {\n"
	baseEntityWithAliasString = basediagram.Indentation + "%s [%s] {\n"
	baseEntityAttributeString = basediagram.Indentation + basediagram.Indentation + "%s %s%s\n"
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
	PK       bool
	FK       bool
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

	if e.Alias != "" {
		sb.WriteString(fmt.Sprintf(string(baseEntityWithAliasString), e.Name, e.Alias))
	} else {
		sb.WriteString(fmt.Sprintf(string(baseEntityNoAliasString), e.Name))
	}

	for _, attr := range e.Attributes {
		keys := ""
		if attr.PK && attr.FK {
			keys = " PK,FK"
		} else if attr.PK {
			keys = " PK"
		} else if attr.FK {
			keys = " FK"
		}
		sb.WriteString(fmt.Sprintf(string(baseEntityAttributeString), attr.Type, attr.Name, keys))
	}

	sb.WriteString(basediagram.Indentation + "}\n")
	return sb.String()
}
