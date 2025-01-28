package classdiagram

import (
	"strings"
	"testing"
)

func TestNewRelation(t *testing.T) {
	// Create two test classes
	class1 := NewClass("ClassA")
	class2 := NewClass("ClassB")

	tests := []struct {
		name   string
		classA *Class
		classB *Class
		want   *Relation
	}{
		{
			name:   "Create new relation",
			classA: class1,
			classB: class2,
			want: &Relation{
				ClassA: class1,
				ClassB: class2,
				Link:   RelationLinkSolid,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewRelation(tt.classA, tt.classB)

			if got.ClassA != tt.want.ClassA {
				t.Errorf("NewRelation() ClassA = %v, want %v", got.ClassA, tt.want.ClassA)
			}

			if got.ClassB != tt.want.ClassB {
				t.Errorf("NewRelation() ClassB = %v, want %v", got.ClassB, tt.want.ClassB)
			}

			if got.Link != tt.want.Link {
				t.Errorf("NewRelation() Link = %v, want %v", got.Link, tt.want.Link)
			}
		})
	}
}

func TestRelation_String(t *testing.T) {
	tests := []struct {
		name     string
		relation *Relation
		contains []string
	}{
		{
			name: "Basic relation with default values",
			relation: func() *Relation {
				class1 := NewClass("ClassA")
				class2 := NewClass("ClassB")
				return NewRelation(class1, class2)
			}(),
			contains: []string{
				"ClassA -- ClassB",
			},
		},
		{
			name: "Relation with cardinality",
			relation: func() *Relation {
				class1 := NewClass("User")
				class2 := NewClass("Order")
				relation := NewRelation(class1, class2)
				relation.CardinalityToClassA = RelationCardinalityOneOrMore
				relation.CardinalityToClassB = RelationCardinalityZeroOrOne
				return relation
			}(),
			contains: []string{
				`User "1..*"--"0..1" Order`,
			},
		},
		{
			name: "Relation with relation types and link",
			relation: func() *Relation {
				class1 := NewClass("Parent")
				class2 := NewClass("Child")
				relation := NewRelation(class1, class2)
				relation.RelationToClassA = RelationTypeInheritance
				relation.RelationToClassB = RelationTypeAggregation
				relation.Link = RelationLinkDashed
				return relation
			}(),
			contains: []string{
				"Parent |>..o Child",
			},
		},
		{
			name: "Relation with label",
			relation: func() *Relation {
				class1 := NewClass("ClassA")
				class2 := NewClass("ClassB")
				relation := NewRelation(class1, class2)
				relation.Label = "Test Relation"
				return relation
			}(),
			contains: []string{
				"ClassA -- ClassB : Test Relation",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := tt.relation.String()

			for _, expectedContent := range tt.contains {
				if !strings.Contains(output, expectedContent) {
					t.Errorf("String() output missing expected content: %q", expectedContent)
				}
			}
		})
	}
}

// TestRelationConstants checks the predefined constants for relations
func TestRelationConstants(t *testing.T) {
	relationTypeTests := []struct {
		name      string
		typeConst relationType
		expected  string
	}{
		{"Association", RelationTypeAssociation, ">"},
		{"Association Left", RelationTypeAssociationLeft, "<"},
		{"Inheritance", RelationTypeInheritance, "|>"},
		{"Inheritance Left", RelationTypeInheritanceLeft, "<|"},
		{"Composition", RelationTypeComposition, "*"},
		{"Aggregation", RelationTypeAggregation, "o"},
	}

	for _, tt := range relationTypeTests {
		t.Run(tt.name, func(t *testing.T) {
			if string(tt.typeConst) != tt.expected {
				t.Errorf("Relation type %s = %v, want %v", tt.name, tt.typeConst, tt.expected)
			}
		})
	}

	relationLinkTests := []struct {
		name      string
		linkConst relationLink
		expected  string
	}{
		{"Solid Link", RelationLinkSolid, "--"},
		{"Dashed Link", RelationLinkDashed, ".."},
	}

	for _, tt := range relationLinkTests {
		t.Run(tt.name, func(t *testing.T) {
			if string(tt.linkConst) != tt.expected {
				t.Errorf("Relation link %s = %v, want %v", tt.name, tt.linkConst, tt.expected)
			}
		})
	}

	cardinalityTests := []struct {
		name             string
		cardinalityConst relationCardinality
		expected         string
	}{
		{"Only One", RelationCardinalityOnlyOne, "\"1\""},
		{"Zero or One", RelationCardinalityZeroOrOne, "\"0..1\""},
		{"One or More", RelationCardinalityOneOrMore, "\"1..*\""},
		{"Many", RelationCardinalityMany, "\"*\""},
		{"N", RelationCardinalityN, "\"n\""},
		{"Zero to N", RelationCardinalityZeroToN, "\"0..n\""},
		{"One to N", RelationCardinalityOneToN, "\"1..n\""},
	}

	for _, tt := range cardinalityTests {
		t.Run(tt.name, func(t *testing.T) {
			if string(tt.cardinalityConst) != tt.expected {
				t.Errorf("Cardinality %s = %v, want %v", tt.name, tt.cardinalityConst, tt.expected)
			}
		})
	}
}
