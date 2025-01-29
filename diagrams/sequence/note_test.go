package sequence

import (
	"reflect"
	"testing"
)

func TestNewNote(t *testing.T) {
	actor1 := NewActor("user1", "User 1", ActorParticipant)
	actor2 := NewActor("user2", "User 2", ActorParticipant)

	tests := []struct {
		name     string
		position NotePosition
		text     string
		actors   []*Actor
		want     *Note
	}{
		{
			name:     "Create left note for single actor",
			position: NoteLeft,
			text:     "This is a left note",
			actors:   []*Actor{actor1},
			want: &Note{
				Position: NoteLeft,
				Text:     "This is a left note",
				Actors:   []*Actor{actor1},
			},
		},
		{
			name:     "Create right note for single actor",
			position: NoteRight,
			text:     "This is a right note",
			actors:   []*Actor{actor1},
			want: &Note{
				Position: NoteRight,
				Text:     "This is a right note",
				Actors:   []*Actor{actor1},
			},
		},
		{
			name:     "Create over note for multiple actors",
			position: NoteOver,
			text:     "This is an over note",
			actors:   []*Actor{actor1, actor2},
			want: &Note{
				Position: NoteOver,
				Text:     "This is an over note",
				Actors:   []*Actor{actor1, actor2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewNote(tt.position, tt.text, tt.actors...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNote() = %v, want %v", got, tt.want)
			}
		})
	}
}
