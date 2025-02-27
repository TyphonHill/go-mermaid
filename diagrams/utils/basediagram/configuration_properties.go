package basediagram

import (
	"fmt"
	"strings"
)

type ConfigurationProperties struct {
	theme       Theme
	maxTextSize int
	maxEdges    int
	fontSize    int
}

const (
	configPropertyBase        = "config:\n"
	configPropertyTheme       = "%stheme: %s\n"
	configPropertyMaxTextSize = "%smaxTextSize: %d\n"
	configPropertyMaxEdges    = "%smaxEdges: %d\n"
	configPropertyFontSize    = "%sfontSize: %d\n"
)

func (c *ConfigurationProperties) SetTheme(theme Theme) *ConfigurationProperties {
	c.theme = theme
	return c
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
	sb.WriteString(fmt.Sprintf(configPropertyTheme, Indentation, c.theme))
	sb.WriteString(fmt.Sprintf(configPropertyMaxTextSize, Indentation, c.maxTextSize))
	sb.WriteString(fmt.Sprintf(configPropertyMaxEdges, Indentation, c.maxEdges))
	sb.WriteString(fmt.Sprintf(configPropertyFontSize, Indentation, c.fontSize))

	return sb.String()
}
