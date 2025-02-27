package flowchart

import (
	"reflect"
	"strings"
	"testing"
)

func TestNewLink(t *testing.T) {
	from := NewNode("1", "Start")
	to := NewNode("2", "End")

	tests := []struct {
		name     string
		from     *Node
		to       *Node
		wantLink *Link
	}{
		{
			name: "Create new link between nodes",
			from: from,
			to:   to,
			wantLink: &Link{
				From:   from,
				To:     to,
				Shape:  LinkShapeOpen,
				Head:   LinkArrowTypeArrow,
				Tail:   LinkArrowTypeNone,
				Length: 0,
			},
		},
		{
			name: "Create new link with same node",
			from: from,
			to:   from,
			wantLink: &Link{
				From:   from,
				To:     from,
				Shape:  LinkShapeOpen,
				Head:   LinkArrowTypeArrow,
				Tail:   LinkArrowTypeNone,
				Length: 0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewLink(tt.from, tt.to)
			if !reflect.DeepEqual(got, tt.wantLink) {
				t.Errorf("NewLink() = %v, want %v", got, tt.wantLink)
			}
		})
	}
}

func TestLink_String(t *testing.T) {
	from := NewNode("1", "Start")
	to := NewNode("2", "End")

	tests := []struct {
		name     string
		link     *Link
		setup    func(*Link)
		contains []string
	}{
		{
			name: "Basic link",
			link: &Link{
				From:   from,
				To:     to,
				Shape:  LinkShapeOpen,
				Head:   LinkArrowTypeArrow,
				Tail:   LinkArrowTypeNone,
				Length: 0,
			},
			contains: []string{
				"1 --> 2",
			},
		},
		{
			name: "Link with text",
			link: &Link{
				From:   from,
				To:     to,
				Shape:  LinkShapeOpen,
				Head:   LinkArrowTypeArrow,
				Tail:   LinkArrowTypeNone,
				Text:   "Connection",
				Length: 0,
			},
			contains: []string{
				"1 -->|Connection| 2",
			},
		},
		{
			name: "Link with length",
			link: &Link{
				From:   from,
				To:     to,
				Shape:  LinkShapeOpen,
				Head:   LinkArrowTypeArrow,
				Tail:   LinkArrowTypeNone,
				Length: 2,
			},
			contains: []string{
				"1 ----> 2",
			},
		},
		{
			name: "Link with different shapes",
			link: &Link{
				From: from,
				To:   to,
				Head: LinkArrowTypeArrow,
				Tail: LinkArrowTypeNone,
			},
			setup: func(l *Link) {
				shapes := []linkShape{
					LinkShapeOpen,
					LinkShapeDotted,
					LinkShapeThick,
					LinkShapeInvisible,
				}
				for _, shape := range shapes {
					l.Shape = shape
					got := l.String()
					if got == "" {
						t.Errorf("Link.String() with shape %v returned empty string", shape)
					}
				}
			},
		},
		{
			name: "Link with different arrow types",
			link: &Link{
				From:  from,
				To:    to,
				Shape: LinkShapeOpen,
			},
			setup: func(l *Link) {
				arrowTypes := []linkArrowType{
					LinkArrowTypeNone,
					LinkArrowTypeArrow,
					LinkArrowTypeLeftArrow,
					LinkArrowTypeBullet,
					LinkArrowTypeCross,
				}
				for _, head := range arrowTypes {
					for _, tail := range arrowTypes {
						l.Head = head
						l.Tail = tail
						got := l.String()
						if got == "" {
							t.Errorf("Link.String() with head %v and tail %v returned empty string", head, tail)
						}
					}
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup(tt.link)
			}

			for _, want := range tt.contains {
				got := tt.link.String()
				if !strings.Contains(got, want) {
					t.Errorf("String() missing expected content %q in:\n%s", want, got)
				}
			}
		})
	}
}
