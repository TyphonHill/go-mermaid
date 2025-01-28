// note.go
package sequence_diagram

type NotePosition string

const (
	NoteLeft  NotePosition = "left of"
	NoteRight NotePosition = "right of"
	NoteOver  NotePosition = "over"
)

type Note struct {
	Position NotePosition
	Text     string
	Actors   []*Actor // For single actor notes, only the first actor is used
}

// NewNote creates a new Note with the specified properties
func NewNote(position NotePosition, text string, actors ...*Actor) *Note {
	return &Note{
		Position: position,
		Text:     text,
		Actors:   actors,
	}
}
