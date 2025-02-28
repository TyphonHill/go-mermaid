package timeline

import (
	"strings"
	"testing"

	"github.com/TyphonHill/go-mermaid/diagrams/utils/basediagram"
)

func TestNewTimeLineConfigurationProperties(t *testing.T) {
	got := NewTimeLineConfigurationProperties()

	if got.properties == nil {
		t.Error("NewTimeLineConfigurationProperties() properties map is nil")
	}

	if len(got.properties) != 0 {
		t.Errorf("NewTimeLineConfigurationProperties() properties map length = %v, want 0", len(got.properties))
	}
}

func TestTimelineConfigurationProperties_String(t *testing.T) {
	tests := []struct {
		name     string
		config   TimelineConfigurationProperties
		setup    func(*TimelineConfigurationProperties)
		contains []string
	}{
		{
			name:   "Empty configuration",
			config: NewTimeLineConfigurationProperties(),
			contains: []string{
				"",
			},
		},
		{
			name:   "Configuration with single property",
			config: NewTimeLineConfigurationProperties(),
			setup: func(c *TimelineConfigurationProperties) {
				c.SetPadding(10.5)
			},
			contains: []string{
				"timeline",
				"padding: 10.5",
			},
		},
		{
			name:   "Configuration with multiple properties",
			config: NewTimeLineConfigurationProperties(),
			setup: func(c *TimelineConfigurationProperties) {
				c.SetPadding(10.5)
				c.SetDisableMulticolor(true)
				c.SetTaskFontFamily("Arial")
			},
			contains: []string{
				"timeline",
				"padding: 10.5",
				"disableMulticolor: true",
				"taskFontFamily: Arial",
			},
		},
		{
			name:   "Configuration with base properties",
			config: NewTimeLineConfigurationProperties(),
			setup: func(c *TimelineConfigurationProperties) {
				c.ConfigurationProperties = basediagram.NewConfigurationProperties()
				c.ConfigurationProperties.SetFontSize(12)
				c.SetPadding(10.5)
			},
			contains: []string{
				"fontSize: 12",
				"timeline",
				"padding: 10.5",
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

func TestTimelineConfigurationProperties_Setters(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(*TimelineConfigurationProperties) *TimelineConfigurationProperties
		property string
		value    interface{}
	}{
		{
			name: "Set disable multicolor",
			setup: func(c *TimelineConfigurationProperties) *TimelineConfigurationProperties {
				return c.SetDisableMulticolor(true)
			},
			property: timelinePropertyDisableMulticolor,
			value:    true,
		},
		{
			name: "Set diagram margin X",
			setup: func(c *TimelineConfigurationProperties) *TimelineConfigurationProperties {
				return c.SetDiagramMarginX(20)
			},
			property: timelinePropertyDiagramMarginX,
			value:    20,
		},
		{
			name: "Set diagram margin Y",
			setup: func(c *TimelineConfigurationProperties) *TimelineConfigurationProperties {
				return c.SetDiagramMarginY(30)
			},
			property: timelinePropertyDiagramMarginY,
			value:    30,
		},
		{
			name: "Set left margin",
			setup: func(c *TimelineConfigurationProperties) *TimelineConfigurationProperties {
				return c.SetLeftMargin(15)
			},
			property: timelinePropertyLeftMargin,
			value:    15,
		},
		{
			name: "Set width",
			setup: func(c *TimelineConfigurationProperties) *TimelineConfigurationProperties {
				return c.SetWidth(800)
			},
			property: timelinePropertyWidth,
			value:    800,
		},
		{
			name: "Set height",
			setup: func(c *TimelineConfigurationProperties) *TimelineConfigurationProperties {
				return c.SetHeight(600)
			},
			property: timelinePropertyHeight,
			value:    600,
		},
		{
			name: "Set padding",
			setup: func(c *TimelineConfigurationProperties) *TimelineConfigurationProperties {
				return c.SetPadding(10.5)
			},
			property: timelinePropertyPadding,
			value:    10.5,
		},
		{
			name: "Set box margin",
			setup: func(c *TimelineConfigurationProperties) *TimelineConfigurationProperties {
				return c.SetBoxMargin(10)
			},
			property: timelinePropertyBoxMargin,
			value:    10,
		},
		{
			name: "Set box text margin",
			setup: func(c *TimelineConfigurationProperties) *TimelineConfigurationProperties {
				return c.SetBoxTextMargin(5)
			},
			property: timelinePropertyBoxTextMargin,
			value:    5,
		},
		{
			name: "Set note margin",
			setup: func(c *TimelineConfigurationProperties) *TimelineConfigurationProperties {
				return c.SetNoteMargin(8)
			},
			property: timelinePropertyNoteMargin,
			value:    8,
		},
		{
			name: "Set message margin",
			setup: func(c *TimelineConfigurationProperties) *TimelineConfigurationProperties {
				return c.SetMessageMargin(12)
			},
			property: timelinePropertyMessageMargin,
			value:    12,
		},
		{
			name: "Set message align",
			setup: func(c *TimelineConfigurationProperties) *TimelineConfigurationProperties {
				return c.SetMessageAlign("left")
			},
			property: timelinePropertyMessageAlign,
			value:    "left",
		},
		{
			name: "Set bottom margin adj",
			setup: func(c *TimelineConfigurationProperties) *TimelineConfigurationProperties {
				return c.SetBottomMarginAdj(25)
			},
			property: timelinePropertyBottomMarginAdj,
			value:    25,
		},
		{
			name: "Set right angles",
			setup: func(c *TimelineConfigurationProperties) *TimelineConfigurationProperties {
				return c.SetRightAngles(true)
			},
			property: timelinePropertyRightAngles,
			value:    true,
		},
		{
			name: "Set task font size",
			setup: func(c *TimelineConfigurationProperties) *TimelineConfigurationProperties {
				return c.SetTaskFontSize(14)
			},
			property: timelinePropertyTaskFontSize,
			value:    14,
		},
		{
			name: "Set task font family",
			setup: func(c *TimelineConfigurationProperties) *TimelineConfigurationProperties {
				return c.SetTaskFontFamily("Arial")
			},
			property: timelinePropertyTaskFontFamily,
			value:    "Arial",
		},
		{
			name: "Set task margin",
			setup: func(c *TimelineConfigurationProperties) *TimelineConfigurationProperties {
				return c.SetTaskMargin(5.5)
			},
			property: timelinePropertyTaskMargin,
			value:    5.5,
		},
		{
			name: "Set activation width",
			setup: func(c *TimelineConfigurationProperties) *TimelineConfigurationProperties {
				return c.SetActivationWidth(2.5)
			},
			property: timelinePropertyActivationWidth,
			value:    2.5,
		},
		{
			name: "Set text placement",
			setup: func(c *TimelineConfigurationProperties) *TimelineConfigurationProperties {
				return c.SetTextPlacement("top")
			},
			property: timelinePropertyTextPlacement,
			value:    "top",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := NewTimeLineConfigurationProperties()
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
			case *basediagram.FloatProperty:
				got = p.Val
			}

			if got != tt.value {
				t.Errorf("Property %q = %v, want %v", tt.property, got, tt.value)
			}
		})
	}
}
