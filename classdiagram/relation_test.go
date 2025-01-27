package classdiagram

import (
	"reflect"
	"testing"
)

func TestNewRelation(t *testing.T) {
	type args struct {
		classA *Class
		classB *Class
	}
	tests := []struct {
		name            string
		args            args
		wantNewRelation *Relation
	}{
		{
			name: "Nominal test",
			args: args{
				classA: NewClass("Test1"),
				classB: NewClass("Test2"),
			},
			wantNewRelation: &Relation{
				ClassA: &Class{Name: "Test1"},
				ClassB: &Class{Name: "Test2"},
				Link:   RelationLinkSolid,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewRelation := NewRelation(tt.args.classA, tt.args.classB); !reflect.DeepEqual(gotNewRelation, tt.wantNewRelation) {
				t.Errorf("NewRelation() = %v, want %v", gotNewRelation, tt.wantNewRelation)
			}
		})
	}
}

func TestRelation_String(t *testing.T) {
	class1 := NewClass("Test1")
	class2 := NewClass("Test2")

	rel := NewRelation(class1, class2)
	rel.RelationToClassB = RelationTypeAssociation
	rel.RelationToClassA = RelationTypeAggregation
	rel.CardinalityToClassA = RelationCardinalityMany
	rel.CardinalityToClassB = RelationCardinalityZeroToN
	rel.Link = RelationLinkDashed
	rel.Label = "TestLabel"

	tests := []struct {
		name     string
		relation *Relation
		want     string
	}{
		{
			name:     "Nominal test",
			relation: NewRelation(class1, class2),
			want:     "\tTest1 -- Test2\n",
		},
		{
			name:     "Nominal test",
			relation: rel,
			want:     "\tTest1 \"*\"o..>\"0..n\" Test2 : TestLabel\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Relation{
				ClassA:              tt.relation.ClassA,
				ClassB:              tt.relation.ClassB,
				RelationToClassA:    tt.relation.RelationToClassA,
				RelationToClassB:    tt.relation.RelationToClassB,
				CardinalityToClassA: tt.relation.CardinalityToClassA,
				CardinalityToClassB: tt.relation.CardinalityToClassB,
				Link:                tt.relation.Link,
				Label:               tt.relation.Label,
			}
			if got := r.String(); got != tt.want {
				t.Errorf("Relation.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
