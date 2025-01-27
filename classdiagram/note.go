package classdiagram

import (
	"fmt"
	"strings"
)

const (
	baseDiagramNoteString string = "\tnote \"%s\"\n"
	baseClassNoteString   string = "\tnote for %s \"%s\"\n"
)

type Note struct {
	Text  string
	Class *Class
}

func NewNote(text string, class *Class) (newNote *Note) {
	newNote = &Note{
		Text:  text,
		Class: class,
	}

	return
}

func (n *Note) String() string {
	var sb strings.Builder

	if n.Class == nil {
		sb.WriteString(fmt.Sprintf(string(baseDiagramNoteString), n.Text))
	} else {
		sb.WriteString(fmt.Sprintf(string(baseClassNoteString), n.Class.Name, n.Text))
	}

	return sb.String()
}
