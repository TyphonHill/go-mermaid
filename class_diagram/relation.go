package class_diagram

import (
	"fmt"
	"strings"
)

type relationType string
type relationLink string
type relationCardinality string

const (
	RelationTypeAssociation     relationType = ">"
	RelationTypeAssociationLeft relationType = "<"
	RelationTypeInheritance     relationType = "|>"
	RelationTypeInheritanceLeft relationType = "<|"
	RelationTypeComposition     relationType = "*"
	RelationTypeAggregation     relationType = "o"
)

const (
	RelationLinkSolid  relationLink = "--"
	RelationLinkDashed relationLink = ".."
)

const (
	RelationCardinalityOnlyOne   relationCardinality = "\"1\""
	RelationCardinalityZeroOrOne relationCardinality = "\"0..1\""
	RelationCardinalityOneOrMore relationCardinality = "\"1..*\""
	RelationCardinalityMany      relationCardinality = "\"*\""
	RelationCardinalityN         relationCardinality = "\"n\""
	RelationCardinalityZeroToN   relationCardinality = "\"0..n\""
	RelationCardinalityOneToN    relationCardinality = "\"1..n\""
)

const (
	baseRelationString     string = "\t%s %s%s%s%s%s %s%s\n"
	baseRelationTextString string = " : %s"
)

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

func NewRelation(classA *Class, classB *Class) (newRelation *Relation) {
	newRelation = &Relation{
		ClassA: classA,
		ClassB: classB,
		Link:   RelationLinkSolid,
	}

	return
}

func (r *Relation) String() string {
	var sb strings.Builder

	label := ""
	if len(r.Label) > 0 {
		label = fmt.Sprintf(string(baseRelationTextString), r.Label)
	}

	sb.WriteString(fmt.Sprintf(string(baseRelationString), r.ClassA.Name, r.CardinalityToClassA, r.RelationToClassA, r.Link, r.RelationToClassB, r.CardinalityToClassB, r.ClassB.Name, label))

	return sb.String()
}
