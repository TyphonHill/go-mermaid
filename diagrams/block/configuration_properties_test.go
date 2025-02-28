package block

import (
	"strings"
	"testing"

	"github.com/TyphonHill/go-mermaid/diagrams/utils/basediagram"
)

func TestBlockConfigurationProperties_String(t *testing.T) {
	tests := []struct {
		name     string
		config   BlockConfigurationProperties
		setup    func(*BlockConfigurationProperties)
		contains []string
	}{
		{
			name:   "Empty configuration",
			config: NewBlockConfigurationProperties(),
			contains: []string{
				"",
			},
		},
		{
			name:   "Configuration with padding",
			config: NewBlockConfigurationProperties(),
			setup: func(c *BlockConfigurationProperties) {
				c.SetPadding(10)
			},
			contains: []string{
				"block:",
				"padding: 10",
			},
		},
		{
			name:   "Configuration with base properties and padding",
			config: NewBlockConfigurationProperties(),
			setup: func(c *BlockConfigurationProperties) {
				c.SetPadding(10)
				c.ConfigurationProperties = basediagram.NewConfigurationProperties()
				c.ConfigurationProperties.SetFontSize(12)
			},
			contains: []string{
				"fontSize: 12",
				"block:",
				"padding: 10",
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

func TestNewBlockConfigurationProperties(t *testing.T) {
	got := NewBlockConfigurationProperties()

	if got.properties == nil {
		t.Error("NewBlockConfigurationProperties() properties map is nil")
	}

	if len(got.properties) != 0 {
		t.Errorf("NewBlockConfigurationProperties() properties map length = %v, want 0", len(got.properties))
	}
}

func TestBlockConfigurationProperties_Setters(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(*BlockConfigurationProperties) *BlockConfigurationProperties
		property string
		value    interface{}
	}{
		{
			name: "Set padding",
			setup: func(c *BlockConfigurationProperties) *BlockConfigurationProperties {
				return c.SetPadding(10)
			},
			property: blockPropertyPadding,
			value:    10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := NewBlockConfigurationProperties()
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
