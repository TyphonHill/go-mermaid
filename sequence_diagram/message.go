package sequence_diagram

type MessageType string

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

type Message struct {
	From   *Actor
	To     *Actor
	Type   MessageType
	Text   string
	Nested []*Message
	Note   *Note // For note support
}

func NewMessage(from, to *Actor, msgType MessageType, text string) *Message {
	return &Message{
		From:   from,
		To:     to,
		Type:   msgType,
		Text:   text,
		Nested: make([]*Message, 0),
	}
}

func (m *Message) AddNestedMessage(from, to *Actor, msgType MessageType, text string) *Message {
	nested := NewMessage(from, to, msgType, text)
	m.Nested = append(m.Nested, nested)
	return nested
}
