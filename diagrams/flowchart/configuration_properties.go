package flowchart

import (
	"strings"

	"github.com/TyphonHill/go-mermaid/diagrams/utils/basediagram"
)

const (
	baseFlowchartConfigurationProperties string = basediagram.Indentation + "flowchart:\n"
	flowchartPropertyTitleTopMargin      string = "titleTopMargin"
	flowchartPropertyDiagramPadding      string = "diagramPadding"
	flowchartPropertyHtmlLabels          string = "htmlLabels"
	flowchartPropertyNodeSpacing         string = "nodeSpacing"
	flowchartPropertyRankSpacing         string = "rankSpacing"
	flowchartPropertyCurve               string = "curve"
	flowchartPropertyPadding             string = "padding"
	flowchartPropertyDefaultRenderer     string = "defaultRenderer"
	flowchartPropertyWrappingWidth       string = "wrappingWidth"
	flowchartPropertyArrowMarkerAbsolute string = "arrowMarkerAbsolute"
)

// FlowchartConfigurationProperties holds flowchart-specific configuration
type FlowchartConfigurationProperties struct {
	basediagram.ConfigurationProperties
	properties map[string]basediagram.DiagramProperty
}

func NewFlowchartConfigurationProperties() FlowchartConfigurationProperties {
	return FlowchartConfigurationProperties{
		ConfigurationProperties: basediagram.NewConfigurationProperties(),
		properties:              make(map[string]basediagram.DiagramProperty),
	}
}

func (c *FlowchartConfigurationProperties) SetTitleTopMargin(v int) *FlowchartConfigurationProperties {
	c.properties[flowchartPropertyTitleTopMargin] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: flowchartPropertyTitleTopMargin,
			Val:  v,
		},
	}
	return c
}

func (c *FlowchartConfigurationProperties) SetDiagramPadding(v int) *FlowchartConfigurationProperties {
	c.properties[flowchartPropertyDiagramPadding] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: flowchartPropertyDiagramPadding,
			Val:  v,
		},
	}
	return c
}

func (c *FlowchartConfigurationProperties) SetHtmlLabels(v bool) *FlowchartConfigurationProperties {
	c.properties[flowchartPropertyHtmlLabels] = &basediagram.BoolProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: flowchartPropertyHtmlLabels,
			Val:  v,
		},
	}
	return c
}

func (c *FlowchartConfigurationProperties) SetNodeSpacing(v int) *FlowchartConfigurationProperties {
	c.properties[flowchartPropertyNodeSpacing] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: flowchartPropertyNodeSpacing,
			Val:  v,
		},
	}
	return c
}

func (c *FlowchartConfigurationProperties) SetRankSpacing(v int) *FlowchartConfigurationProperties {
	c.properties[flowchartPropertyRankSpacing] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: flowchartPropertyRankSpacing,
			Val:  v,
		},
	}
	return c
}

func (c *FlowchartConfigurationProperties) SetCurve(v string) *FlowchartConfigurationProperties {
	c.properties[flowchartPropertyCurve] = &basediagram.StringProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: flowchartPropertyCurve,
			Val:  v,
		},
	}
	return c
}

func (c *FlowchartConfigurationProperties) SetPadding(v int) *FlowchartConfigurationProperties {
	c.properties[flowchartPropertyPadding] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: flowchartPropertyPadding,
			Val:  v,
		},
	}
	return c
}

func (c *FlowchartConfigurationProperties) SetDefaultRenderer(v string) *FlowchartConfigurationProperties {
	c.properties[flowchartPropertyDefaultRenderer] = &basediagram.StringProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: flowchartPropertyDefaultRenderer,
			Val:  v,
		},
	}
	return c
}

func (c *FlowchartConfigurationProperties) SetWrappingWidth(v int) *FlowchartConfigurationProperties {
	c.properties[flowchartPropertyWrappingWidth] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: flowchartPropertyWrappingWidth,
			Val:  v,
		},
	}
	return c
}

func (c *FlowchartConfigurationProperties) SetArrowMarkerAbsolute(v bool) *FlowchartConfigurationProperties {
	c.properties[flowchartPropertyArrowMarkerAbsolute] = &basediagram.BoolProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: flowchartPropertyArrowMarkerAbsolute,
			Val:  v,
		},
	}
	return c
}

func (c FlowchartConfigurationProperties) String() string {
	var sb strings.Builder
	sb.WriteString(c.ConfigurationProperties.String())

	if len(c.properties) > 0 {
		sb.WriteString(baseFlowchartConfigurationProperties)
		for _, prop := range c.properties {
			sb.WriteString(prop.Format())
		}
	}

	return sb.String()
}
