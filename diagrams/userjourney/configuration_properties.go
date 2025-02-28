package userjourney

import (
	"strings"

	"github.com/TyphonHill/go-mermaid/diagrams/utils/basediagram"
)

const (
	baseJourneyConfigurationProperties string = basediagram.Indentation + "journey:\n"

	journeyPropertyDiagramMarginX  string = "diagramMarginX"
	journeyPropertyDiagramMarginY  string = "diagramMarginY"
	journeyPropertyLeftMargin      string = "leftMargin"
	journeyPropertyWidth           string = "width"
	journeyPropertyHeight          string = "height"
	journeyPropertyBoxMargin       string = "boxMargin"
	journeyPropertyBoxTextMargin   string = "boxTextMargin"
	journeyPropertyNoteMargin      string = "noteMargin"
	journeyPropertyMessageMargin   string = "messageMargin"
	journeyPropertyMessageAlign    string = "messageAlign"
	journeyPropertyBottomMarginAdj string = "bottomMarginAdj"
	journeyPropertyRightAngles     string = "rightAngles"
	journeyPropertyTaskFontSize    string = "taskFontSize"
	journeyPropertyTaskFontFamily  string = "taskFontFamily"
	journeyPropertyTaskMargin      string = "taskMargin"
	journeyPropertyActivationWidth string = "activationWidth"
	journeyPropertyTextPlacement   string = "textPlacement"
	journeyPropertyActorColours    string = "actorColours"
	journeyPropertySectionFills    string = "sectionFills"
	journeyPropertySectionColours  string = "sectionColours"
)

// JourneyConfigurationProperties holds journey-specific configuration
type JourneyConfigurationProperties struct {
	basediagram.ConfigurationProperties
	properties map[string]basediagram.DiagramProperty
}

func NewJourneyConfigurationProperties() JourneyConfigurationProperties {
	return JourneyConfigurationProperties{
		ConfigurationProperties: basediagram.NewConfigurationProperties(),
		properties:              make(map[string]basediagram.DiagramProperty),
	}
}

func (c *JourneyConfigurationProperties) SetDiagramMarginX(v int) *JourneyConfigurationProperties {
	c.properties[journeyPropertyDiagramMarginX] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: journeyPropertyDiagramMarginX,
			Val:  v,
		},
	}
	return c
}

func (c *JourneyConfigurationProperties) SetDiagramMarginY(v int) *JourneyConfigurationProperties {
	c.properties[journeyPropertyDiagramMarginY] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: journeyPropertyDiagramMarginY,
			Val:  v,
		},
	}
	return c
}

func (c *JourneyConfigurationProperties) SetLeftMargin(v int) *JourneyConfigurationProperties {
	c.properties[journeyPropertyLeftMargin] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: journeyPropertyLeftMargin,
			Val:  v,
		},
	}
	return c
}

func (c *JourneyConfigurationProperties) SetWidth(v int) *JourneyConfigurationProperties {
	c.properties[journeyPropertyWidth] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: journeyPropertyWidth,
			Val:  v,
		},
	}
	return c
}

func (c *JourneyConfigurationProperties) SetHeight(v int) *JourneyConfigurationProperties {
	c.properties[journeyPropertyHeight] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: journeyPropertyHeight,
			Val:  v,
		},
	}
	return c
}

func (c *JourneyConfigurationProperties) SetBoxMargin(v int) *JourneyConfigurationProperties {
	c.properties[journeyPropertyBoxMargin] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: journeyPropertyBoxMargin,
			Val:  v,
		},
	}
	return c
}

func (c *JourneyConfigurationProperties) SetBoxTextMargin(v int) *JourneyConfigurationProperties {
	c.properties[journeyPropertyBoxTextMargin] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: journeyPropertyBoxTextMargin,
			Val:  v,
		},
	}
	return c
}

func (c *JourneyConfigurationProperties) SetNoteMargin(v int) *JourneyConfigurationProperties {
	c.properties[journeyPropertyNoteMargin] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: journeyPropertyNoteMargin,
			Val:  v,
		},
	}
	return c
}

func (c *JourneyConfigurationProperties) SetMessageMargin(v int) *JourneyConfigurationProperties {
	c.properties[journeyPropertyMessageMargin] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: journeyPropertyMessageMargin,
			Val:  v,
		},
	}
	return c
}

func (c *JourneyConfigurationProperties) SetMessageAlign(v string) *JourneyConfigurationProperties {
	c.properties[journeyPropertyMessageAlign] = &basediagram.StringProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: journeyPropertyMessageAlign,
			Val:  v,
		},
	}
	return c
}

func (c *JourneyConfigurationProperties) SetBottomMarginAdj(v int) *JourneyConfigurationProperties {
	c.properties[journeyPropertyBottomMarginAdj] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: journeyPropertyBottomMarginAdj,
			Val:  v,
		},
	}
	return c
}

func (c *JourneyConfigurationProperties) SetRightAngles(v bool) *JourneyConfigurationProperties {
	c.properties[journeyPropertyRightAngles] = &basediagram.BoolProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: journeyPropertyRightAngles,
			Val:  v,
		},
	}
	return c
}

func (c *JourneyConfigurationProperties) SetTaskFontSize(v int) *JourneyConfigurationProperties {
	c.properties[journeyPropertyTaskFontSize] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: journeyPropertyTaskFontSize,
			Val:  v,
		},
	}
	return c
}

func (c *JourneyConfigurationProperties) SetTaskFontFamily(v string) *JourneyConfigurationProperties {
	c.properties[journeyPropertyTaskFontFamily] = &basediagram.StringProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: journeyPropertyTaskFontFamily,
			Val:  v,
		},
	}
	return c
}

func (c *JourneyConfigurationProperties) SetTaskMargin(v int) *JourneyConfigurationProperties {
	c.properties[journeyPropertyTaskMargin] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: journeyPropertyTaskMargin,
			Val:  v,
		},
	}
	return c
}

func (c *JourneyConfigurationProperties) SetActivationWidth(v int) *JourneyConfigurationProperties {
	c.properties[journeyPropertyActivationWidth] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: journeyPropertyActivationWidth,
			Val:  v,
		},
	}
	return c
}

func (c *JourneyConfigurationProperties) SetTextPlacement(v string) *JourneyConfigurationProperties {
	c.properties[journeyPropertyTextPlacement] = &basediagram.StringProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: journeyPropertyTextPlacement,
			Val:  v,
		},
	}
	return c
}

func (c *JourneyConfigurationProperties) SetActorColours(v []string) *JourneyConfigurationProperties {
	c.properties[journeyPropertyActorColours] = &basediagram.StringArrayProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: journeyPropertyActorColours,
			Val:  v,
		},
	}
	return c
}

func (c *JourneyConfigurationProperties) SetSectionFills(v []string) *JourneyConfigurationProperties {
	c.properties[journeyPropertySectionFills] = &basediagram.StringArrayProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: journeyPropertySectionFills,
			Val:  v,
		},
	}
	return c
}

func (c *JourneyConfigurationProperties) SetSectionColours(v []string) *JourneyConfigurationProperties {
	c.properties[journeyPropertySectionColours] = &basediagram.StringArrayProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: journeyPropertySectionColours,
			Val:  v,
		},
	}
	return c
}

func (c JourneyConfigurationProperties) String() string {
	var sb strings.Builder
	sb.WriteString(c.ConfigurationProperties.String())

	if len(c.properties) > 0 {
		sb.WriteString(baseJourneyConfigurationProperties)
		for _, prop := range c.properties {
			sb.WriteString(prop.Format())
		}
	}

	return sb.String()
}
