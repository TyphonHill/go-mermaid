package class

import (
	"strings"

	"github.com/TyphonHill/go-mermaid/diagrams/utils/basediagram"
)

const (
	baseClassConfigurationProperties string = basediagram.Indentation + "class:\n"
	classPropertyTitleTopMargin      string = "titleTopMargin"
	classPropertyArrowMarkerAbsolute string = "arrowMarkerAbsolute"
	classPropertyDividerMargin       string = "dividerMargin"
	classPropertyPadding             string = "padding"
	classPropertyTextHeight          string = "textHeight"
	classPropertyDefaultRenderer     string = "defaultRenderer"
	classPropertyNodeSpacing         string = "nodeSpacing"
	classPropertyRankSpacing         string = "rankSpacing"
	classPropertyDiagramPadding      string = "diagramPadding"
	classPropertyHtmlLabels          string = "htmlLabels"
	classPropertyHideEmptyMembersBox string = "hideEmptyMembersBox"
)

// ClassConfigurationProperties holds class-specific configuration
type ClassConfigurationProperties struct {
	basediagram.ConfigurationProperties
	properties map[string]basediagram.DiagramProperty
}

func NewClassConfigurationProperties() ClassConfigurationProperties {
	return ClassConfigurationProperties{
		ConfigurationProperties: basediagram.NewConfigurationProperties(),
		properties:              make(map[string]basediagram.DiagramProperty),
	}
}

// Setters for each property from the schema
func (c *ClassConfigurationProperties) SetTitleTopMargin(v int) *ClassConfigurationProperties {
	c.properties[classPropertyTitleTopMargin] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: classPropertyTitleTopMargin,
			Val:  v,
		},
	}
	return c
}

func (c *ClassConfigurationProperties) SetArrowMarkerAbsolute(v bool) *ClassConfigurationProperties {
	c.properties[classPropertyArrowMarkerAbsolute] = &basediagram.BoolProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: classPropertyArrowMarkerAbsolute,
			Val:  v,
		},
	}
	return c
}

func (c *ClassConfigurationProperties) SetDividerMargin(v int) *ClassConfigurationProperties {
	c.properties[classPropertyDividerMargin] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: classPropertyDividerMargin,
			Val:  v,
		},
	}
	return c
}

func (c *ClassConfigurationProperties) SetPadding(v int) *ClassConfigurationProperties {
	c.properties[classPropertyPadding] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: classPropertyPadding,
			Val:  v,
		},
	}
	return c
}

func (c *ClassConfigurationProperties) SetTextHeight(v int) *ClassConfigurationProperties {
	c.properties[classPropertyTextHeight] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: classPropertyTextHeight,
			Val:  v,
		},
	}
	return c
}

func (c *ClassConfigurationProperties) SetDefaultRenderer(v string) *ClassConfigurationProperties {
	c.properties[classPropertyDefaultRenderer] = &basediagram.StringProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: classPropertyDefaultRenderer,
			Val:  v,
		},
	}
	return c
}

func (c *ClassConfigurationProperties) SetNodeSpacing(v int) *ClassConfigurationProperties {
	c.properties[classPropertyNodeSpacing] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: classPropertyNodeSpacing,
			Val:  v,
		},
	}
	return c
}

func (c *ClassConfigurationProperties) SetRankSpacing(v int) *ClassConfigurationProperties {
	c.properties[classPropertyRankSpacing] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: classPropertyRankSpacing,
			Val:  v,
		},
	}
	return c
}

func (c *ClassConfigurationProperties) SetDiagramPadding(v int) *ClassConfigurationProperties {
	c.properties[classPropertyDiagramPadding] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: classPropertyDiagramPadding,
			Val:  v,
		},
	}
	return c
}

func (c *ClassConfigurationProperties) SetHtmlLabels(v bool) *ClassConfigurationProperties {
	c.properties[classPropertyHtmlLabels] = &basediagram.BoolProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: classPropertyHtmlLabels,
			Val:  v,
		},
	}
	return c
}

func (c *ClassConfigurationProperties) SetHideEmptyMembersBox(v bool) *ClassConfigurationProperties {
	c.properties[classPropertyHideEmptyMembersBox] = &basediagram.BoolProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: classPropertyHideEmptyMembersBox,
			Val:  v,
		},
	}
	return c
}

func (c ClassConfigurationProperties) String() string {
	var sb strings.Builder
	sb.WriteString(c.ConfigurationProperties.String())

	if len(c.properties) > 0 {
		sb.WriteString(baseClassConfigurationProperties)
		for _, prop := range c.properties {
			sb.WriteString(prop.Format())
		}
	}

	return sb.String()
}
