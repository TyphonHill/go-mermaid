package sequence

import (
	"fmt"
	"strings"
)

// MessageType represents the different types of messages in a sequence diagram.
type MessageType string

// Predefined message types for sequence diagrams.
const (
	// Regular message types
	MessageSolid      MessageType = "-->"   // Solid line
	MessageSolidArrow MessageType = "-->>"  // Solid line with arrow
	MessageDashed     MessageType = "-->"   // Dashed line (same as solid in Mermaid)
	MessageAsync      MessageType = "->>"   // Async message
	MessageDotted     MessageType = "-->>>" // Dotted line with arrow
	MessageResponse   MessageType = "->>"   // Response message (typically async)

	// Activation/Deactivation arrows
	MessageActivate   MessageType = "+" // Activation
	MessageDeactivate MessageType = "-" // Deactivation

	// Special message types for tracking creation/destruction
	MessageCreate  MessageType = "create"
	MessageDestroy MessageType = "destroy"
)

// Base string formats for message elements
const (
	baseMessage       string = "%s%s%s%s: %s\n" // indent, from, arrow, to, text
	baseMessageNoDesc string = "%s%s%s%s\n"     // indent, from, arrow, to
	baseCreate        string = "create "
	baseDestroy       string = "destroy %s\n"
	baseActivate      string = "activate %s\n"
	baseDeactivate    string = "deactivate %s\n"
)

// Message represents a message between actors in a sequence diagram.
type Message struct {
	From   *Actor
	To     *Actor
	Type   MessageType
	Text   string
	Nested []*Message
	Note   *Note
}

// NewMessage creates a new Message between two actors.
func NewMessage(from, to *Actor, msgType MessageType, text string) *Message {
	return &Message{
		From:   from,
		To:     to,
		Type:   msgType,
		Text:   text,
		Nested: make([]*Message, 0),
	}
}

// AddNestedMessage creates and adds a new nested message.
func (m *Message) AddNestedMessage(from, to *Actor, msgType MessageType, text string) *Message {
	nested := NewMessage(from, to, msgType, text)
	m.Nested = append(m.Nested, nested)
	return nested
}

// SetType sets the message type and returns the message for chaining.
func (m *Message) SetType(msgType MessageType) *Message {
	m.Type = msgType
	return m
}

// SetText sets the message text and returns the message for chaining.
func (m *Message) SetText(text string) *Message {
	m.Text = text
	return m
}

// String generates a Mermaid-formatted string representation of the message with custom indentation.
func (m *Message) String(curIndentation string) string {
	var sb strings.Builder

	// Handle notes first
	if m.Note != nil {
		return m.Note.String(curIndentation)
	}

	switch m.Type {
	case MessageCreate:
		if m.Text != "" {
			sb.WriteString(fmt.Sprintf("%s\t%s", curIndentation,
				fmt.Sprintf(baseMessage, baseCreate, m.From.ID, MessageSolid, m.To.ID, m.Text)))
		}
	case MessageDestroy:
		sb.WriteString(fmt.Sprintf("%s\t%s", curIndentation,
			fmt.Sprintf(baseDestroy, m.To.ID)))
	case MessageActivate:
		// Write the message and activation
		if m.Text != "" {
			sb.WriteString(fmt.Sprintf("%s\t%s", curIndentation,
				fmt.Sprintf(baseMessage, "", m.From.ID, "-->", m.To.ID, m.Text)))
		}
		sb.WriteString(fmt.Sprintf("%s\t%s", curIndentation,
			fmt.Sprintf(baseActivate, m.To.ID)))
	case MessageDeactivate:
		// Only deactivate the sender (From), not the receiver (To)
		if m.Text != "" {
			sb.WriteString(fmt.Sprintf("%s\t%s", curIndentation,
				fmt.Sprintf(baseMessage, "", m.From.ID, MessageSolid, m.To.ID, m.Text)))
		}
		sb.WriteString(fmt.Sprintf("%s\t%s", curIndentation,
			fmt.Sprintf(baseDeactivate, m.To.ID))) // Use To instead of From
	default:
		arrow := string(m.Type)
		if m.Text != "" {
			sb.WriteString(fmt.Sprintf("%s\t%s", curIndentation,
				fmt.Sprintf(baseMessage, "", m.From.ID, arrow, m.To.ID, m.Text)))
		} else {
			sb.WriteString(fmt.Sprintf("%s\t%s", curIndentation,
				fmt.Sprintf(baseMessageNoDesc, "", m.From.ID, arrow, m.To.ID)))
		}
	}

	// Handle nested messages
	if len(m.Nested) > 0 {
		nextIndentation := fmt.Sprintf("%s\t", curIndentation)
		for _, nested := range m.Nested {
			sb.WriteString(nested.String(nextIndentation))
		}
	}

	return sb.String()
}
