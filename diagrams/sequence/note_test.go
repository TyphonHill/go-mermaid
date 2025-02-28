package sequence

import (
	"strings"
	"testing"
)

func TestNewNote(t *testing.T) {
	actor1 := &Actor{ID: "A"}
	actor2 := &Actor{ID: "B"}

	tests := []struct {
		name     string
		position NotePosition
		text     string
		actors   []*Actor
		want     *Note
	}{
		{
			name:     "Create left note",
			position: NoteLeft,
			text:     "Left note",
			actors:   []*Actor{actor1},
			want: &Note{
				Position: NoteLeft,
				Text:     "Left note",
				Actors:   []*Actor{actor1},
			},
		},
		{
			name:     "Create right note",
			position: NoteRight,
			text:     "Right note",
			actors:   []*Actor{actor1},
			want: &Note{
				Position: NoteRight,
				Text:     "Right note",
				Actors:   []*Actor{actor1},
			},
		},
		{
			name:     "Create over note with one actor",
			position: NoteOver,
			text:     "Over note",
			actors:   []*Actor{actor1},
			want: &Note{
				Position: NoteOver,
				Text:     "Over note",
				Actors:   []*Actor{actor1},
			},
		},
		{
			name:     "Create over note with two actors",
			position: NoteOver,
			text:     "Over both note",
			actors:   []*Actor{actor1, actor2},
			want: &Note{
				Position: NoteOver,
				Text:     "Over both note",
				Actors:   []*Actor{actor1, actor2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newNote(tt.position, tt.text, tt.actors...)
			if got.Position != tt.want.Position {
				t.Errorf("NewNote().Position = %v, want %v", got.Position, tt.want.Position)
			}
			if got.Text != tt.want.Text {
				t.Errorf("NewNote().Text = %v, want %v", got.Text, tt.want.Text)
			}
			if len(got.Actors) != len(tt.want.Actors) {
				t.Errorf("NewNote().Actors length = %v, want %v", len(got.Actors), len(tt.want.Actors))
			}
		})
	}
}

func TestNote_String(t *testing.T) {
	tests := []struct {
		name     string
		note     *Note
		contains []string
		empty    bool
	}{
		{
			name: "Left note",
			note: newNote(NoteLeft, "Left side note", &Actor{ID: "A"}),
			contains: []string{
				"Note left of A: Left side note",
			},
		},
		{
			name: "Right note",
			note: newNote(NoteRight, "Right side note", &Actor{ID: "B"}),
			contains: []string{
				"Note right of B: Right side note",
			},
		},
		{
			name: "Over note single actor",
			note: newNote(NoteOver, "Over note", &Actor{ID: "C"}),
			contains: []string{
				"Note over C: Over note",
			},
		},
		{
			name: "Over note multiple actors",
			note: newNote(NoteOver, "Over both note", &Actor{ID: "A"}, &Actor{ID: "B"}),
			contains: []string{
				"Note over A,B: Over both note",
			},
		},
		{
			name: "Note with indentation",
			note: newNote(NoteLeft, "Indented note", &Actor{ID: "A"}),
			contains: []string{
				"Note left of A: Indented note",
			},
		},
		{
			name:  "Note without actors",
			note:  newNote(NoteOver, "Empty note"),
			empty: true,
		},
		{
			name: "Left note with multiple actors",
			note: newNote(NoteLeft, "Left note", &Actor{ID: "A"}),
			contains: []string{
				"Note left of A: Left note",
			},
		},
		{
			name: "Right note with multiple actors",
			note: newNote(NoteRight, "Right note", &Actor{ID: "A"}),
			contains: []string{
				"Note right of A: Right note",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.note.String("")

			if tt.empty {
				if got != "" {
					t.Errorf("String() = %q, want empty string", got)
				}
				return
			}

			for _, want := range tt.contains {
				if !strings.Contains(got, want) {
					t.Errorf("String() missing expected content %q in:\n%s", want, got)
				}
			}
		})
	}
}
