package entityrelationship

import (
	"strings"
	"testing"
)

func TestNewRelationship(t *testing.T) {
	from := NewEntity("FROM")
	to := NewEntity("TO")

	rel := NewRelationship(from, to)

	if rel.From != from {
		t.Errorf("NewRelationship().From = %v, want %v", rel.From, from)
	}
	if rel.To != to {
		t.Errorf("NewRelationship().To = %v, want %v", rel.To, to)
	}
	if rel.Cardinality != ExactlyOne {
		t.Errorf("NewRelationship().Cardinality = %v, want %v", rel.Cardinality, ExactlyOne)
	}
	if rel.Label != "" {
		t.Error("NewRelationship() should not set a label")
	}
}

func TestRelationship_SetLabel(t *testing.T) {
	rel := NewRelationship(NewEntity("A"), NewEntity("B"))
	label := "test_label"

	result := rel.SetLabel(label)

	if result != rel {
		t.Error("SetLabel() should return the relationship for chaining")
	}
	if rel.Label != label {
		t.Errorf("SetLabel() = %v, want %v", rel.Label, label)
	}
}

func TestRelationship_SetCardinality(t *testing.T) {
	rel := NewRelationship(NewEntity("A"), NewEntity("B"))
	card := OneToZeroOrMore

	result := rel.SetCardinality(card)

	if result != rel {
		t.Error("SetCardinality() should return the relationship for chaining")
	}
	if rel.Cardinality != card {
		t.Errorf("SetCardinality() = %v, want %v", rel.Cardinality, card)
	}
}

func TestRelationship_String(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() *Relationship
		contains []string
	}{
		{
			name: "Basic relationship without label",
			setup: func() *Relationship {
				return NewRelationship(
					NewEntity("A"),
					NewEntity("B"),
				)
			},
			contains: []string{
				"A",
				"B",
				"relates",
				string(ExactlyOne),
			},
		},
		{
			name: "Relationship with label",
			setup: func() *Relationship {
				rel := NewRelationship(
					NewEntity("User"),
					NewEntity("Post"),
				)
				rel.SetLabel("writes")
				return rel
			},
			contains: []string{
				"User",
				"Post",
				"writes",
			},
		},
		{
			name: "Relationship with custom cardinality",
			setup: func() *Relationship {
				rel := NewRelationship(
					NewEntity("Book"),
					NewEntity("Author"),
				)
				rel.SetCardinality(ManyToMany)
				return rel
			},
			contains: []string{
				"Book",
				"Author",
				string(ManyToMany),
			},
		},
		{
			name: "Relationship with aliased entities",
			setup: func() *Relationship {
				from := NewEntity("USERS")
				from.SetAlias("User")
				to := NewEntity("POSTS")
				to.SetAlias("Post")
				return NewRelationship(from, to)
			},
			contains: []string{
				"USERS",
				"POSTS",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rel := tt.setup()
			result := rel.String()

			for _, want := range tt.contains {
				if !strings.Contains(result, want) {
					t.Errorf("String() missing expected content %q in:\n%s", want, result)
				}
			}
		})
	}
}
