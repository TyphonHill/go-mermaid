package block

import (
	"reflect"
	"testing"
)

func TestNewLink(t *testing.T) {
	from := NewBlock("0", "From")
	to := NewBlock("1", "To")

	tests := []struct {
		name     string
		from     *Block
		to       *Block
		wantLink *Link
	}{
		{
			name: "Create new link between blocks",
			from: from,
			to:   to,
			wantLink: &Link{
				From: from,
				To:   to,
				Text: "",
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

func TestLink_SetText(t *testing.T) {
	from := NewBlock("0", "From")
	to := NewBlock("1", "To")
	link := NewLink(from, to)

	result := link.SetText("Test Link")

	if link.Text != "Test Link" {
		t.Errorf("SetText() = %v, want %v", link.Text, "Test Link")
	}

	if result != link {
		t.Error("SetText() should return link for chaining")
	}
}

func TestLink_String(t *testing.T) {
	tests := []struct {
		name    string
		link    *Link
		setup   func(*Link)
		wantStr string
	}{
		{
			name: "Basic link without text",
			link: NewLink(
				NewBlock("0", "From"),
				NewBlock("1", "To"),
			),
			wantStr: "\t0 --> 1\n",
		},
		{
			name: "Link with text",
			link: NewLink(
				NewBlock("2", "Start"),
				NewBlock("3", "End"),
			),
			setup: func(l *Link) {
				l.SetText("Test Link")
			},
			wantStr: "\t2 -- \"Test Link\" --> 3\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup(tt.link)
			}

			got := tt.link.String()
			if got != tt.wantStr {
				t.Errorf("Link.String() = %v, want %v", got, tt.wantStr)
			}
		})
	}
}
