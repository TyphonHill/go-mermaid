package flowchart

import (
	"reflect"
	"testing"
)

func TestNewNodeStyle(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)

	tests := []struct {
		name             string
		wantNewNodeStyle *NodeStyle
	}{
		{
			name: "Nominal test",
			wantNewNodeStyle: &NodeStyle{
				StrokeWidth: 1,
				StrokeDash:  "0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewNodeStyle := NewNodeStyle(); !reflect.DeepEqual(gotNewNodeStyle, tt.wantNewNodeStyle) {
				t.Errorf("NewNodeStyle() = %v, want %v", gotNewNodeStyle, tt.wantNewNodeStyle)
			}
		})
	}
}

func TestNodeStyle_String(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)

	tests := []struct {
		name      string
		nodeStyle *NodeStyle
		want      string
	}{
		{
			name: "Nominal test",
			nodeStyle: &NodeStyle{
				Color:       "red",
				Fill:        "blue",
				Stroke:      "purple",
				StrokeWidth: 1,
				StrokeDash:  "0",
			},
			want: "color:red,fill:blue,stroke:purple,stroke-width:1,stroke-dasharray:0",
		},
		{
			name: "Missing fields",
			nodeStyle: &NodeStyle{
				StrokeWidth: 1,
				StrokeDash:  "0",
			},
			want: "stroke-width:1,stroke-dasharray:0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.nodeStyle.String(); got != tt.want {
				t.Errorf("NodeStyle.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
