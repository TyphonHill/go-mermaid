package sequence_diagram

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
	MessageActivate   MessageType = "-->+" // Solid arrow with activation
	MessageDeactivate MessageType = "-->-" // Solid arrow with deactivation

	// Special message types for tracking creation/destruction
	messageCreate  MessageType = "create"
	messageDestroy MessageType = "destroy"
)

// Message represents a communication between actors in a sequence diagram.
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

// AddNestedMessage adds a nested message to the current message.
func (m *Message) AddNestedMessage(from, to *Actor, msgType MessageType, text string) *Message {
	nested := NewMessage(from, to, msgType, text)
	m.Nested = append(m.Nested, nested)
	return nested
}
