package sequence

import (
	"strings"

	"github.com/TyphonHill/go-mermaid/diagrams/utils/basediagram"
)

const (
	baseSequenceConfigurationProperties    string = basediagram.Indentation + "sequence:\n"
	sequencePropertyArrowMarkerAbsolute    string = "arrowMarkerAbsolute"
	sequencePropertyHideUnusedParticipants string = "hideUnusedParticipants"
	sequencePropertyActivationWidth        string = "activationWidth"
	sequencePropertyDiagramMarginX         string = "diagramMarginX"
	sequencePropertyDiagramMarginY         string = "diagramMarginY"
	sequencePropertyActorMargin            string = "actorMargin"
	sequencePropertyWidth                  string = "width"
	sequencePropertyHeight                 string = "height"
	sequencePropertyBoxMargin              string = "boxMargin"
	sequencePropertyBoxTextMargin          string = "boxTextMargin"
	sequencePropertyNoteMargin             string = "noteMargin"
	sequencePropertyMessageMargin          string = "messageMargin"
	sequencePropertyMessageAlign           string = "messageAlign"
	sequencePropertyMirrorActors           string = "mirrorActors"
	sequencePropertyForceMenus             string = "forceMenus"
	sequencePropertyBottomMarginAdj        string = "bottomMarginAdj"
	sequencePropertyRightAngles            string = "rightAngles"
	sequencePropertyShowSequenceNumbers    string = "showSequenceNumbers"
	sequencePropertyActorFontSize          string = "actorFontSize"
	sequencePropertyActorFontFamily        string = "actorFontFamily"
	sequencePropertyActorFontWeight        string = "actorFontWeight"
	sequencePropertyNoteFontSize           string = "noteFontSize"
	sequencePropertyNoteFontFamily         string = "noteFontFamily"
	sequencePropertyNoteFontWeight         string = "noteFontWeight"
	sequencePropertyNoteAlign              string = "noteAlign"
	sequencePropertyMessageFontSize        string = "messageFontSize"
	sequencePropertyMessageFontFamily      string = "messageFontFamily"
	sequencePropertyMessageFontWeight      string = "messageFontWeight"
	sequencePropertyWrap                   string = "wrap"
	sequencePropertyWrapPadding            string = "wrapPadding"
	sequencePropertyLabelBoxWidth          string = "labelBoxWidth"
	sequencePropertyLabelBoxHeight         string = "labelBoxHeight"
)

// SequenceConfigurationProperties holds sequence-specific configuration
type SequenceConfigurationProperties struct {
	basediagram.ConfigurationProperties
	properties map[string]basediagram.DiagramProperty
}

func NewSequenceConfigurationProperties() SequenceConfigurationProperties {
	return SequenceConfigurationProperties{
		ConfigurationProperties: basediagram.NewConfigurationProperties(),
		properties:              make(map[string]basediagram.DiagramProperty),
	}
}

// Setters for each property
func (c *SequenceConfigurationProperties) SetArrowMarkerAbsolute(v bool) *SequenceConfigurationProperties {
	c.properties[sequencePropertyArrowMarkerAbsolute] = &basediagram.BoolProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: sequencePropertyArrowMarkerAbsolute,
			Val:  v,
		},
	}
	return c
}

func (c *SequenceConfigurationProperties) SetHideUnusedParticipants(v bool) *SequenceConfigurationProperties {
	c.properties[sequencePropertyHideUnusedParticipants] = &basediagram.BoolProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: sequencePropertyHideUnusedParticipants,
			Val:  v,
		},
	}
	return c
}

func (c *SequenceConfigurationProperties) SetActivationWidth(v int) *SequenceConfigurationProperties {
	c.properties[sequencePropertyActivationWidth] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: sequencePropertyActivationWidth,
			Val:  v,
		},
	}
	return c
}

func (c *SequenceConfigurationProperties) SetDiagramMarginX(v int) *SequenceConfigurationProperties {
	c.properties[sequencePropertyDiagramMarginX] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: sequencePropertyDiagramMarginX,
			Val:  v,
		},
	}
	return c
}

func (c *SequenceConfigurationProperties) SetDiagramMarginY(v int) *SequenceConfigurationProperties {
	c.properties[sequencePropertyDiagramMarginY] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: sequencePropertyDiagramMarginY,
			Val:  v,
		},
	}
	return c
}

func (c *SequenceConfigurationProperties) SetActorMargin(v int) *SequenceConfigurationProperties {
	c.properties[sequencePropertyActorMargin] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: sequencePropertyActorMargin,
			Val:  v,
		},
	}
	return c
}

func (c *SequenceConfigurationProperties) SetWidth(v int) *SequenceConfigurationProperties {
	c.properties[sequencePropertyWidth] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: sequencePropertyWidth,
			Val:  v,
		},
	}
	return c
}

func (c *SequenceConfigurationProperties) SetHeight(v int) *SequenceConfigurationProperties {
	c.properties[sequencePropertyHeight] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: sequencePropertyHeight,
			Val:  v,
		},
	}
	return c
}

func (c *SequenceConfigurationProperties) SetBoxMargin(v int) *SequenceConfigurationProperties {
	c.properties[sequencePropertyBoxMargin] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: sequencePropertyBoxMargin,
			Val:  v,
		},
	}
	return c
}

func (c *SequenceConfigurationProperties) SetBoxTextMargin(v int) *SequenceConfigurationProperties {
	c.properties[sequencePropertyBoxTextMargin] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: sequencePropertyBoxTextMargin,
			Val:  v,
		},
	}
	return c
}

func (c *SequenceConfigurationProperties) SetNoteMargin(v int) *SequenceConfigurationProperties {
	c.properties[sequencePropertyNoteMargin] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: sequencePropertyNoteMargin,
			Val:  v,
		},
	}
	return c
}

func (c *SequenceConfigurationProperties) SetMessageMargin(v int) *SequenceConfigurationProperties {
	c.properties[sequencePropertyMessageMargin] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: sequencePropertyMessageMargin,
			Val:  v,
		},
	}
	return c
}

func (c *SequenceConfigurationProperties) SetMessageAlign(v string) *SequenceConfigurationProperties {
	c.properties[sequencePropertyMessageAlign] = &basediagram.StringProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: sequencePropertyMessageAlign,
			Val:  v,
		},
	}
	return c
}

func (c *SequenceConfigurationProperties) SetMirrorActors(v bool) *SequenceConfigurationProperties {
	c.properties[sequencePropertyMirrorActors] = &basediagram.BoolProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: sequencePropertyMirrorActors,
			Val:  v,
		},
	}
	return c
}

func (c *SequenceConfigurationProperties) SetForceMenus(v bool) *SequenceConfigurationProperties {
	c.properties[sequencePropertyForceMenus] = &basediagram.BoolProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: sequencePropertyForceMenus,
			Val:  v,
		},
	}
	return c
}

func (c *SequenceConfigurationProperties) SetBottomMarginAdj(v int) *SequenceConfigurationProperties {
	c.properties[sequencePropertyBottomMarginAdj] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: sequencePropertyBottomMarginAdj,
			Val:  v,
		},
	}
	return c
}

func (c *SequenceConfigurationProperties) SetRightAngles(v bool) *SequenceConfigurationProperties {
	c.properties[sequencePropertyRightAngles] = &basediagram.BoolProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: sequencePropertyRightAngles,
			Val:  v,
		},
	}
	return c
}

func (c *SequenceConfigurationProperties) SetShowSequenceNumbers(v bool) *SequenceConfigurationProperties {
	c.properties[sequencePropertyShowSequenceNumbers] = &basediagram.BoolProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: sequencePropertyShowSequenceNumbers,
			Val:  v,
		},
	}
	return c
}

func (c *SequenceConfigurationProperties) SetActorFontSize(v int) *SequenceConfigurationProperties {
	c.properties[sequencePropertyActorFontSize] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: sequencePropertyActorFontSize,
			Val:  v,
		},
	}
	return c
}

func (c *SequenceConfigurationProperties) SetActorFontFamily(v string) *SequenceConfigurationProperties {
	c.properties[sequencePropertyActorFontFamily] = &basediagram.StringProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: sequencePropertyActorFontFamily,
			Val:  v,
		},
	}
	return c
}

func (c *SequenceConfigurationProperties) SetActorFontWeight(v int) *SequenceConfigurationProperties {
	c.properties[sequencePropertyActorFontWeight] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: sequencePropertyActorFontWeight,
			Val:  v,
		},
	}
	return c
}

func (c *SequenceConfigurationProperties) SetNoteFontSize(v int) *SequenceConfigurationProperties {
	c.properties[sequencePropertyNoteFontSize] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: sequencePropertyNoteFontSize,
			Val:  v,
		},
	}
	return c
}

func (c *SequenceConfigurationProperties) SetNoteFontFamily(v string) *SequenceConfigurationProperties {
	c.properties[sequencePropertyNoteFontFamily] = &basediagram.StringProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: sequencePropertyNoteFontFamily,
			Val:  v,
		},
	}
	return c
}

func (c *SequenceConfigurationProperties) SetNoteFontWeight(v int) *SequenceConfigurationProperties {
	c.properties[sequencePropertyNoteFontWeight] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: sequencePropertyNoteFontWeight,
			Val:  v,
		},
	}
	return c
}

func (c *SequenceConfigurationProperties) SetNoteAlign(v string) *SequenceConfigurationProperties {
	c.properties[sequencePropertyNoteAlign] = &basediagram.StringProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: sequencePropertyNoteAlign,
			Val:  v,
		},
	}
	return c
}

func (c *SequenceConfigurationProperties) SetMessageFontSize(v int) *SequenceConfigurationProperties {
	c.properties[sequencePropertyMessageFontSize] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: sequencePropertyMessageFontSize,
			Val:  v,
		},
	}
	return c
}

func (c *SequenceConfigurationProperties) SetMessageFontFamily(v string) *SequenceConfigurationProperties {
	c.properties[sequencePropertyMessageFontFamily] = &basediagram.StringProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: sequencePropertyMessageFontFamily,
			Val:  v,
		},
	}
	return c
}

func (c *SequenceConfigurationProperties) SetMessageFontWeight(v int) *SequenceConfigurationProperties {
	c.properties[sequencePropertyMessageFontWeight] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: sequencePropertyMessageFontWeight,
			Val:  v,
		},
	}
	return c
}

func (c *SequenceConfigurationProperties) SetWrap(v bool) *SequenceConfigurationProperties {
	c.properties[sequencePropertyWrap] = &basediagram.BoolProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: sequencePropertyWrap,
			Val:  v,
		},
	}
	return c
}

func (c *SequenceConfigurationProperties) SetWrapPadding(v int) *SequenceConfigurationProperties {
	c.properties[sequencePropertyWrapPadding] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: sequencePropertyWrapPadding,
			Val:  v,
		},
	}
	return c
}

func (c *SequenceConfigurationProperties) SetLabelBoxWidth(v int) *SequenceConfigurationProperties {
	c.properties[sequencePropertyLabelBoxWidth] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: sequencePropertyLabelBoxWidth,
			Val:  v,
		},
	}
	return c
}

func (c *SequenceConfigurationProperties) SetLabelBoxHeight(v int) *SequenceConfigurationProperties {
	c.properties[sequencePropertyLabelBoxHeight] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: sequencePropertyLabelBoxHeight,
			Val:  v,
		},
	}
	return c
}

func (c SequenceConfigurationProperties) String() string {
	var sb strings.Builder
	sb.WriteString(c.ConfigurationProperties.String())

	if len(c.properties) > 0 {
		sb.WriteString(baseSequenceConfigurationProperties)
		for _, prop := range c.properties {
			sb.WriteString(prop.Format())
		}
	}

	return sb.String()
}
