package state

import (
	"strings"

	"github.com/TyphonHill/go-mermaid/diagrams/utils/basediagram"
)

const (
	baseStateConfigurationProperties string = basediagram.Indentation + "state:\n"
	statePropertyTitleTopMargin      string = "titleTopMargin"
	statePropertyArrowMarkerAbsolute string = "arrowMarkerAbsolute"
	statePropertyDividerMargin       string = "dividerMargin"
	statePropertySizeUnit            string = "sizeUnit"
	statePropertyPadding             string = "padding"
	statePropertyTextHeight          string = "textHeight"
	statePropertyTitleShift          string = "titleShift"
	statePropertyNoteMargin          string = "noteMargin"
	statePropertyNodeSpacing         string = "nodeSpacing"
	statePropertyRankSpacing         string = "rankSpacing"
	statePropertyForkWidth           string = "forkWidth"
	statePropertyForkHeight          string = "forkHeight"
	statePropertyMiniPadding         string = "miniPadding"
	statePropertyFontSizeFactor      string = "fontSizeFactor"
	statePropertyFontSize            string = "fontSize"
	statePropertyLabelHeight         string = "labelHeight"
	statePropertyEdgeLengthFactor    string = "edgeLengthFactor"
	statePropertyCompositTitleSize   string = "compositTitleSize"
	statePropertyRadius              string = "radius"
	statePropertyDefaultRenderer     string = "defaultRenderer"
)

// StateConfigurationProperties holds state-specific configuration
type StateConfigurationProperties struct {
	basediagram.ConfigurationProperties
	properties map[string]basediagram.DiagramProperty
}

func NewStateConfigurationProperties() StateConfigurationProperties {
	return StateConfigurationProperties{
		ConfigurationProperties: basediagram.NewConfigurationProperties(),
		properties:              make(map[string]basediagram.DiagramProperty),
	}
}

func (c *StateConfigurationProperties) SetTitleTopMargin(v int) *StateConfigurationProperties {
	c.properties[statePropertyTitleTopMargin] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: statePropertyTitleTopMargin,
			Val:  v,
		},
	}
	return c
}

func (c *StateConfigurationProperties) SetArrowMarkerAbsolute(v bool) *StateConfigurationProperties {
	c.properties[statePropertyArrowMarkerAbsolute] = &basediagram.BoolProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: statePropertyArrowMarkerAbsolute,
			Val:  v,
		},
	}
	return c
}

func (c *StateConfigurationProperties) SetDividerMargin(v int) *StateConfigurationProperties {
	c.properties[statePropertyDividerMargin] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: statePropertyDividerMargin,
			Val:  v,
		},
	}
	return c
}

func (c *StateConfigurationProperties) SetSizeUnit(v int) *StateConfigurationProperties {
	c.properties[statePropertySizeUnit] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: statePropertySizeUnit,
			Val:  v,
		},
	}
	return c
}

func (c *StateConfigurationProperties) SetPadding(v int) *StateConfigurationProperties {
	c.properties[statePropertyPadding] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: statePropertyPadding,
			Val:  v,
		},
	}
	return c
}

func (c *StateConfigurationProperties) SetTextHeight(v int) *StateConfigurationProperties {
	c.properties[statePropertyTextHeight] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: statePropertyTextHeight,
			Val:  v,
		},
	}
	return c
}

func (c *StateConfigurationProperties) SetTitleShift(v int) *StateConfigurationProperties {
	c.properties[statePropertyTitleShift] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: statePropertyTitleShift,
			Val:  v,
		},
	}
	return c
}

func (c *StateConfigurationProperties) SetNoteMargin(v int) *StateConfigurationProperties {
	c.properties[statePropertyNoteMargin] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: statePropertyNoteMargin,
			Val:  v,
		},
	}
	return c
}

func (c *StateConfigurationProperties) SetNodeSpacing(v int) *StateConfigurationProperties {
	c.properties[statePropertyNodeSpacing] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: statePropertyNodeSpacing,
			Val:  v,
		},
	}
	return c
}

func (c *StateConfigurationProperties) SetRankSpacing(v int) *StateConfigurationProperties {
	c.properties[statePropertyRankSpacing] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: statePropertyRankSpacing,
			Val:  v,
		},
	}
	return c
}

func (c *StateConfigurationProperties) SetForkWidth(v int) *StateConfigurationProperties {
	c.properties[statePropertyForkWidth] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: statePropertyForkWidth,
			Val:  v,
		},
	}
	return c
}

func (c *StateConfigurationProperties) SetForkHeight(v int) *StateConfigurationProperties {
	c.properties[statePropertyForkHeight] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: statePropertyForkHeight,
			Val:  v,
		},
	}
	return c
}

func (c *StateConfigurationProperties) SetMiniPadding(v int) *StateConfigurationProperties {
	c.properties[statePropertyMiniPadding] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: statePropertyMiniPadding,
			Val:  v,
		},
	}
	return c
}

func (c *StateConfigurationProperties) SetFontSizeFactor(v int) *StateConfigurationProperties {
	c.properties[statePropertyFontSizeFactor] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: statePropertyFontSizeFactor,
			Val:  v,
		},
	}
	return c
}

func (c *StateConfigurationProperties) SetFontSize(v int) *StateConfigurationProperties {
	c.properties[statePropertyFontSize] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: statePropertyFontSize,
			Val:  v,
		},
	}
	return c
}

func (c *StateConfigurationProperties) SetLabelHeight(v int) *StateConfigurationProperties {
	c.properties[statePropertyLabelHeight] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: statePropertyLabelHeight,
			Val:  v,
		},
	}
	return c
}

func (c *StateConfigurationProperties) SetEdgeLengthFactor(v string) *StateConfigurationProperties {
	c.properties[statePropertyEdgeLengthFactor] = &basediagram.StringProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: statePropertyEdgeLengthFactor,
			Val:  v,
		},
	}
	return c
}

func (c *StateConfigurationProperties) SetCompositTitleSize(v int) *StateConfigurationProperties {
	c.properties[statePropertyCompositTitleSize] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: statePropertyCompositTitleSize,
			Val:  v,
		},
	}
	return c
}

func (c *StateConfigurationProperties) SetRadius(v int) *StateConfigurationProperties {
	c.properties[statePropertyRadius] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: statePropertyRadius,
			Val:  v,
		},
	}
	return c
}

func (c *StateConfigurationProperties) SetDefaultRenderer(v string) *StateConfigurationProperties {
	c.properties[statePropertyDefaultRenderer] = &basediagram.StringProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: statePropertyDefaultRenderer,
			Val:  v,
		},
	}
	return c
}

func (c StateConfigurationProperties) String() string {
	var sb strings.Builder
	sb.WriteString(c.ConfigurationProperties.String())

	if len(c.properties) > 0 {
		sb.WriteString(baseStateConfigurationProperties)
		for _, prop := range c.properties {
			sb.WriteString(prop.Format())
		}
	}

	return sb.String()
}
