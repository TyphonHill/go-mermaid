package class

import (
	"strings"
	"testing"

	"github.com/TyphonHill/go-mermaid/diagrams/utils/basediagram"
)

func TestNewClassConfigurationProperties(t *testing.T) {
	got := NewClassConfigurationProperties()

	if got.properties == nil {
		t.Error("NewClassConfigurationProperties() properties map is nil")
	}

	if len(got.properties) != 0 {
		t.Errorf("NewClassConfigurationProperties() properties map length = %v, want 0", len(got.properties))
	}
}

func TestClassConfigurationProperties_String(t *testing.T) {
	tests := []struct {
		name     string
		config   ClassConfigurationProperties
		setup    func(*ClassConfigurationProperties)
		contains []string
	}{
		{
			name:   "Empty configuration",
			config: NewClassConfigurationProperties(),
			contains: []string{
				"",
			},
		},
		{
			name:   "Configuration with single property",
			config: NewClassConfigurationProperties(),
			setup: func(c *ClassConfigurationProperties) {
				c.SetPadding(10)
			},
			contains: []string{
				"class:",
				"padding: 10",
			},
		},
		{
			name:   "Configuration with multiple properties",
			config: NewClassConfigurationProperties(),
			setup: func(c *ClassConfigurationProperties) {
				c.SetPadding(10)
				c.SetTextHeight(20)
				c.SetHtmlLabels(true)
			},
			contains: []string{
				"class:",
				"padding: 10",
				"textHeight: 20",
				"htmlLabels: true",
			},
		},
		{
			name:   "Configuration with base properties",
			config: NewClassConfigurationProperties(),
			setup: func(c *ClassConfigurationProperties) {
				c.ConfigurationProperties = basediagram.NewConfigurationProperties()
				c.ConfigurationProperties.SetFontSize(12)
				c.SetPadding(10)
			},
			contains: []string{
				"fontSize: 12",
				"class:",
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

func TestClassConfigurationProperties_Setters(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(*ClassConfigurationProperties) *ClassConfigurationProperties
		property string
		value    interface{}
	}{
		{
			name: "Set title top margin",
			setup: func(c *ClassConfigurationProperties) *ClassConfigurationProperties {
				return c.SetTitleTopMargin(10)
			},
			property: classPropertyTitleTopMargin,
			value:    10,
		},
		{
			name: "Set arrow marker absolute",
			setup: func(c *ClassConfigurationProperties) *ClassConfigurationProperties {
				return c.SetArrowMarkerAbsolute(true)
			},
			property: classPropertyArrowMarkerAbsolute,
			value:    true,
		},
		{
			name: "Set divider margin",
			setup: func(c *ClassConfigurationProperties) *ClassConfigurationProperties {
				return c.SetDividerMargin(5)
			},
			property: classPropertyDividerMargin,
			value:    5,
		},
		{
			name: "Set padding",
			setup: func(c *ClassConfigurationProperties) *ClassConfigurationProperties {
				return c.SetPadding(8)
			},
			property: classPropertyPadding,
			value:    8,
		},
		{
			name: "Set text height",
			setup: func(c *ClassConfigurationProperties) *ClassConfigurationProperties {
				return c.SetTextHeight(15)
			},
			property: classPropertyTextHeight,
			value:    15,
		},
		{
			name: "Set default renderer",
			setup: func(c *ClassConfigurationProperties) *ClassConfigurationProperties {
				return c.SetDefaultRenderer("custom")
			},
			property: classPropertyDefaultRenderer,
			value:    "custom",
		},
		{
			name: "Set node spacing",
			setup: func(c *ClassConfigurationProperties) *ClassConfigurationProperties {
				return c.SetNodeSpacing(50)
			},
			property: classPropertyNodeSpacing,
			value:    50,
		},
		{
			name: "Set rank spacing",
			setup: func(c *ClassConfigurationProperties) *ClassConfigurationProperties {
				return c.SetRankSpacing(30)
			},
			property: classPropertyRankSpacing,
			value:    30,
		},
		{
			name: "Set diagram padding",
			setup: func(c *ClassConfigurationProperties) *ClassConfigurationProperties {
				return c.SetDiagramPadding(20)
			},
			property: classPropertyDiagramPadding,
			value:    20,
		},
		{
			name: "Set HTML labels",
			setup: func(c *ClassConfigurationProperties) *ClassConfigurationProperties {
				return c.SetHtmlLabels(true)
			},
			property: classPropertyHtmlLabels,
			value:    true,
		},
		{
			name: "Set hide empty members box",
			setup: func(c *ClassConfigurationProperties) *ClassConfigurationProperties {
				return c.SetHideEmptyMembersBox(true)
			},
			property: classPropertyHideEmptyMembersBox,
			value:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := NewClassConfigurationProperties()
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
