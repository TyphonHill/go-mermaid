package block

import (
	"strings"

	"github.com/TyphonHill/go-mermaid/diagrams/utils/basediagram"
)

const (
	baseBlockConfigurationProperties string = "block:\n"
	blockPropertyPadding             string = "padding"
)

// BlockConfigurationProperties holds block-specific configuration
type BlockConfigurationProperties struct {
	basediagram.ConfigurationProperties
	properties map[string]basediagram.DiagramProperty
}

func NewBlockConfigurationProperties() BlockConfigurationProperties {
	return BlockConfigurationProperties{
		ConfigurationProperties: basediagram.NewConfigurationProperties(),
		properties:              make(map[string]basediagram.DiagramProperty),
	}
}

func (c *BlockConfigurationProperties) SetPadding(v int) *BlockConfigurationProperties {
	c.properties[blockPropertyPadding] = &basediagram.IntProperty{
		BaseProperty: basediagram.BaseProperty{
			Name: blockPropertyPadding,
			Val:  v,
		},
	}
	return c
}

func (c BlockConfigurationProperties) String() string {
	var sb strings.Builder
	sb.WriteString(c.ConfigurationProperties.String())

	if len(c.properties) > 0 {
		sb.WriteString(baseBlockConfigurationProperties)
		for _, prop := range c.properties {
			sb.WriteString(prop.Format())
		}
	}

	return sb.String()
}
