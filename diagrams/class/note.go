package class

import (
	"fmt"
	"strings"
)

// Note constants for formatting the Mermaid syntax representation.
const (
	baseDiagramNoteString string = "\tnote \"%s\"\n"
	baseClassNoteString   string = "\tnote for %s \"%s\"\n"
)

// Note represents an annotation or comment in a class diagram.
// It can be either a general diagram note or a note associated with a specific class.
type Note struct {
	Text  string
	Class *Class
}

// NewNote creates a new Note with the given text and optional associated class.
// If no class is provided, the note is treated as a general diagram note.
func NewNote(text string, class *Class) (newNote *Note) {
	newNote = &Note{
		Text:  text,
		Class: class,
	}

	return
}

// String generates the Mermaid syntax representation of the note.
// If the note is associated with a class, it uses the class-specific note format.
// Otherwise, it uses the general diagram note format.
func (n *Note) String() string {
	var sb strings.Builder

	if n.Class == nil {
		sb.WriteString(fmt.Sprintf(string(baseDiagramNoteString), n.Text))
	} else {
		sb.WriteString(fmt.Sprintf(string(baseClassNoteString), n.Class.Name, n.Text))
	}

	return sb.String()
}
