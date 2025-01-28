package sequencediagram

// NotePosition represents the positioning of a note in a sequence diagram.
type NotePosition string

// Predefined note positioning options.
const (
	NoteLeft  NotePosition = "left of"
	NoteRight NotePosition = "right of"
	NoteOver  NotePosition = "over"
)

// Note represents an annotation or comment in a sequence diagram.
type Note struct {
	Position NotePosition
	Text     string
	Actors   []*Actor
}

// NewNote creates a new Note with the specified properties.
func NewNote(position NotePosition, text string, actors ...*Actor) *Note {
	return &Note{
		Position: position,
		Text:     text,
		Actors:   actors,
	}
}
