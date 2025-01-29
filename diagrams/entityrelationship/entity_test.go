package entityrelationship

import (
	"strings"
	"testing"
)

func TestNewEntity(t *testing.T) {
	name := "TEST_TABLE"
	entity := NewEntity(name)

	if entity.Name != name {
		t.Errorf("NewEntity().Name = %v, want %v", entity.Name, name)
	}
	if len(entity.Attributes) != 0 {
		t.Error("NewEntity() should create empty attributes slice")
	}
	if entity.Alias != "" {
		t.Error("NewEntity() should not set an alias")
	}
}

func TestEntity_SetAlias(t *testing.T) {
	entity := NewEntity("TEST_TABLE")
	alias := "Test"

	result := entity.SetAlias(alias)

	if result != entity {
		t.Error("SetAlias() should return the entity for chaining")
	}
	if entity.Alias != alias {
		t.Errorf("SetAlias() = %v, want %v", entity.Alias, alias)
	}
}

func TestEntity_AddAttribute(t *testing.T) {
	entity := NewEntity("TEST_TABLE")

	attr := entity.AddAttribute("id", TypeInteger)

	if len(entity.Attributes) != 1 {
		t.Error("AddAttribute() should add attribute to entity")
	}
	if attr.Name != "id" {
		t.Errorf("AddAttribute().Name = %v, want %v", attr.Name, "id")
	}
	if attr.Type != TypeInteger {
		t.Errorf("AddAttribute().Type = %v, want %v", attr.Type, TypeInteger)
	}
}

func TestAttribute_Setters(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(*Attribute)
		checkPK  bool
		checkFK  bool
		checkReq bool
	}{
		{
			name:     "Set Primary Key",
			setup:    func(a *Attribute) { a.SetPrimaryKey() },
			checkPK:  true,
			checkFK:  false,
			checkReq: false,
		},
		{
			name:     "Set Foreign Key",
			setup:    func(a *Attribute) { a.SetForeignKey() },
			checkPK:  false,
			checkFK:  true,
			checkReq: false,
		},
		{
			name:     "Set Required",
			setup:    func(a *Attribute) { a.SetRequired() },
			checkPK:  false,
			checkFK:  false,
			checkReq: true,
		},
		{
			name: "Set All",
			setup: func(a *Attribute) {
				a.SetPrimaryKey().SetForeignKey().SetRequired()
			},
			checkPK:  true,
			checkFK:  true,
			checkReq: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attr := &Attribute{Name: "test", Type: TypeString}
			tt.setup(attr)

			if attr.PK != tt.checkPK {
				t.Errorf("PK = %v, want %v", attr.PK, tt.checkPK)
			}
			if attr.FK != tt.checkFK {
				t.Errorf("FK = %v, want %v", attr.FK, tt.checkFK)
			}
			if attr.Required != tt.checkReq {
				t.Errorf("Required = %v, want %v", attr.Required, tt.checkReq)
			}
		})
	}
}

func TestEntity_String(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() *Entity
		contains []string
		excludes []string
	}{
		{
			name: "Basic entity without alias",
			setup: func() *Entity {
				e := NewEntity("TEST")
				e.AddAttribute("id", TypeInteger).SetPrimaryKey()
				return e
			},
			contains: []string{
				"TEST {",
				"int id PK",
				"}",
			},
		},
		{
			name: "Entity with alias",
			setup: func() *Entity {
				e := NewEntity("TEST_TABLE")
				e.SetAlias("Test")
				e.AddAttribute("id", TypeInteger).SetForeignKey()
				return e
			},
			contains: []string{
				"TEST_TABLE [Test] {",
				"int id FK",
				"}",
			},
		},
		{
			name: "Entity with PK and FK",
			setup: func() *Entity {
				e := NewEntity("TEST")
				e.AddAttribute("id", TypeInteger).SetPrimaryKey().SetForeignKey()
				return e
			},
			contains: []string{
				"int id PK,FK",
			},
		},
		{
			name: "Entity with all attribute types",
			setup: func() *Entity {
				e := NewEntity("TEST")
				e.AddAttribute("id", TypeInteger)
				e.AddAttribute("name", TypeString)
				e.AddAttribute("active", TypeBoolean)
				e.AddAttribute("price", TypeFloat)
				e.AddAttribute("created", TypeDateTime)
				return e
			},
			contains: []string{
				"int id",
				"string name",
				"boolean active",
				"float price",
				"datetime created",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entity := tt.setup()
			result := entity.String()

			for _, want := range tt.contains {
				if !strings.Contains(result, want) {
					t.Errorf("String() missing expected content %q in:\n%s", want, result)
				}
			}

			for _, unwanted := range tt.excludes {
				if strings.Contains(result, unwanted) {
					t.Errorf("String() contains unexpected content %q in:\n%s", unwanted, result)
				}
			}
		})
	}
}
