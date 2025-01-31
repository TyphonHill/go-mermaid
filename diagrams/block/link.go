package block

import "fmt"

// Mermaid link syntax templates
const (
	tplLinkWithText = "\t%s -- \"%s\" --> %s\n"
	tplLink         = "\t%s --> %s\n"
)

// Link represents a connection between two blocks
type Link struct {
	From *Block
	To   *Block
	Text string
}

// NewLink creates a link between two blocks
func NewLink(from, to *Block) *Link {
	return &Link{
		From: from,
		To:   to,
	}
}

// SetText sets the text label for this link
func (l *Link) SetText(text string) *Link {
	l.Text = text
	return l
}

// String returns the Mermaid syntax representation of this link
func (l *Link) String() string {
	if l.Text != "" {
		return fmt.Sprintf(tplLinkWithText, l.From.ID, l.Text, l.To.ID)
	}
	return fmt.Sprintf(tplLink, l.From.ID, l.To.ID)
}
