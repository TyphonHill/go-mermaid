package basediagram

import (
	"strings"
	"testing"
)

func TestNewConfigurationProperties(t *testing.T) {
	got := NewConfigurationProperties()
	if got.String() == "" {
		t.Error("NewConfigurationProperties() returned empty configuration")
	}
}

func TestConfigurationProperties_String(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(*ConfigurationProperties)
		contains []string
	}{
		{
			name: "Empty configuration",
			contains: []string{
				"",
			},
		},
		{
			name: "Configuration with single property",
			setup: func(c *ConfigurationProperties) {
				c.SetFontSize(14)
			},
			contains: []string{
				"fontSize: 14",
			},
		},
		{
			name: "Configuration with multiple properties",
			setup: func(c *ConfigurationProperties) {
				c.SetFontSize(14)
				c.SetFontFamily("Arial")
				c.SetTheme(ThemeDark)
			},
			contains: []string{
				"fontSize: 14",
				"fontFamily: Arial",
				"theme: dark",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := &ConfigurationProperties{}
			if tt.setup != nil {
				tt.setup(config)
			}

			got := config.String()
			for _, want := range tt.contains {
				if !strings.Contains(got, want) {
					t.Errorf("String() missing expected content %q in:\n%s", want, got)
				}
			}
		})
	}
}

func TestConfigurationProperties_Setters(t *testing.T) {
	tests := []struct {
		name  string
		setup func(*ConfigurationProperties)
		check func(*testing.T, *ConfigurationProperties)
	}{
		{
			name: "Set font size",
			setup: func(c *ConfigurationProperties) {
				c.SetFontSize(14)
			},
			check: func(t *testing.T, c *ConfigurationProperties) {
				if !strings.Contains(c.String(), "fontSize: 14") {
					t.Error("FontSize not set correctly")
				}
			},
		},
		{
			name: "Set font family",
			setup: func(c *ConfigurationProperties) {
				c.SetFontFamily("Arial")
			},
			check: func(t *testing.T, c *ConfigurationProperties) {
				if !strings.Contains(c.String(), "fontFamily: Arial") {
					t.Error("FontFamily not set correctly")
				}
			},
		},
		{
			name: "Set theme",
			setup: func(c *ConfigurationProperties) {
				c.SetTheme(ThemeDark)
			},
			check: func(t *testing.T, c *ConfigurationProperties) {
				if !strings.Contains(c.String(), "theme: dark") {
					t.Error("Theme not set correctly")
				}
			},
		},
		{
			name: "Set max text size",
			setup: func(c *ConfigurationProperties) {
				c.SetMaxTextSize(20)
			},
			check: func(t *testing.T, c *ConfigurationProperties) {
				if !strings.Contains(c.String(), "maxTextSize: 20") {
					t.Error("MaxTextSize not set correctly")
				}
			},
		},
		{
			name: "Set max edges",
			setup: func(c *ConfigurationProperties) {
				c.SetMaxEdges(100)
			},
			check: func(t *testing.T, c *ConfigurationProperties) {
				if !strings.Contains(c.String(), "maxEdges: 100") {
					t.Error("MaxEdges not set correctly")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := &ConfigurationProperties{}
			tt.setup(config)
			tt.check(t, config)
		})
	}
}
