package state

import (
	"strings"
	"testing"

	"github.com/TyphonHill/go-mermaid/diagrams/utils/basediagram"
)

func TestNewStateConfigurationProperties(t *testing.T) {
	got := NewStateConfigurationProperties()

	if got.properties == nil {
		t.Error("NewStateConfigurationProperties() properties map is nil")
	}

	if len(got.properties) != 0 {
		t.Errorf("NewStateConfigurationProperties() properties map length = %v, want 0", len(got.properties))
	}
}

func TestStateConfigurationProperties_String(t *testing.T) {
	tests := []struct {
		name     string
		config   StateConfigurationProperties
		setup    func(*StateConfigurationProperties)
		contains []string
	}{
		{
			name:   "Empty configuration",
			config: NewStateConfigurationProperties(),
			contains: []string{
				"",
			},
		},
		{
			name:   "Configuration with single property",
			config: NewStateConfigurationProperties(),
			setup: func(c *StateConfigurationProperties) {
				c.SetPadding(10)
			},
			contains: []string{
				"state:",
				"padding: 10",
			},
		},
		{
			name:   "Configuration with multiple properties",
			config: NewStateConfigurationProperties(),
			setup: func(c *StateConfigurationProperties) {
				c.SetPadding(10)
				c.SetRadius(5)
				c.SetArrowMarkerAbsolute(true)
			},
			contains: []string{
				"state:",
				"padding: 10",
				"radius: 5",
				"arrowMarkerAbsolute: true",
			},
		},
		{
			name:   "Configuration with base properties",
			config: NewStateConfigurationProperties(),
			setup: func(c *StateConfigurationProperties) {
				c.ConfigurationProperties = basediagram.NewConfigurationProperties()
				c.ConfigurationProperties.SetFontSize(12)
				c.SetPadding(10)
			},
			contains: []string{
				"fontSize: 12",
				"state:",
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

func TestStateConfigurationProperties_Setters(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(*StateConfigurationProperties) *StateConfigurationProperties
		property string
		value    interface{}
	}{
		{
			name: "Set title top margin",
			setup: func(c *StateConfigurationProperties) *StateConfigurationProperties {
				return c.SetTitleTopMargin(10)
			},
			property: statePropertyTitleTopMargin,
			value:    10,
		},
		{
			name: "Set arrow marker absolute",
			setup: func(c *StateConfigurationProperties) *StateConfigurationProperties {
				return c.SetArrowMarkerAbsolute(true)
			},
			property: statePropertyArrowMarkerAbsolute,
			value:    true,
		},
		{
			name: "Set divider margin",
			setup: func(c *StateConfigurationProperties) *StateConfigurationProperties {
				return c.SetDividerMargin(15)
			},
			property: statePropertyDividerMargin,
			value:    15,
		},
		{
			name: "Set size unit",
			setup: func(c *StateConfigurationProperties) *StateConfigurationProperties {
				return c.SetSizeUnit(5)
			},
			property: statePropertySizeUnit,
			value:    5,
		},
		{
			name: "Set padding",
			setup: func(c *StateConfigurationProperties) *StateConfigurationProperties {
				return c.SetPadding(10)
			},
			property: statePropertyPadding,
			value:    10,
		},
		{
			name: "Set text height",
			setup: func(c *StateConfigurationProperties) *StateConfigurationProperties {
				return c.SetTextHeight(20)
			},
			property: statePropertyTextHeight,
			value:    20,
		},
		{
			name: "Set title shift",
			setup: func(c *StateConfigurationProperties) *StateConfigurationProperties {
				return c.SetTitleShift(5)
			},
			property: statePropertyTitleShift,
			value:    5,
		},
		{
			name: "Set note margin",
			setup: func(c *StateConfigurationProperties) *StateConfigurationProperties {
				return c.SetNoteMargin(8)
			},
			property: statePropertyNoteMargin,
			value:    8,
		},
		{
			name: "Set node spacing",
			setup: func(c *StateConfigurationProperties) *StateConfigurationProperties {
				return c.SetNodeSpacing(25)
			},
			property: statePropertyNodeSpacing,
			value:    25,
		},
		{
			name: "Set rank spacing",
			setup: func(c *StateConfigurationProperties) *StateConfigurationProperties {
				return c.SetRankSpacing(30)
			},
			property: statePropertyRankSpacing,
			value:    30,
		},
		{
			name: "Set fork width",
			setup: func(c *StateConfigurationProperties) *StateConfigurationProperties {
				return c.SetForkWidth(40)
			},
			property: statePropertyForkWidth,
			value:    40,
		},
		{
			name: "Set fork height",
			setup: func(c *StateConfigurationProperties) *StateConfigurationProperties {
				return c.SetForkHeight(30)
			},
			property: statePropertyForkHeight,
			value:    30,
		},
		{
			name: "Set mini padding",
			setup: func(c *StateConfigurationProperties) *StateConfigurationProperties {
				return c.SetMiniPadding(4)
			},
			property: statePropertyMiniPadding,
			value:    4,
		},
		{
			name: "Set font size factor",
			setup: func(c *StateConfigurationProperties) *StateConfigurationProperties {
				return c.SetFontSizeFactor(2)
			},
			property: statePropertyFontSizeFactor,
			value:    2,
		},
		{
			name: "Set font size",
			setup: func(c *StateConfigurationProperties) *StateConfigurationProperties {
				return c.SetFontSize(14)
			},
			property: statePropertyFontSize,
			value:    14,
		},
		{
			name: "Set label height",
			setup: func(c *StateConfigurationProperties) *StateConfigurationProperties {
				return c.SetLabelHeight(16)
			},
			property: statePropertyLabelHeight,
			value:    16,
		},
		{
			name: "Set edge length factor",
			setup: func(c *StateConfigurationProperties) *StateConfigurationProperties {
				return c.SetEdgeLengthFactor("1.5")
			},
			property: statePropertyEdgeLengthFactor,
			value:    "1.5",
		},
		{
			name: "Set composit title size",
			setup: func(c *StateConfigurationProperties) *StateConfigurationProperties {
				return c.SetCompositTitleSize(18)
			},
			property: statePropertyCompositTitleSize,
			value:    18,
		},
		{
			name: "Set radius",
			setup: func(c *StateConfigurationProperties) *StateConfigurationProperties {
				return c.SetRadius(5)
			},
			property: statePropertyRadius,
			value:    5,
		},
		{
			name: "Set default renderer",
			setup: func(c *StateConfigurationProperties) *StateConfigurationProperties {
				return c.SetDefaultRenderer("dagre")
			},
			property: statePropertyDefaultRenderer,
			value:    "dagre",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := NewStateConfigurationProperties()
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
