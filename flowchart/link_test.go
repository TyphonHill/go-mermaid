package flowchart

import (
	"reflect"
	"testing"
)

func TestNewLink(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)

	type args struct {
		from *Node
		to   *Node
	}
	tests := []struct {
		name        string
		args        args
		wantNewLink *Link
	}{
		{
			name: "Nominal test",
			args: args{
				from: &Node{ID: 123},
				to:   &Node{ID: 456},
			},
			wantNewLink: &Link{
				From:  &Node{ID: 123},
				To:    &Node{ID: 456},
				Shape: LinkShapeOpen,
				Head:  LinkArrowTypeArrow,
				Tail:  LinkArrowTypeNone,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewLink := NewLink(tt.args.from, tt.args.to); !reflect.DeepEqual(gotNewLink, tt.wantNewLink) {
				t.Errorf("NewLink() = %v, want %v", gotNewLink, tt.wantNewLink)
			}
		})
	}
}

func TestLink_String(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)

	tests := []struct {
		name string
		link *Link
		want string
	}{
		{
			name: "Nominal test",
			link: &Link{
				Shape:  LinkShapeOpen,
				Head:   LinkArrowTypeArrow,
				Tail:   LinkArrowTypeNone,
				Text:   "",
				From:   &Node{ID: 123},
				To:     &Node{ID: 456},
				Length: 0,
			},
			want: "\t123 --> 456",
		},
		{
			name: "Link with text",
			link: &Link{
				Shape:  LinkShapeOpen,
				Head:   LinkArrowTypeArrow,
				Tail:   LinkArrowTypeNone,
				Text:   "This is a test",
				From:   &Node{ID: 123},
				To:     &Node{ID: 456},
				Length: 0,
			},
			want: "\t123 -->|This is a test| 456",
		},
		{
			name: "Longer link",
			link: &Link{
				Shape:  LinkShapeOpen,
				Head:   LinkArrowTypeArrow,
				Tail:   LinkArrowTypeNone,
				Text:   "",
				From:   &Node{ID: 123},
				To:     &Node{ID: 456},
				Length: 10,
			},
			want: "\t123 ------------> 456",
		},
		{
			name: "Negative length",
			link: &Link{
				Shape:  LinkShapeOpen,
				Head:   LinkArrowTypeArrow,
				Tail:   LinkArrowTypeNone,
				Text:   "",
				From:   &Node{ID: 123},
				To:     &Node{ID: 456},
				Length: -10,
			},
			want: "\t123 --> 456",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.link.String(); got != tt.want {
				t.Errorf("Link.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
