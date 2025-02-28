package flowchart

import (
	"strings"
	"testing"

	"github.com/TyphonHill/go-mermaid/diagrams/utils/basediagram"
)

func TestNewFlowchartConfigurationProperties(t *testing.T) {
	got := NewFlowchartConfigurationProperties()

	if got.properties == nil {
		t.Error("NewFlowchartConfigurationProperties() properties map is nil")
	}

	if len(got.properties) != 0 {
		t.Errorf("NewFlowchartConfigurationProperties() properties map length = %v, want 0", len(got.properties))
	}
}

func TestFlowchartConfigurationProperties_String(t *testing.T) {
	tests := []struct {
		name     string
		config   FlowchartConfigurationProperties
		setup    func(*FlowchartConfigurationProperties)
		contains []string
	}{
		{
			name:   "Empty configuration",
			config: NewFlowchartConfigurationProperties(),
			contains: []string{
				"",
			},
		},
		{
			name:   "Configuration with single property",
			config: NewFlowchartConfigurationProperties(),
			setup: func(c *FlowchartConfigurationProperties) {
				c.SetPadding(10)
			},
			contains: []string{
				"flowchart:",
				"padding: 10",
			},
		},
		{
			name:   "Configuration with multiple properties",
			config: NewFlowchartConfigurationProperties(),
			setup: func(c *FlowchartConfigurationProperties) {
				c.SetPadding(10)
				c.SetCurve("basis")
				c.SetHtmlLabels(true)
			},
			contains: []string{
				"flowchart:",
				"padding: 10",
				"curve: basis",
				"htmlLabels: true",
			},
		},
		{
			name:   "Configuration with base properties",
			config: NewFlowchartConfigurationProperties(),
			setup: func(c *FlowchartConfigurationProperties) {
				c.ConfigurationProperties = basediagram.NewConfigurationProperties()
				c.ConfigurationProperties.SetFontSize(12)
				c.SetPadding(10)
			},
			contains: []string{
				"fontSize: 12",
				"flowchart:",
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

func TestFlowchartConfigurationProperties_Setters(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(*FlowchartConfigurationProperties) *FlowchartConfigurationProperties
		property string
		value    interface{}
	}{
		{
			name: "Set title top margin",
			setup: func(c *FlowchartConfigurationProperties) *FlowchartConfigurationProperties {
				return c.SetTitleTopMargin(10)
			},
			property: flowchartPropertyTitleTopMargin,
			value:    10,
		},
		{
			name: "Set diagram padding",
			setup: func(c *FlowchartConfigurationProperties) *FlowchartConfigurationProperties {
				return c.SetDiagramPadding(20)
			},
			property: flowchartPropertyDiagramPadding,
			value:    20,
		},
		{
			name: "Set HTML labels",
			setup: func(c *FlowchartConfigurationProperties) *FlowchartConfigurationProperties {
				return c.SetHtmlLabels(true)
			},
			property: flowchartPropertyHtmlLabels,
			value:    true,
		},
		{
			name: "Set node spacing",
			setup: func(c *FlowchartConfigurationProperties) *FlowchartConfigurationProperties {
				return c.SetNodeSpacing(50)
			},
			property: flowchartPropertyNodeSpacing,
			value:    50,
		},
		{
			name: "Set rank spacing",
			setup: func(c *FlowchartConfigurationProperties) *FlowchartConfigurationProperties {
				return c.SetRankSpacing(30)
			},
			property: flowchartPropertyRankSpacing,
			value:    30,
		},
		{
			name: "Set curve",
			setup: func(c *FlowchartConfigurationProperties) *FlowchartConfigurationProperties {
				return c.SetCurve("basis")
			},
			property: flowchartPropertyCurve,
			value:    "basis",
		},
		{
			name: "Set padding",
			setup: func(c *FlowchartConfigurationProperties) *FlowchartConfigurationProperties {
				return c.SetPadding(10)
			},
			property: flowchartPropertyPadding,
			value:    10,
		},
		{
			name: "Set default renderer",
			setup: func(c *FlowchartConfigurationProperties) *FlowchartConfigurationProperties {
				return c.SetDefaultRenderer("dagre")
			},
			property: flowchartPropertyDefaultRenderer,
			value:    "dagre",
		},
		{
			name: "Set wrapping width",
			setup: func(c *FlowchartConfigurationProperties) *FlowchartConfigurationProperties {
				return c.SetWrappingWidth(100)
			},
			property: flowchartPropertyWrappingWidth,
			value:    100,
		},
		{
			name: "Set arrow marker absolute",
			setup: func(c *FlowchartConfigurationProperties) *FlowchartConfigurationProperties {
				return c.SetArrowMarkerAbsolute(true)
			},
			property: flowchartPropertyArrowMarkerAbsolute,
			value:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := NewFlowchartConfigurationProperties()
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
