package sequence

import (
	"strings"
	"testing"

	"github.com/TyphonHill/go-mermaid/diagrams/utils/basediagram"
)

func TestNewSequenceConfigurationProperties(t *testing.T) {
	got := NewSequenceConfigurationProperties()

	if got.properties == nil {
		t.Error("NewSequenceConfigurationProperties() properties map is nil")
	}

	if len(got.properties) != 0 {
		t.Errorf("NewSequenceConfigurationProperties() properties map length = %v, want 0", len(got.properties))
	}
}

func TestSequenceConfigurationProperties_String(t *testing.T) {
	tests := []struct {
		name     string
		config   SequenceConfigurationProperties
		setup    func(*SequenceConfigurationProperties)
		contains []string
	}{
		{
			name:   "Empty configuration",
			config: NewSequenceConfigurationProperties(),
			contains: []string{
				"",
			},
		},
		{
			name:   "Configuration with single property",
			config: NewSequenceConfigurationProperties(),
			setup: func(c *SequenceConfigurationProperties) {
				c.SetActivationWidth(10)
			},
			contains: []string{
				"sequence:",
				"activationWidth: 10",
			},
		},
		{
			name:   "Configuration with multiple properties",
			config: NewSequenceConfigurationProperties(),
			setup: func(c *SequenceConfigurationProperties) {
				c.SetActivationWidth(10)
				c.SetDiagramMarginX(20)
				c.SetHideUnusedParticipants(true)
			},
			contains: []string{
				"sequence:",
				"activationWidth: 10",
				"diagramMarginX: 20",
				"hideUnusedParticipants: true",
			},
		},
		{
			name:   "Configuration with base properties",
			config: NewSequenceConfigurationProperties(),
			setup: func(c *SequenceConfigurationProperties) {
				c.ConfigurationProperties = basediagram.NewConfigurationProperties()
				c.ConfigurationProperties.SetFontSize(12)
				c.SetActivationWidth(10)
			},
			contains: []string{
				"fontSize: 12",
				"sequence:",
				"activationWidth: 10",
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

func TestSequenceConfigurationProperties_Setters(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(*SequenceConfigurationProperties) *SequenceConfigurationProperties
		property string
		value    interface{}
	}{
		{
			name: "Set arrow marker absolute",
			setup: func(c *SequenceConfigurationProperties) *SequenceConfigurationProperties {
				return c.SetArrowMarkerAbsolute(true)
			},
			property: sequencePropertyArrowMarkerAbsolute,
			value:    true,
		},
		{
			name: "Set hide unused participants",
			setup: func(c *SequenceConfigurationProperties) *SequenceConfigurationProperties {
				return c.SetHideUnusedParticipants(true)
			},
			property: sequencePropertyHideUnusedParticipants,
			value:    true,
		},
		{
			name: "Set activation width",
			setup: func(c *SequenceConfigurationProperties) *SequenceConfigurationProperties {
				return c.SetActivationWidth(10)
			},
			property: sequencePropertyActivationWidth,
			value:    10,
		},
		{
			name: "Set message align",
			setup: func(c *SequenceConfigurationProperties) *SequenceConfigurationProperties {
				return c.SetMessageAlign("left")
			},
			property: sequencePropertyMessageAlign,
			value:    "left",
		},
		{
			name: "Set actor font family",
			setup: func(c *SequenceConfigurationProperties) *SequenceConfigurationProperties {
				return c.SetActorFontFamily("Arial")
			},
			property: sequencePropertyActorFontFamily,
			value:    "Arial",
		},
		{
			name: "Set diagram margin X",
			setup: func(c *SequenceConfigurationProperties) *SequenceConfigurationProperties {
				return c.SetDiagramMarginX(20)
			},
			property: sequencePropertyDiagramMarginX,
			value:    20,
		},
		{
			name: "Set diagram margin Y",
			setup: func(c *SequenceConfigurationProperties) *SequenceConfigurationProperties {
				return c.SetDiagramMarginY(30)
			},
			property: sequencePropertyDiagramMarginY,
			value:    30,
		},
		{
			name: "Set actor margin",
			setup: func(c *SequenceConfigurationProperties) *SequenceConfigurationProperties {
				return c.SetActorMargin(15)
			},
			property: sequencePropertyActorMargin,
			value:    15,
		},
		{
			name: "Set width",
			setup: func(c *SequenceConfigurationProperties) *SequenceConfigurationProperties {
				return c.SetWidth(800)
			},
			property: sequencePropertyWidth,
			value:    800,
		},
		{
			name: "Set height",
			setup: func(c *SequenceConfigurationProperties) *SequenceConfigurationProperties {
				return c.SetHeight(600)
			},
			property: sequencePropertyHeight,
			value:    600,
		},
		{
			name: "Set box margin",
			setup: func(c *SequenceConfigurationProperties) *SequenceConfigurationProperties {
				return c.SetBoxMargin(10)
			},
			property: sequencePropertyBoxMargin,
			value:    10,
		},
		{
			name: "Set box text margin",
			setup: func(c *SequenceConfigurationProperties) *SequenceConfigurationProperties {
				return c.SetBoxTextMargin(5)
			},
			property: sequencePropertyBoxTextMargin,
			value:    5,
		},
		{
			name: "Set note margin",
			setup: func(c *SequenceConfigurationProperties) *SequenceConfigurationProperties {
				return c.SetNoteMargin(8)
			},
			property: sequencePropertyNoteMargin,
			value:    8,
		},
		{
			name: "Set message margin",
			setup: func(c *SequenceConfigurationProperties) *SequenceConfigurationProperties {
				return c.SetMessageMargin(12)
			},
			property: sequencePropertyMessageMargin,
			value:    12,
		},
		{
			name: "Set mirror actors",
			setup: func(c *SequenceConfigurationProperties) *SequenceConfigurationProperties {
				return c.SetMirrorActors(true)
			},
			property: sequencePropertyMirrorActors,
			value:    true,
		},
		{
			name: "Set force menus",
			setup: func(c *SequenceConfigurationProperties) *SequenceConfigurationProperties {
				return c.SetForceMenus(true)
			},
			property: sequencePropertyForceMenus,
			value:    true,
		},
		{
			name: "Set bottom margin adj",
			setup: func(c *SequenceConfigurationProperties) *SequenceConfigurationProperties {
				return c.SetBottomMarginAdj(25)
			},
			property: sequencePropertyBottomMarginAdj,
			value:    25,
		},
		{
			name: "Set right angles",
			setup: func(c *SequenceConfigurationProperties) *SequenceConfigurationProperties {
				return c.SetRightAngles(true)
			},
			property: sequencePropertyRightAngles,
			value:    true,
		},
		{
			name: "Set show sequence numbers",
			setup: func(c *SequenceConfigurationProperties) *SequenceConfigurationProperties {
				return c.SetShowSequenceNumbers(true)
			},
			property: sequencePropertyShowSequenceNumbers,
			value:    true,
		},
		{
			name: "Set actor font size",
			setup: func(c *SequenceConfigurationProperties) *SequenceConfigurationProperties {
				return c.SetActorFontSize(14)
			},
			property: sequencePropertyActorFontSize,
			value:    14,
		},
		{
			name: "Set actor font weight",
			setup: func(c *SequenceConfigurationProperties) *SequenceConfigurationProperties {
				return c.SetActorFontWeight(600)
			},
			property: sequencePropertyActorFontWeight,
			value:    600,
		},
		{
			name: "Set note font size",
			setup: func(c *SequenceConfigurationProperties) *SequenceConfigurationProperties {
				return c.SetNoteFontSize(12)
			},
			property: sequencePropertyNoteFontSize,
			value:    12,
		},
		{
			name: "Set note font family",
			setup: func(c *SequenceConfigurationProperties) *SequenceConfigurationProperties {
				return c.SetNoteFontFamily("Helvetica")
			},
			property: sequencePropertyNoteFontFamily,
			value:    "Helvetica",
		},
		{
			name: "Set note font weight",
			setup: func(c *SequenceConfigurationProperties) *SequenceConfigurationProperties {
				return c.SetNoteFontWeight(400)
			},
			property: sequencePropertyNoteFontWeight,
			value:    400,
		},
		{
			name: "Set note align",
			setup: func(c *SequenceConfigurationProperties) *SequenceConfigurationProperties {
				return c.SetNoteAlign("right")
			},
			property: sequencePropertyNoteAlign,
			value:    "right",
		},
		{
			name: "Set message font size",
			setup: func(c *SequenceConfigurationProperties) *SequenceConfigurationProperties {
				return c.SetMessageFontSize(13)
			},
			property: sequencePropertyMessageFontSize,
			value:    13,
		},
		{
			name: "Set message font family",
			setup: func(c *SequenceConfigurationProperties) *SequenceConfigurationProperties {
				return c.SetMessageFontFamily("Courier")
			},
			property: sequencePropertyMessageFontFamily,
			value:    "Courier",
		},
		{
			name: "Set message font weight",
			setup: func(c *SequenceConfigurationProperties) *SequenceConfigurationProperties {
				return c.SetMessageFontWeight(500)
			},
			property: sequencePropertyMessageFontWeight,
			value:    500,
		},
		{
			name: "Set wrap",
			setup: func(c *SequenceConfigurationProperties) *SequenceConfigurationProperties {
				return c.SetWrap(true)
			},
			property: sequencePropertyWrap,
			value:    true,
		},
		{
			name: "Set wrap padding",
			setup: func(c *SequenceConfigurationProperties) *SequenceConfigurationProperties {
				return c.SetWrapPadding(15)
			},
			property: sequencePropertyWrapPadding,
			value:    15,
		},
		{
			name: "Set label box width",
			setup: func(c *SequenceConfigurationProperties) *SequenceConfigurationProperties {
				return c.SetLabelBoxWidth(50)
			},
			property: sequencePropertyLabelBoxWidth,
			value:    50,
		},
		{
			name: "Set label box height",
			setup: func(c *SequenceConfigurationProperties) *SequenceConfigurationProperties {
				return c.SetLabelBoxHeight(30)
			},
			property: sequencePropertyLabelBoxHeight,
			value:    30,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := NewSequenceConfigurationProperties()
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
