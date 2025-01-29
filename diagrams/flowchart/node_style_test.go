package flowchart

import (
	"reflect"
	"testing"
)

func TestNewNodeStyle(t *testing.T) {
	tests := []struct {
		name          string
		wantNodeStyle *NodeStyle
	}{
		{
			name: "Create new node style with default values",
			wantNodeStyle: &NodeStyle{
				StrokeWidth: 1,
				StrokeDash:  "0",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewNodeStyle()
			if !reflect.DeepEqual(got, tt.wantNodeStyle) {
				t.Errorf("NewNodeStyle() = %v, want %v", got, tt.wantNodeStyle)
			}
		})
	}
}

func TestNodeStyle_String(t *testing.T) {
	tests := []struct {
		name      string
		nodeStyle *NodeStyle
		setup     func(*NodeStyle)
		wantStr   string
	}{
		{
			name:      "Empty style",
			nodeStyle: &NodeStyle{},
			wantStr:   "",
		},
		{
			name:      "Default style",
			nodeStyle: NewNodeStyle(),
			wantStr:   "stroke-width:1,stroke-dasharray:0",
		},
		{
			name:      "Style with color only",
			nodeStyle: &NodeStyle{},
			setup: func(ns *NodeStyle) {
				ns.Color = "#ff0000"
			},
			wantStr: "color:#ff0000",
		},
		{
			name:      "Style with fill only",
			nodeStyle: &NodeStyle{},
			setup: func(ns *NodeStyle) {
				ns.Fill = "#00ff00"
			},
			wantStr: "fill:#00ff00",
		},
		{
			name:      "Style with stroke only",
			nodeStyle: &NodeStyle{},
			setup: func(ns *NodeStyle) {
				ns.Stroke = "#0000ff"
			},
			wantStr: "stroke:#0000ff",
		},
		{
			name:      "Style with all properties",
			nodeStyle: &NodeStyle{},
			setup: func(ns *NodeStyle) {
				ns.Color = "#ff0000"
				ns.Fill = "#00ff00"
				ns.Stroke = "#0000ff"
				ns.StrokeWidth = 2
				ns.StrokeDash = "5"
			},
			wantStr: "color:#ff0000,fill:#00ff00,stroke:#0000ff,stroke-width:2,stroke-dasharray:5",
		},
		{
			name:      "Style with zero stroke width",
			nodeStyle: &NodeStyle{},
			setup: func(ns *NodeStyle) {
				ns.StrokeWidth = 0
			},
			wantStr: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup(tt.nodeStyle)
			}

			got := tt.nodeStyle.String()
			if got != tt.wantStr {
				t.Errorf("NodeStyle.String() = %v, want %v", got, tt.wantStr)
			}
		})
	}
}
