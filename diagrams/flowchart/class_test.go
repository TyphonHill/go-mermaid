package flowchart

import (
	"reflect"
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
		name    string
		class   *Class
		setup   func(*Class)
		wantStr string
	}{
		{
			name: "Class with default style",
			class: &Class{
				Name:  "testClass",
				Style: NewNodeStyle(),
			},
			wantStr: "\tclassDef testClass stroke-width:1,stroke-dasharray:0\n",
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
			wantStr: "\tclassDef customClass color:#333333,fill:#f9f9f9,stroke:#666666,stroke-width:2,stroke-dasharray:5\n",
		},
		{
			name: "Class with nil style",
			class: &Class{
				Name:  "nilStyle",
				Style: nil,
			},
			wantStr: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup(tt.class)
			}

			got := tt.class.String()
			if got != tt.wantStr {
				t.Errorf("Class.String() = %v, want %v", got, tt.wantStr)
			}
		})
	}
}
