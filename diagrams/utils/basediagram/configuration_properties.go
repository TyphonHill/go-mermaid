package basediagram

import (
	"fmt"
	"strings"
)

type ConfigurationProperties struct {
	Theme
	maxTextSize int
	maxEdges    int
	fontSize    int
}

const (
	configPropertyBase        = "config:\n"
	configPropertyMaxTextSize = "%smaxTextSize: %d\n"
	configPropertyMaxEdges    = "%smaxEdges: %d\n"
	configPropertyFontSize    = "%sfontSize: %d\n"
)

func NewConfigurationProperties() ConfigurationProperties {
	return ConfigurationProperties{
		Theme: Theme{
			Name: ThemeDefault,
		},
		maxTextSize: 50000,
		maxEdges:    500,
		fontSize:    16,
	}
}

func (c *ConfigurationProperties) SetMaxTextSize(maxTextSize int) *ConfigurationProperties {
	c.maxTextSize = maxTextSize
	return c
}

func (c *ConfigurationProperties) SetMaxEdges(maxEdges int) *ConfigurationProperties {
	c.maxEdges = maxEdges
	return c
}

func (c *ConfigurationProperties) SetFontSize(fontSize int) *ConfigurationProperties {
	c.fontSize = fontSize
	return c
}

func (c *ConfigurationProperties) String() string {
	var sb strings.Builder

	sb.WriteString(configPropertyBase)

	sb.WriteString(c.Theme.String())

	sb.WriteString(fmt.Sprintf(configPropertyMaxTextSize, Indentation, c.maxTextSize))
	sb.WriteString(fmt.Sprintf(configPropertyMaxEdges, Indentation, c.maxEdges))
	sb.WriteString(fmt.Sprintf(configPropertyFontSize, Indentation, c.fontSize))

	return sb.String()
}
