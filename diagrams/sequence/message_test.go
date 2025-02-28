package sequence

import (
	"reflect"
	"strings"
	"testing"
)

func TestNewMessage(t *testing.T) {
	actor1 := NewActor("user", "User", ActorParticipant)
	actor2 := NewActor("system", "System", ActorParticipant)

	tests := []struct {
		name    string
		from    *Actor
		to      *Actor
		msgType MessageType
		text    string
		want    *Message
	}{
		{
			name:    "Create solid message",
			from:    actor1,
			to:      actor2,
			msgType: MessageSolid,
			text:    "Request data",
			want: &Message{
				From:   actor1,
				To:     actor2,
				Type:   MessageSolid,
				Text:   "Request data",
				Nested: make([]*Message, 0),
			},
		},
		{
			name:    "Create async message",
			from:    actor1,
			to:      actor2,
			msgType: MessageAsync,
			text:    "Background task",
			want: &Message{
				From:   actor1,
				To:     actor2,
				Type:   MessageAsync,
				Text:   "Background task",
				Nested: make([]*Message, 0),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewMessage(tt.from, tt.to, tt.msgType, tt.text)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_String(t *testing.T) {
	tests := []struct {
		name     string
		message  *Message
		setup    func(*Message)
		indent   string
		contains []string
	}{
		{
			name: "Message with note",
			message: &Message{
				Note: newNote(NoteOver, "This is a note", NewActor("A", "A", ActorParticipant), NewActor("B", "B", ActorParticipant)),
			},
			indent: "",
			contains: []string{
				"Note over A,B: This is a note",
			},
		},
		{
			name: "Message without text",
			message: NewMessage(
				&Actor{ID: "A"},
				&Actor{ID: "B"},
				MessageSolid,
				"",
			),
			indent: "",
			contains: []string{
				"A-->B",
			},
		},
		{
			name: "Create message with text",
			message: NewMessage(
				&Actor{ID: "A"},
				&Actor{ID: "B"},
				MessageCreate,
				"Creates B",
			),
			indent: "",
			contains: []string{
				"create A-->B: Creates B",
			},
		},
		{
			name: "Destroy message",
			message: NewMessage(
				&Actor{ID: "A"},
				&Actor{ID: "B"},
				MessageDestroy,
				"",
			),
			indent: "",
			contains: []string{
				"destroy B",
			},
		},
		{
			name: "Activate message with text",
			message: NewMessage(
				&Actor{ID: "A"},
				&Actor{ID: "B"},
				MessageActivate,
				"Start process",
			),
			indent: "",
			contains: []string{
				"A-->B: Start process",
				"activate B",
			},
		},
		{
			name: "Deactivate message with text",
			message: NewMessage(
				&Actor{ID: "A"},
				&Actor{ID: "B"},
				MessageDeactivate,
				"End process",
			),
			indent: "",
			contains: []string{
				"A-->B: End process",
				"deactivate B",
			},
		},
		{
			name: "Message with nested messages",
			message: NewMessage(
				&Actor{ID: "A"},
				&Actor{ID: "B"},
				MessageSolid,
				"Parent",
			),
			setup: func(m *Message) {
				nested := NewMessage(
					&Actor{ID: "B"},
					&Actor{ID: "C"},
					MessageSolid,
					"Child",
				)
				m.Nested = append(m.Nested, nested)
			},
			indent: "",
			contains: []string{
				"A-->B: Parent",
				"\tB-->C: Child",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup(tt.message)
			}

			got := tt.message.String(tt.indent)
			for _, want := range tt.contains {
				if !strings.Contains(got, want) {
					t.Errorf("String() missing expected content %q in:\n%s", want, got)
				}
			}
		})
	}
}

func TestMessage_SetType(t *testing.T) {
	tests := []struct {
		name    string
		msgType MessageType
		want    MessageType
	}{
		{"Set solid type", MessageSolid, MessageSolid},
		{"Set async type", MessageAsync, MessageAsync},
		{"Set dotted type", MessageDotted, MessageDotted},
		{"Set create type", MessageCreate, MessageCreate},
		{"Set destroy type", MessageDestroy, MessageDestroy},
		{"Set activate type", MessageActivate, MessageActivate},
		{"Set deactivate type", MessageDeactivate, MessageDeactivate},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := NewMessage(&Actor{ID: "A"}, &Actor{ID: "B"}, MessageSolid, "")
			result := msg.SetType(tt.msgType)

			if msg.Type != tt.want {
				t.Errorf("SetType() = %v, want %v", msg.Type, tt.want)
			}

			if result != msg {
				t.Error("SetType() should return message for chaining")
			}
		})
	}
}

func TestMessage_SetText(t *testing.T) {
	tests := []struct {
		name string
		text string
		want string
	}{
		{"Set empty text", "", ""},
		{"Set simple text", "Hello", "Hello"},
		{"Set text with spaces", "Hello World", "Hello World"},
		{"Set text with special chars", "Hello: World!", "Hello: World!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := NewMessage(&Actor{ID: "A"}, &Actor{ID: "B"}, MessageSolid, "")
			result := msg.SetText(tt.text)

			if msg.Text != tt.want {
				t.Errorf("SetText() = %v, want %v", msg.Text, tt.want)
			}

			if result != msg {
				t.Error("SetText() should return message for chaining")
			}
		})
	}
}

func TestMessage_AddNestedMessage(t *testing.T) {
	tests := []struct {
		name    string
		parent  *Message
		from    *Actor
		to      *Actor
		msgType MessageType
		text    string
	}{
		{
			name: "Add nested solid message",
			parent: NewMessage(
				&Actor{ID: "A"},
				&Actor{ID: "B"},
				MessageSolid,
				"Parent",
			),
			from:    &Actor{ID: "B"},
			to:      &Actor{ID: "C"},
			msgType: MessageSolid,
			text:    "Child",
		},
		{
			name: "Add nested async message",
			parent: NewMessage(
				&Actor{ID: "A"},
				&Actor{ID: "B"},
				MessageSolid,
				"Parent",
			),
			from:    &Actor{ID: "B"},
			to:      &Actor{ID: "C"},
			msgType: MessageAsync,
			text:    "Async Child",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.parent.AddNestedMessage(tt.from, tt.to, tt.msgType, tt.text)
			want := NewMessage(tt.from, tt.to, tt.msgType, tt.text)

			if !reflect.DeepEqual(got, want) {
				t.Errorf("AddNestedMessage() = %v, want %v", got, want)
			}

			if len(tt.parent.Nested) != 1 || !reflect.DeepEqual(tt.parent.Nested[0], got) {
				t.Error("Parent's Nested messages not updated correctly")
			}
		})
	}
}
