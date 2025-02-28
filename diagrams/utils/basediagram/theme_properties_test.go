package basediagram

import (
	"strings"
	"testing"
)

func TestNewTheme(t *testing.T) {
	got := NewTheme()
	if got.Name != ThemeDefault {
		t.Errorf("NewTheme().Name = %v, want %v", got.Name, ThemeDefault)
	}
	if got.Variables == nil {
		t.Error("NewTheme().Variables is nil")
	}
	if len(got.Variables) != 0 {
		t.Errorf("NewTheme().Variables length = %v, want 0", len(got.Variables))
	}
}

func TestTheme_String(t *testing.T) {
	tests := []struct {
		name     string
		theme    Theme
		setup    func(*Theme)
		contains []string
	}{
		{
			name: "Theme without variables",
			theme: Theme{
				Name:      ThemeDark,
				Variables: make(map[string]interface{}),
			},
			contains: []string{
				"theme: dark",
			},
		},
		{
			name: "Theme with single variable",
			theme: Theme{
				Name:      ThemeDark,
				Variables: make(map[string]interface{}),
			},
			setup: func(t *Theme) {
				t.SetDarkMode(true)
			},
			contains: []string{
				"theme: dark",
				"themeVariables:",
				"darkMode: true",
			},
		},
		{
			name: "Theme with multiple variables",
			theme: Theme{
				Name:      ThemeDark,
				Variables: make(map[string]interface{}),
			},
			setup: func(t *Theme) {
				t.SetDarkMode(true)
				t.SetBackground("#000000")
				t.SetFontFamily("Arial")
			},
			contains: []string{
				"theme: dark",
				"themeVariables:",
				"darkMode: true",
				"background: #000000",
				"fontFamily: Arial",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup(&tt.theme)
			}

			got := tt.theme.String()
			for _, want := range tt.contains {
				if !strings.Contains(got, want) {
					t.Errorf("String() missing expected content %q in:\n%s", want, got)
				}
			}
		})
	}
}

func TestTheme_Setters(t *testing.T) {
	tests := []struct {
		name      string
		setup     func(*Theme)
		varName   string
		wantValue interface{}
	}{
		{
			name: "Set theme name",
			setup: func(t *Theme) {
				t.SetTheme(ThemeDark)
			},
			varName:   "name",
			wantValue: ThemeDark,
		},
		{
			name: "Set dark mode",
			setup: func(t *Theme) {
				t.SetDarkMode(true)
			},
			varName:   ThemeVarDarkMode,
			wantValue: true,
		},
		{
			name: "Set background",
			setup: func(t *Theme) {
				t.SetBackground("#000000")
			},
			varName:   ThemeVarBackground,
			wantValue: "#000000",
		},
		{
			name: "Set font family",
			setup: func(t *Theme) {
				t.SetFontFamily("Arial")
			},
			varName:   ThemeVarFontFamily,
			wantValue: "Arial",
		},
		{
			name: "Set font size",
			setup: func(t *Theme) {
				t.SetFontSize("14px")
			},
			varName:   ThemeVarFontSize,
			wantValue: "14px",
		},
		{
			name: "Set primary color",
			setup: func(t *Theme) {
				t.SetPrimaryColor("#ff0000")
			},
			varName:   ThemeVarPrimaryColor,
			wantValue: "#ff0000",
		},
		{
			name: "Set primary text color",
			setup: func(t *Theme) {
				t.SetPrimaryTextColor("#ffffff")
			},
			varName:   ThemeVarPrimaryTextColor,
			wantValue: "#ffffff",
		},
		{
			name: "Set secondary color",
			setup: func(t *Theme) {
				t.SetSecondaryColor("#00ff00")
			},
			varName:   ThemeVarSecondaryColor,
			wantValue: "#00ff00",
		},
		{
			name: "Set primary border color",
			setup: func(t *Theme) {
				t.SetPrimaryBorderColor("#cccccc")
			},
			varName:   ThemeVarPrimaryBorderColor,
			wantValue: "#cccccc",
		},
		{
			name: "Set tertiary color",
			setup: func(t *Theme) {
				t.SetTertiaryColor("#0000ff")
			},
			varName:   ThemeVarTertiaryColor,
			wantValue: "#0000ff",
		},
		{
			name: "Set note background color",
			setup: func(t *Theme) {
				t.SetNoteBkgColor("#ffffcc")
			},
			varName:   ThemeVarNoteBkgColor,
			wantValue: "#ffffcc",
		},
		{
			name: "Set note text color",
			setup: func(t *Theme) {
				t.SetNoteTextColor("#333333")
			},
			varName:   ThemeVarNoteTextColor,
			wantValue: "#333333",
		},
		{
			name: "Set note border color",
			setup: func(t *Theme) {
				t.SetNoteBorderColor("#999999")
			},
			varName:   ThemeVarNoteBorderColor,
			wantValue: "#999999",
		},
		{
			name: "Set line color",
			setup: func(t *Theme) {
				t.SetLineColor("#666666")
			},
			varName:   ThemeVarLineColor,
			wantValue: "#666666",
		},
		{
			name: "Set text color",
			setup: func(t *Theme) {
				t.SetTextColor("#000000")
			},
			varName:   ThemeVarTextColor,
			wantValue: "#000000",
		},
		{
			name: "Set main background",
			setup: func(t *Theme) {
				t.SetMainBkg("#ffffff")
			},
			varName:   ThemeVarMainBkg,
			wantValue: "#ffffff",
		},
		{
			name: "Set error background color",
			setup: func(t *Theme) {
				t.SetErrorBkgColor("#ffcccc")
			},
			varName:   ThemeVarErrorBkgColor,
			wantValue: "#ffcccc",
		},
		{
			name: "Set error text color",
			setup: func(t *Theme) {
				t.SetErrorTextColor("#ff0000")
			},
			varName:   ThemeVarErrorTextColor,
			wantValue: "#ff0000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var theme Theme
			tt.setup(&theme)

			// Test variable was set correctly
			if tt.varName == "name" {
				if theme.Name != tt.wantValue {
					t.Errorf("Theme name = %v, want %v", theme.Name, tt.wantValue)
				}
			} else {
				got, exists := theme.Variables[tt.varName]
				if !exists {
					t.Errorf("Variable %q was not set", tt.varName)
					return
				}
				if got != tt.wantValue {
					t.Errorf("Variable %q = %v, want %v", tt.varName, got, tt.wantValue)
				}
			}
		})
	}
}
