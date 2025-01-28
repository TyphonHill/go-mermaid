package classdiagram

import (
	"fmt"
	"strings"
)

// relationType represents the type of relationship between classes.
type relationType string

// relationLink represents the visual style of the relationship line.
type relationLink string

// relationCardinality represents the multiplicity of the relationship.
type relationCardinality string

// Relationship type constants define different types of class relationships.
const (
	RelationTypeAssociation     relationType = ">"
	RelationTypeAssociationLeft relationType = "<"
	RelationTypeInheritance     relationType = "|>"
	RelationTypeInheritanceLeft relationType = "<|"
	RelationTypeComposition     relationType = "*"
	RelationTypeAggregation     relationType = "o"
)

// Relationship link constants define the line style for relationships.
const (
	RelationLinkSolid  relationLink = "--"
	RelationLinkDashed relationLink = ".."
)

// Relationship cardinality constants define the multiplicity of relationships.
const (
	RelationCardinalityOnlyOne   relationCardinality = "\"1\""
	RelationCardinalityZeroOrOne relationCardinality = "\"0..1\""
	RelationCardinalityOneOrMore relationCardinality = "\"1..*\""
	RelationCardinalityMany      relationCardinality = "\"*\""
	RelationCardinalityN         relationCardinality = "\"n\""
	RelationCardinalityZeroToN   relationCardinality = "\"0..n\""
	RelationCardinalityOneToN    relationCardinality = "\"1..n\""
)

// Formatting constants for relation string representation.
const (
	baseRelationString     string = "\t%s %s%s%s%s%s %s%s\n"
	baseRelationTextString string = " : %s"
)

// Relation represents a relationship between two classes in a class diagram.
// It includes information about the related classes, relationship types,
// cardinalities, link style, and optional label.
type Relation struct {
	ClassA              *Class
	ClassB              *Class
	RelationToClassA    relationType
	RelationToClassB    relationType
	CardinalityToClassA relationCardinality
	CardinalityToClassB relationCardinality
	Link                relationLink
	Label               string
}

// NewRelation creates a new Relation between two classes.
// It initializes the relation with a default solid link.
func NewRelation(classA *Class, classB *Class) (newRelation *Relation) {
	newRelation = &Relation{
		ClassA: classA,
		ClassB: classB,
		Link:   RelationLinkSolid,
	}

	return
}

// String generates the Mermaid syntax representation of the relationship.
// It includes the related classes, relationship types, cardinalities, link style, and optional label.
func (r *Relation) String() string {
	var sb strings.Builder

	label := ""
	if len(r.Label) > 0 {
		label = fmt.Sprintf(string(baseRelationTextString), r.Label)
	}

	sb.WriteString(fmt.Sprintf(string(baseRelationString), r.ClassA.Name, r.CardinalityToClassA, r.RelationToClassA, r.Link, r.RelationToClassB, r.CardinalityToClassB, r.ClassB.Name, label))

	return sb.String()
}
