package basediagram

import (
	"strings"
	"testing"
)

type testConfig = *ConfigurationProperties

func TestNewBaseDiagram(t *testing.T) {
	tests := []struct {
		name string
		conf testConfig
		want *BaseDiagram[testConfig]
	}{
		{
			name: "Create new base diagram with empty config",
			conf: &ConfigurationProperties{},
			want: &BaseDiagram[testConfig]{
				Config: &ConfigurationProperties{},
			},
		},
		{
			name: "Create new base diagram with config",
			conf: func() testConfig {
				c := new(ConfigurationProperties)
				c.SetFontSize(14)
				return c
			}(),
			want: &BaseDiagram[testConfig]{
				Config: func() testConfig {
					c := new(ConfigurationProperties)
					c.SetFontSize(14)
					return c
				}(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewBaseDiagram(tt.conf)
			if got.Config.String() != tt.want.Config.String() {
				t.Errorf("NewBaseDiagram() = %v, want %v", got.Config, tt.want.Config)
			}
		})
	}
}

func TestBaseDiagram_String(t *testing.T) {
	tests := []struct {
		name     string
		diagram  *BaseDiagram[testConfig]
		content  string
		setup    func(*BaseDiagram[testConfig])
		contains []string
	}{
		{
			name: "Empty diagram",
			diagram: &BaseDiagram[testConfig]{
				Config: &ConfigurationProperties{},
			},
			content: "test content",
			contains: []string{
				"test content",
			},
		},
		{
			name: "Diagram with configuration",
			diagram: &BaseDiagram[testConfig]{
				Config: &ConfigurationProperties{},
			},
			content: "test content",
			setup: func(d *BaseDiagram[testConfig]) {
				d.Config.SetFontSize(14)
				d.Config.SetFontFamily("Arial")
			},
			contains: []string{
				"fontSize: 14",
				"fontFamily: Arial",
				"test content",
			},
		},
		{
			name: "Diagram with theme",
			diagram: &BaseDiagram[testConfig]{
				Config: &ConfigurationProperties{},
			},
			content: "test content",
			setup: func(d *BaseDiagram[testConfig]) {
				d.Config.SetTheme("dark")
			},
			contains: []string{
				"config:",
				"theme: dark",
				"test content",
			},
		},
		{
			name: "Diagram with theme and configuration",
			diagram: &BaseDiagram[testConfig]{
				Config: &ConfigurationProperties{},
			},
			content: "test content",
			setup: func(d *BaseDiagram[testConfig]) {
				d.Config.SetTheme("dark")
				d.Config.SetFontSize(14)
			},
			contains: []string{
				"config:",
				"theme: dark",
				"fontSize: 14",
				"test content",
			},
		},
		{
			name: "Diagram with title",
			diagram: &BaseDiagram[testConfig]{
				Config: &ConfigurationProperties{},
			},
			content: "test content",
			setup: func(d *BaseDiagram[testConfig]) {
				d.SetTitle("Test Diagram")
			},
			contains: []string{
				"---",
				"title: Test Diagram",
				"---",
				"test content",
			},
		},
		{
			name: "Diagram with all features",
			diagram: &BaseDiagram[testConfig]{
				Config: &ConfigurationProperties{},
			},
			content: "test content",
			setup: func(d *BaseDiagram[testConfig]) {
				d.Config.SetTheme("dark")
				d.SetTitle("Test Diagram")
				d.Config.SetFontSize(14)
				d.Config.SetFontFamily("Arial")
			},
			contains: []string{
				"---",
				"title: Test Diagram",
				"---",
				"config:",
				"theme: dark",
				"fontSize: 14",
				"fontFamily: Arial",
				"test content",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup(tt.diagram)
			}

			got := tt.diagram.String(tt.content)
			for _, want := range tt.contains {
				if !strings.Contains(got, want) {
					t.Errorf("String() missing expected content %q in:\n%s", want, got)
				}
			}
		})
	}
}

func TestBaseDiagram_SetTheme(t *testing.T) {
	tests := []struct {
		name  string
		theme ThemeName
		want  ThemeName
	}{
		{
			name:  "Set default theme",
			theme: ThemeDefault,
			want:  ThemeDefault,
		},
		{
			name:  "Set dark theme",
			theme: ThemeDark,
			want:  ThemeDark,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diagram := &BaseDiagram[testConfig]{
				Config: &ConfigurationProperties{},
			}
			diagram.Config.SetTheme(tt.theme)

			if diagram.Config.Theme.Name != tt.want {
				t.Errorf("SetTheme() = %v, want %v", diagram.Config.Theme, tt.want)
			}
		})
	}
}

func TestBaseDiagram_SetTitle(t *testing.T) {
	tests := []struct {
		name  string
		title string
		want  string
	}{
		{
			name:  "Set empty title",
			title: "",
			want:  "",
		},
		{
			name:  "Set simple title",
			title: "Test Diagram",
			want:  "Test Diagram",
		},
		{
			name:  "Set title with special characters",
			title: "Test: Diagram!",
			want:  "Test: Diagram!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diagram := &BaseDiagram[testConfig]{
				Config: &ConfigurationProperties{},
			}
			result := diagram.SetTitle(tt.title)

			if diagram.Title != tt.want {
				t.Errorf("SetTitle() = %v, want %v", diagram.Title, tt.want)
			}

			if result != diagram {
				t.Error("SetTitle() should return diagram for chaining")
			}
		})
	}
}
