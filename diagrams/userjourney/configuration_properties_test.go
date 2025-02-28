package userjourney

import (
	"reflect"
	"strings"
	"testing"

	"github.com/TyphonHill/go-mermaid/diagrams/utils/basediagram"
)

func TestNewJourneyConfigurationProperties(t *testing.T) {
	got := NewJourneyConfigurationProperties()

	if got.properties == nil {
		t.Error("NewJourneyConfigurationProperties() properties map is nil")
	}

	if len(got.properties) != 0 {
		t.Errorf("NewJourneyConfigurationProperties() properties map length = %v, want 0", len(got.properties))
	}
}

func TestJourneyConfigurationProperties_String(t *testing.T) {
	tests := []struct {
		name     string
		config   JourneyConfigurationProperties
		setup    func(*JourneyConfigurationProperties)
		contains []string
	}{
		{
			name:   "Empty configuration",
			config: NewJourneyConfigurationProperties(),
			contains: []string{
				"",
			},
		},
		{
			name:   "Configuration with single property",
			config: NewJourneyConfigurationProperties(),
			setup: func(c *JourneyConfigurationProperties) {
				c.SetWidth(800)
			},
			contains: []string{
				"journey:",
				"width: 800",
			},
		},
		{
			name:   "Configuration with multiple properties",
			config: NewJourneyConfigurationProperties(),
			setup: func(c *JourneyConfigurationProperties) {
				c.SetWidth(800)
				c.SetRightAngles(true)
				c.SetTaskFontFamily("Arial")
			},
			contains: []string{
				"journey:",
				"width: 800",
				"rightAngles: true",
				"taskFontFamily: Arial",
			},
		},
		{
			name:   "Configuration with base properties",
			config: NewJourneyConfigurationProperties(),
			setup: func(c *JourneyConfigurationProperties) {
				c.ConfigurationProperties = basediagram.NewConfigurationProperties()
				c.ConfigurationProperties.SetFontSize(12)
				c.SetWidth(800)
			},
			contains: []string{
				"fontSize: 12",
				"journey:",
				"width: 800",
			},
		},
		{
			name:   "Configuration with array properties",
			config: NewJourneyConfigurationProperties(),
			setup: func(c *JourneyConfigurationProperties) {
				c.SetActorColours([]string{"#ff0000", "#00ff00"})
				c.SetSectionFills([]string{"#f9f9f9", "#f1f1f1"})
			},
			contains: []string{
				"journey:",
				"actorColours: [\"#ff0000\", \"#00ff00\"]",
				"sectionFills: [\"#f9f9f9\", \"#f1f1f1\"]",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup(&tt.config)
			}

			got := tt.config.String()
			for _, want := range tt.contains {
				if !strings.Contains(got, want) {
					t.Errorf("String() missing expected content %q in:\n%s", want, got)
				}
			}
		})
	}
}

func TestJourneyConfigurationProperties_Setters(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(*JourneyConfigurationProperties) *JourneyConfigurationProperties
		property string
		value    interface{}
	}{
		{
			name: "Set diagram margin X",
			setup: func(c *JourneyConfigurationProperties) *JourneyConfigurationProperties {
				return c.SetDiagramMarginX(20)
			},
			property: journeyPropertyDiagramMarginX,
			value:    20,
		},
		{
			name: "Set diagram margin Y",
			setup: func(c *JourneyConfigurationProperties) *JourneyConfigurationProperties {
				return c.SetDiagramMarginY(30)
			},
			property: journeyPropertyDiagramMarginY,
			value:    30,
		},
		{
			name: "Set left margin",
			setup: func(c *JourneyConfigurationProperties) *JourneyConfigurationProperties {
				return c.SetLeftMargin(15)
			},
			property: journeyPropertyLeftMargin,
			value:    15,
		},
		{
			name: "Set width",
			setup: func(c *JourneyConfigurationProperties) *JourneyConfigurationProperties {
				return c.SetWidth(800)
			},
			property: journeyPropertyWidth,
			value:    800,
		},
		{
			name: "Set height",
			setup: func(c *JourneyConfigurationProperties) *JourneyConfigurationProperties {
				return c.SetHeight(600)
			},
			property: journeyPropertyHeight,
			value:    600,
		},
		{
			name: "Set box margin",
			setup: func(c *JourneyConfigurationProperties) *JourneyConfigurationProperties {
				return c.SetBoxMargin(10)
			},
			property: journeyPropertyBoxMargin,
			value:    10,
		},
		{
			name: "Set box text margin",
			setup: func(c *JourneyConfigurationProperties) *JourneyConfigurationProperties {
				return c.SetBoxTextMargin(5)
			},
			property: journeyPropertyBoxTextMargin,
			value:    5,
		},
		{
			name: "Set note margin",
			setup: func(c *JourneyConfigurationProperties) *JourneyConfigurationProperties {
				return c.SetNoteMargin(8)
			},
			property: journeyPropertyNoteMargin,
			value:    8,
		},
		{
			name: "Set message margin",
			setup: func(c *JourneyConfigurationProperties) *JourneyConfigurationProperties {
				return c.SetMessageMargin(12)
			},
			property: journeyPropertyMessageMargin,
			value:    12,
		},
		{
			name: "Set message align",
			setup: func(c *JourneyConfigurationProperties) *JourneyConfigurationProperties {
				return c.SetMessageAlign("left")
			},
			property: journeyPropertyMessageAlign,
			value:    "left",
		},
		{
			name: "Set bottom margin adj",
			setup: func(c *JourneyConfigurationProperties) *JourneyConfigurationProperties {
				return c.SetBottomMarginAdj(25)
			},
			property: journeyPropertyBottomMarginAdj,
			value:    25,
		},
		{
			name: "Set right angles",
			setup: func(c *JourneyConfigurationProperties) *JourneyConfigurationProperties {
				return c.SetRightAngles(true)
			},
			property: journeyPropertyRightAngles,
			value:    true,
		},
		{
			name: "Set task font size",
			setup: func(c *JourneyConfigurationProperties) *JourneyConfigurationProperties {
				return c.SetTaskFontSize(14)
			},
			property: journeyPropertyTaskFontSize,
			value:    14,
		},
		{
			name: "Set task font family",
			setup: func(c *JourneyConfigurationProperties) *JourneyConfigurationProperties {
				return c.SetTaskFontFamily("Arial")
			},
			property: journeyPropertyTaskFontFamily,
			value:    "Arial",
		},
		{
			name: "Set task margin",
			setup: func(c *JourneyConfigurationProperties) *JourneyConfigurationProperties {
				return c.SetTaskMargin(5)
			},
			property: journeyPropertyTaskMargin,
			value:    5,
		},
		{
			name: "Set activation width",
			setup: func(c *JourneyConfigurationProperties) *JourneyConfigurationProperties {
				return c.SetActivationWidth(2)
			},
			property: journeyPropertyActivationWidth,
			value:    2,
		},
		{
			name: "Set text placement",
			setup: func(c *JourneyConfigurationProperties) *JourneyConfigurationProperties {
				return c.SetTextPlacement("top")
			},
			property: journeyPropertyTextPlacement,
			value:    "top",
		},
		{
			name: "Set actor colours",
			setup: func(c *JourneyConfigurationProperties) *JourneyConfigurationProperties {
				return c.SetActorColours([]string{"#ff0000", "#00ff00"})
			},
			property: journeyPropertyActorColours,
			value:    []string{"#ff0000", "#00ff00"},
		},
		{
			name: "Set section fills",
			setup: func(c *JourneyConfigurationProperties) *JourneyConfigurationProperties {
				return c.SetSectionFills([]string{"#f9f9f9", "#f1f1f1"})
			},
			property: journeyPropertySectionFills,
			value:    []string{"#f9f9f9", "#f1f1f1"},
		},
		{
			name: "Set section colours",
			setup: func(c *JourneyConfigurationProperties) *JourneyConfigurationProperties {
				return c.SetSectionColours([]string{"#333333", "#666666"})
			},
			property: journeyPropertySectionColours,
			value:    []string{"#333333", "#666666"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := NewJourneyConfigurationProperties()
			result := tt.setup(&config)

			// Test method chaining
			if result != &config {
				t.Error("Setter should return pointer to config for chaining")
			}

			// Test property was set
			prop, exists := config.properties[tt.property]
			if !exists {
				t.Errorf("Property %q was not set", tt.property)
				return
			}

			// Test property value
			var got interface{}
			switch p := prop.(type) {
			case *basediagram.IntProperty:
				got = p.Val
			case *basediagram.BoolProperty:
				got = p.Val
			case *basediagram.StringProperty:
				got = p.Val
			case *basediagram.StringArrayProperty:
				got = p.Val
			}

			if !reflect.DeepEqual(got, tt.value) {
				t.Errorf("Property %q = %v, want %v", tt.property, got, tt.value)
			}
		})
	}
}
