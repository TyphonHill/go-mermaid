package timeline

import (
	"strings"

	"github.com/TyphonHill/go-mermaid/diagrams/utils/basediagram"
)

const (
	baseTimelineConfigurationProperties string = basediagram.Indentation + "timeline\n"
	timelinePropertyDisableMulticolor   string = "disableMulticolor"
	timelinePropertyDiagramMarginX      string = "diagramMarginX"
	timelinePropertyDiagramMarginY      string = "diagramMarginY"
	timelinePropertyLeftMargin          string = "leftMargin"
	timelinePropertyWidth               string = "width"
	timelinePropertyHeight              string = "height"
	timelinePropertyPadding             string = "padding"
	timelinePropertyBoxMargin           string = "boxMargin"
	timelinePropertyBoxTextMargin       string = "boxTextMargin"
	timelinePropertyNoteMargin          string = "noteMargin"
	timelinePropertyMessageMargin       string = "messageMargin"
	timelinePropertyMessageAlign        string = "messageAlign"
	timelinePropertyBottomMarginAdj     string = "bottomMarginAdj"
	timelinePropertyRightAngles         string = "rightAngles"
	timelinePropertyTaskFontSize        string = "taskFontSize"
	timelinePropertyTaskFontFamily      string = "taskFontFamily"
	timelinePropertyTaskMargin          string = "taskMargin"
	timelinePropertyActivationWidth     string = "activationWidth"
	timelinePropertyTextPlacement       string = "textPlacement"
)

// TimelineConfigurationProperties holds timeline-specific configuration
type TimelineConfigurationProperties struct {
	basediagram.ConfigurationProperties
	properties map[string]basediagram.DiagramProperty
}

func NewTimeLineConfigurationProperties() TimelineConfigurationProperties {
	return TimelineConfigurationProperties{
		ConfigurationProperties: basediagram.NewConfigurationProperties(),
		properties:              make(map[string]basediagram.DiagramProperty),
	}
}

func (c *TimelineConfigurationProperties) SetDisableMulticolor(v bool) *TimelineConfigurationProperties {
	c.properties[timelinePropertyDisableMulticolor] = &basediagram.BoolProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: timelinePropertyDisableMulticolor,
			Val:  v,
		},
	}
	return c
}

func (c *TimelineConfigurationProperties) SetDiagramMarginX(v int) *TimelineConfigurationProperties {
	c.properties[timelinePropertyDiagramMarginX] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: timelinePropertyDiagramMarginX,
			Val:  v,
		},
	}
	return c
}

func (c *TimelineConfigurationProperties) SetDiagramMarginY(v int) *TimelineConfigurationProperties {
	c.properties[timelinePropertyDiagramMarginY] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: timelinePropertyDiagramMarginY,
			Val:  v,
		},
	}
	return c
}

func (c *TimelineConfigurationProperties) SetLeftMargin(v int) *TimelineConfigurationProperties {
	c.properties[timelinePropertyLeftMargin] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: timelinePropertyLeftMargin,
			Val:  v,
		},
	}
	return c
}

func (c *TimelineConfigurationProperties) SetWidth(v int) *TimelineConfigurationProperties {
	c.properties[timelinePropertyWidth] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: timelinePropertyWidth,
			Val:  v,
		},
	}
	return c
}

func (c *TimelineConfigurationProperties) SetHeight(v int) *TimelineConfigurationProperties {
	c.properties[timelinePropertyHeight] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: timelinePropertyHeight,
			Val:  v,
		},
	}
	return c
}

func (c *TimelineConfigurationProperties) SetPadding(v float64) *TimelineConfigurationProperties {
	c.properties[timelinePropertyPadding] = &basediagram.FloatProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: timelinePropertyPadding,
			Val:  v,
		},
	}
	return c
}

func (c *TimelineConfigurationProperties) SetBoxMargin(v int) *TimelineConfigurationProperties {
	c.properties[timelinePropertyBoxMargin] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: timelinePropertyBoxMargin,
			Val:  v,
		},
	}
	return c
}

func (c *TimelineConfigurationProperties) SetBoxTextMargin(v int) *TimelineConfigurationProperties {
	c.properties[timelinePropertyBoxTextMargin] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: timelinePropertyBoxTextMargin,
			Val:  v,
		},
	}
	return c
}

func (c *TimelineConfigurationProperties) SetNoteMargin(v int) *TimelineConfigurationProperties {
	c.properties[timelinePropertyNoteMargin] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: timelinePropertyNoteMargin,
			Val:  v,
		},
	}
	return c
}

func (c *TimelineConfigurationProperties) SetMessageMargin(v int) *TimelineConfigurationProperties {
	c.properties[timelinePropertyMessageMargin] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: timelinePropertyMessageMargin,
			Val:  v,
		},
	}
	return c
}

func (c *TimelineConfigurationProperties) SetMessageAlign(v string) *TimelineConfigurationProperties {
	c.properties[timelinePropertyMessageAlign] = &basediagram.StringProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: timelinePropertyMessageAlign,
			Val:  v,
		},
	}
	return c
}

func (c *TimelineConfigurationProperties) SetBottomMarginAdj(v int) *TimelineConfigurationProperties {
	c.properties[timelinePropertyBottomMarginAdj] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: timelinePropertyBottomMarginAdj,
			Val:  v,
		},
	}
	return c
}

func (c *TimelineConfigurationProperties) SetRightAngles(v bool) *TimelineConfigurationProperties {
	c.properties[timelinePropertyRightAngles] = &basediagram.BoolProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: timelinePropertyRightAngles,
			Val:  v,
		},
	}
	return c
}

func (c *TimelineConfigurationProperties) SetTaskFontSize(v int) *TimelineConfigurationProperties {
	c.properties[timelinePropertyTaskFontSize] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: timelinePropertyTaskFontSize,
			Val:  v,
		},
	}
	return c
}

func (c *TimelineConfigurationProperties) SetTaskFontFamily(v string) *TimelineConfigurationProperties {
	c.properties[timelinePropertyTaskFontFamily] = &basediagram.StringProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: timelinePropertyTaskFontFamily,
			Val:  v,
		},
	}
	return c
}

func (c *TimelineConfigurationProperties) SetTaskMargin(v float64) *TimelineConfigurationProperties {
	c.properties[timelinePropertyTaskMargin] = &basediagram.FloatProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: timelinePropertyTaskMargin,
			Val:  v,
		},
	}
	return c
}

func (c *TimelineConfigurationProperties) SetActivationWidth(v float64) *TimelineConfigurationProperties {
	c.properties[timelinePropertyActivationWidth] = &basediagram.FloatProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: timelinePropertyActivationWidth,
			Val:  v,
		},
	}
	return c
}

func (c *TimelineConfigurationProperties) SetTextPlacement(v string) *TimelineConfigurationProperties {
	c.properties[timelinePropertyTextPlacement] = &basediagram.StringProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: timelinePropertyTextPlacement,
			Val:  v,
		},
	}
	return c
}

func (c TimelineConfigurationProperties) String() string {
	var sb strings.Builder
	sb.WriteString(c.ConfigurationProperties.String())

	if len(c.properties) > 0 {
		sb.WriteString(baseTimelineConfigurationProperties)
		for _, prop := range c.properties {
			sb.WriteString(prop.Format())
		}
	}

	return sb.String()
}
