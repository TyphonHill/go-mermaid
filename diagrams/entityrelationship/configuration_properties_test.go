package entityrelationship

import (
	"strings"
	"testing"

	"github.com/TyphonHill/go-mermaid/diagrams/utils/basediagram"
)

func TestNewErConfigurationProperties(t *testing.T) {
	got := NewErConfigurationProperties()

	if got.properties == nil {
		t.Error("NewErConfigurationProperties() properties map is nil")
	}

	if len(got.properties) != 0 {
		t.Errorf("NewErConfigurationProperties() properties map length = %v, want 0", len(got.properties))
	}
}

func TestErConfigurationProperties_String(t *testing.T) {
	tests := []struct {
		name     string
		config   ErConfigurationProperties
		setup    func(*ErConfigurationProperties)
		contains []string
	}{
		{
			name:   "Empty configuration",
			config: NewErConfigurationProperties(),
			contains: []string{
				"",
			},
		},
		{
			name:   "Configuration with single property",
			config: NewErConfigurationProperties(),
			setup: func(c *ErConfigurationProperties) {
				c.SetFontSize(12)
			},
			contains: []string{
				"er:",
				"fontSize: 12",
			},
		},
		{
			name:   "Configuration with multiple properties",
			config: NewErConfigurationProperties(),
			setup: func(c *ErConfigurationProperties) {
				c.SetFontSize(12)
				c.SetStroke("#333")
				c.SetFill("#fff")
			},
			contains: []string{
				"er:",
				"fontSize: 12",
				"stroke: #333",
				"fill: #fff",
			},
		},
		{
			name:   "Configuration with base properties",
			config: NewErConfigurationProperties(),
			setup: func(c *ErConfigurationProperties) {
				c.ConfigurationProperties = basediagram.NewConfigurationProperties()
				c.ConfigurationProperties.SetFontSize(12)
				c.SetEntityPadding(20)
			},
			contains: []string{
				"fontSize: 12",
				"er:",
				"entityPadding: 20",
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

func TestErConfigurationProperties_Setters(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(*ErConfigurationProperties) *ErConfigurationProperties
		property string
		value    interface{}
	}{
		{
			name: "Set title top margin",
			setup: func(c *ErConfigurationProperties) *ErConfigurationProperties {
				return c.SetTitleTopMargin(10)
			},
			property: erPropertyTitleTopMargin,
			value:    10,
		},
		{
			name: "Set diagram padding",
			setup: func(c *ErConfigurationProperties) *ErConfigurationProperties {
				return c.SetDiagramPadding(20)
			},
			property: erPropertyDiagramPadding,
			value:    20,
		},
		{
			name: "Set layout direction",
			setup: func(c *ErConfigurationProperties) *ErConfigurationProperties {
				return c.SetLayoutDirection("TB")
			},
			property: erPropertyLayoutDirection,
			value:    "TB",
		},
		{
			name: "Set min entity width",
			setup: func(c *ErConfigurationProperties) *ErConfigurationProperties {
				return c.SetMinEntityWidth(100)
			},
			property: erPropertyMinEntityWidth,
			value:    100,
		},
		{
			name: "Set min entity height",
			setup: func(c *ErConfigurationProperties) *ErConfigurationProperties {
				return c.SetMinEntityHeight(50)
			},
			property: erPropertyMinEntityHeight,
			value:    50,
		},
		{
			name: "Set entity padding",
			setup: func(c *ErConfigurationProperties) *ErConfigurationProperties {
				return c.SetEntityPadding(10)
			},
			property: erPropertyEntityPadding,
			value:    10,
		},
		{
			name: "Set stroke",
			setup: func(c *ErConfigurationProperties) *ErConfigurationProperties {
				return c.SetStroke("#333")
			},
			property: erPropertyStroke,
			value:    "#333",
		},
		{
			name: "Set fill",
			setup: func(c *ErConfigurationProperties) *ErConfigurationProperties {
				return c.SetFill("#fff")
			},
			property: erPropertyFill,
			value:    "#fff",
		},
		{
			name: "Set font size",
			setup: func(c *ErConfigurationProperties) *ErConfigurationProperties {
				return c.SetFontSize(12)
			},
			property: erPropertyFontSize,
			value:    12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := NewErConfigurationProperties()
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
			}

			if got != tt.value {
				t.Errorf("Property %q = %v, want %v", tt.property, got, tt.value)
			}
		})
	}
}
