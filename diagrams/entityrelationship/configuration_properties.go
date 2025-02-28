package entityrelationship

import (
	"strings"

	"github.com/TyphonHill/go-mermaid/diagrams/utils/basediagram"
)

const (
	baseErConfigurationProperties string = basediagram.Indentation + "er:\n"
	erPropertyTitleTopMargin      string = "titleTopMargin"
	erPropertyDiagramPadding      string = "diagramPadding"
	erPropertyLayoutDirection     string = "layoutDirection"
	erPropertyMinEntityWidth      string = "minEntityWidth"
	erPropertyMinEntityHeight     string = "minEntityHeight"
	erPropertyEntityPadding       string = "entityPadding"
	erPropertyStroke              string = "stroke"
	erPropertyFill                string = "fill"
	erPropertyFontSize            string = "fontSize"
)

type ErConfigurationProperties struct {
	basediagram.ConfigurationProperties
	properties map[string]basediagram.DiagramProperty
}

func NewErConfigurationProperties() ErConfigurationProperties {
	return ErConfigurationProperties{
		ConfigurationProperties: basediagram.NewConfigurationProperties(),
		properties:              make(map[string]basediagram.DiagramProperty),
	}
}

// Setters for each property from the schema
func (c *ErConfigurationProperties) SetTitleTopMargin(v int) *ErConfigurationProperties {
	c.properties[erPropertyTitleTopMargin] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: erPropertyTitleTopMargin,
			Val:  v,
		},
	}
	return c
}

func (c *ErConfigurationProperties) SetDiagramPadding(v int) *ErConfigurationProperties {
	c.properties[erPropertyDiagramPadding] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: erPropertyDiagramPadding,
			Val:  v,
		},
	}
	return c
}

func (c *ErConfigurationProperties) SetLayoutDirection(v string) *ErConfigurationProperties {
	c.properties[erPropertyLayoutDirection] = &basediagram.StringProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: erPropertyLayoutDirection,
			Val:  v,
		},
	}
	return c
}

func (c *ErConfigurationProperties) SetMinEntityWidth(v int) *ErConfigurationProperties {
	c.properties[erPropertyMinEntityWidth] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: erPropertyMinEntityWidth,
			Val:  v,
		},
	}
	return c
}

func (c *ErConfigurationProperties) SetMinEntityHeight(v int) *ErConfigurationProperties {
	c.properties[erPropertyMinEntityHeight] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: erPropertyMinEntityHeight,
			Val:  v,
		},
	}
	return c
}

func (c *ErConfigurationProperties) SetEntityPadding(v int) *ErConfigurationProperties {
	c.properties[erPropertyEntityPadding] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: erPropertyEntityPadding,
			Val:  v,
		},
	}
	return c
}

func (c *ErConfigurationProperties) SetStroke(v string) *ErConfigurationProperties {
	c.properties[erPropertyStroke] = &basediagram.StringProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: erPropertyStroke,
			Val:  v,
		},
	}
	return c
}

func (c *ErConfigurationProperties) SetFill(v string) *ErConfigurationProperties {
	c.properties[erPropertyFill] = &basediagram.StringProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: erPropertyFill,
			Val:  v,
		},
	}
	return c
}

func (c *ErConfigurationProperties) SetFontSize(v int) *ErConfigurationProperties {
	c.properties[erPropertyFontSize] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: erPropertyFontSize,
			Val:  v,
		},
	}
	return c
}

func (c ErConfigurationProperties) String() string {
	var sb strings.Builder
	sb.WriteString(c.ConfigurationProperties.String())

	if len(c.properties) > 0 {
		sb.WriteString(baseErConfigurationProperties)
		for _, prop := range c.properties {
			sb.WriteString(prop.Format())
		}
	}

	return sb.String()
}
