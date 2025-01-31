package block

import "fmt"

// Link represents a connection between blocks
type Link struct {
	From *Block
	To   *Block
	Text string
}

// NewLink creates a new link between blocks
func NewLink(from, to *Block) *Link {
	return &Link{
		From: from,
		To:   to,
	}
}

// SetText sets the link's text
func (l *Link) SetText(text string) *Link {
	l.Text = text
	return l
}

// String generates the Mermaid syntax for the link
func (l *Link) String() string {
	if l.Text != "" {
		return fmt.Sprintf("\t%s -- \"%s\" --> %s\n", l.From.ID, l.Text, l.To.ID)
	}
	return fmt.Sprintf("\t%s --> %s\n", l.From.ID, l.To.ID)
}
