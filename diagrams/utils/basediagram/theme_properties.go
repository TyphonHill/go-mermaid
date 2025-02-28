package basediagram

import (
	"fmt"
	"strings"
)

type ThemeName string

type Theme struct {
	Name      ThemeName
	Variables map[string]interface{}
}

const (
	baseThemeString     = Indentation + "theme: %s\n"
	themeVariableString = Indentation + Indentation + "%s: %v\n"
)

const (
	ThemeDefault ThemeName = "default"
	ThemeNeutral ThemeName = "neutral"
	ThemeDark    ThemeName = "dark"
	ThemeForest  ThemeName = "forest"
	ThemeBase    ThemeName = "base"
)

const (
	ThemeVarDarkMode           = "darkMode"
	ThemeVarBackground         = "background"
	ThemeVarFontFamily         = "fontFamily"
	ThemeVarFontSize           = "fontSize"
	ThemeVarPrimaryColor       = "primaryColor"
	ThemeVarPrimaryTextColor   = "primaryTextColor"
	ThemeVarSecondaryColor     = "secondaryColor"
	ThemeVarPrimaryBorderColor = "primaryBorderColor"
	ThemeVarTertiaryColor      = "tertiaryColor"
	ThemeVarNoteBkgColor       = "noteBkgColor"
	ThemeVarNoteTextColor      = "noteTextColor"
	ThemeVarNoteBorderColor    = "noteBorderColor"
	ThemeVarLineColor          = "lineColor"
	ThemeVarTextColor          = "textColor"
	ThemeVarMainBkg            = "mainBkg"
	ThemeVarErrorBkgColor      = "errorBkgColor"
	ThemeVarErrorTextColor     = "errorTextColor"
)

func NewTheme() Theme {
	return Theme{
		Name:      ThemeDefault,
		Variables: make(map[string]interface{}),
	}
}

func (t *Theme) SetTheme(name ThemeName) *Theme {
	t.Name = name
	return t
}

func (t *Theme) SetDarkMode(v bool) *Theme {
	if t.Variables == nil {
		t.Variables = make(map[string]interface{})
	}
	t.Variables[ThemeVarDarkMode] = v
	return t
}

func (t *Theme) SetBackground(v string) *Theme {
	if t.Variables == nil {
		t.Variables = make(map[string]interface{})
	}
	t.Variables[ThemeVarBackground] = v
	return t
}

func (t *Theme) SetFontFamily(v string) *Theme {
	if t.Variables == nil {
		t.Variables = make(map[string]interface{})
	}
	t.Variables[ThemeVarFontFamily] = v
	return t
}

func (t *Theme) SetFontSize(v string) *Theme {
	if t.Variables == nil {
		t.Variables = make(map[string]interface{})
	}
	t.Variables[ThemeVarFontSize] = v
	return t
}

func (t *Theme) SetPrimaryColor(v string) *Theme {
	if t.Variables == nil {
		t.Variables = make(map[string]interface{})
	}
	t.Variables[ThemeVarPrimaryColor] = v
	return t
}

func (t *Theme) SetPrimaryTextColor(v string) *Theme {
	if t.Variables == nil {
		t.Variables = make(map[string]interface{})
	}
	t.Variables[ThemeVarPrimaryTextColor] = v
	return t
}

func (t *Theme) SetSecondaryColor(v string) *Theme {
	if t.Variables == nil {
		t.Variables = make(map[string]interface{})
	}
	t.Variables[ThemeVarSecondaryColor] = v
	return t
}

func (t *Theme) SetPrimaryBorderColor(v string) *Theme {
	if t.Variables == nil {
		t.Variables = make(map[string]interface{})
	}
	t.Variables[ThemeVarPrimaryBorderColor] = v
	return t
}

func (t *Theme) SetTertiaryColor(v string) *Theme {
	if t.Variables == nil {
		t.Variables = make(map[string]interface{})
	}
	t.Variables[ThemeVarTertiaryColor] = v
	return t
}

func (t *Theme) SetNoteBkgColor(v string) *Theme {
	if t.Variables == nil {
		t.Variables = make(map[string]interface{})
	}
	t.Variables[ThemeVarNoteBkgColor] = v
	return t
}

func (t *Theme) SetNoteTextColor(v string) *Theme {
	if t.Variables == nil {
		t.Variables = make(map[string]interface{})
	}
	t.Variables[ThemeVarNoteTextColor] = v
	return t
}

func (t *Theme) SetNoteBorderColor(v string) *Theme {
	if t.Variables == nil {
		t.Variables = make(map[string]interface{})
	}
	t.Variables[ThemeVarNoteBorderColor] = v
	return t
}

func (t *Theme) SetLineColor(v string) *Theme {
	if t.Variables == nil {
		t.Variables = make(map[string]interface{})
	}
	t.Variables[ThemeVarLineColor] = v
	return t
}

func (t *Theme) SetTextColor(v string) *Theme {
	if t.Variables == nil {
		t.Variables = make(map[string]interface{})
	}
	t.Variables[ThemeVarTextColor] = v
	return t
}

func (t *Theme) SetMainBkg(v string) *Theme {
	if t.Variables == nil {
		t.Variables = make(map[string]interface{})
	}
	t.Variables[ThemeVarMainBkg] = v
	return t
}

func (t *Theme) SetErrorBkgColor(v string) *Theme {
	if t.Variables == nil {
		t.Variables = make(map[string]interface{})
	}
	t.Variables[ThemeVarErrorBkgColor] = v
	return t
}

func (t *Theme) SetErrorTextColor(v string) *Theme {
	if t.Variables == nil {
		t.Variables = make(map[string]interface{})
	}
	t.Variables[ThemeVarErrorTextColor] = v
	return t
}

func (t *Theme) String() string {
	if len(t.Variables) == 0 {
		return fmt.Sprintf(baseThemeString, t.Name)
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf(baseThemeString, t.Name))
	sb.WriteString(Indentation + "themeVariables:\n")

	for k, v := range t.Variables {
		sb.WriteString(fmt.Sprintf(themeVariableString, k, v))
	}

	return sb.String()
}
