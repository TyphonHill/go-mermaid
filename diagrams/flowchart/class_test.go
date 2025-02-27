package flowchart

import (
	"reflect"
	"strings"
	"testing"
)

func TestNewClass(t *testing.T) {
	tests := []struct {
		name      string
		className string
		wantClass *Class
	}{
		{
			name:      "Create new class with basic name",
			className: "testClass",
			wantClass: &Class{
				Name:  "testClass",
				Style: NewNodeStyle(),
			},
		},
		{
			name:      "Create new class with empty name",
			className: "",
			wantClass: &Class{
				Name:  "",
				Style: NewNodeStyle(),
			},
		},
		{
			name:      "Create new class with special characters",
			className: "test-class_123",
			wantClass: &Class{
				Name:  "test-class_123",
				Style: NewNodeStyle(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewClass(tt.className)
			if !reflect.DeepEqual(got, tt.wantClass) {
				t.Errorf("NewClass() = %v, want %v", got, tt.wantClass)
			}
		})
	}
}

func TestClass_String(t *testing.T) {
	tests := []struct {
		name     string
		class    *Class
		setup    func(*Class)
		contains []string
	}{
		{
			name: "Class with default style",
			class: &Class{
				Name:  "testClass",
				Style: NewNodeStyle(),
			},
			contains: []string{
				"classDef testClass stroke-width:1,stroke-dasharray:0",
			},
		},
		{
			name: "Class with custom style",
			class: &Class{
				Name:  "customClass",
				Style: NewNodeStyle(),
			},
			setup: func(c *Class) {
				c.Style.Fill = "#f9f9f9"
				c.Style.Color = "#333333"
				c.Style.Stroke = "#666666"
				c.Style.StrokeWidth = 2
				c.Style.StrokeDash = "5"
			},
			contains: []string{
				"classDef customClass color:#333333,fill:#f9f9f9,stroke:#666666,stroke-width:2,stroke-dasharray:5",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup(tt.class)
			}

			got := tt.class.String()
			for _, want := range tt.contains {
				if !strings.Contains(got, want) {
					t.Errorf("String() missing expected content %q in:\n%s", want, got)
				}
			}
		})
	}
}
