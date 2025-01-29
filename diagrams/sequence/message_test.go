package sequence

import (
	"reflect"
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

func TestMessage_AddNestedMessage(t *testing.T) {
	actor1 := NewActor("user", "User", ActorParticipant)
	actor2 := NewActor("system", "System", ActorParticipant)
	parent := NewMessage(actor1, actor2, MessageSolid, "Parent message")

	tests := []struct {
		name    string
		parent  *Message
		from    *Actor
		to      *Actor
		msgType MessageType
		text    string
		want    *Message
	}{
		{
			name:    "Add nested message",
			parent:  parent,
			from:    actor2,
			to:      actor1,
			msgType: MessageResponse,
			text:    "Nested response",
			want: &Message{
				From:   actor2,
				To:     actor1,
				Type:   MessageResponse,
				Text:   "Nested response",
				Nested: make([]*Message, 0),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.parent.AddNestedMessage(tt.from, tt.to, tt.msgType, tt.text)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddNestedMessage() = %v, want %v", got, tt.want)
			}
			if len(tt.parent.Nested) != 1 || !reflect.DeepEqual(tt.parent.Nested[0], got) {
				t.Errorf("Parent's Nested messages not updated correctly")
			}
		})
	}
}

func TestMessage_SetText(t *testing.T) {
	actor1 := NewActor("A1", "Actor1", ActorParticipant)
	actor2 := NewActor("A2", "Actor2", ActorParticipant)
	msg := NewMessage(actor1, actor2, MessageSolid, "Initial")

	result := msg.SetText("Updated Text")

	if msg.Text != "Updated Text" {
		t.Errorf("SetText() = %v, want %v", msg.Text, "Updated Text")
	}

	if result != msg {
		t.Error("SetText() should return message for chaining")
	}
}

func TestMessage_SetType(t *testing.T) {
	actor1 := NewActor("A1", "Actor1", ActorParticipant)
	actor2 := NewActor("A2", "Actor2", ActorParticipant)
	msg := NewMessage(actor1, actor2, MessageSolid, "Test")

	tests := []struct {
		name     string
		msgType  MessageType
		expected MessageType
	}{
		{"Solid", MessageSolid, MessageSolid},
		{"Dashed", MessageDashed, MessageDashed},
		{"Async", MessageAsync, MessageAsync},
		{"Dotted", MessageDotted, MessageDotted},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := msg.SetType(tt.msgType)

			if msg.Type != tt.expected {
				t.Errorf("SetType() = %v, want %v", msg.Type, tt.expected)
			}

			if result != msg {
				t.Error("SetType() should return message for chaining")
			}
		})
	}
}

func TestMessage_SetType_ActivateDeactivate(t *testing.T) {
	actor1 := NewActor("A1", "Actor1", ActorParticipant)
	actor2 := NewActor("A2", "Actor2", ActorParticipant)
	msg := NewMessage(actor1, actor2, MessageSolid, "Test")

	tests := []struct {
		name     string
		msgType  MessageType
		expected MessageType
		want     string
	}{
		{
			name:     "Activate",
			msgType:  MessageActivate,
			expected: MessageActivate,
			want:     "\tA1-->A2: Test\n\tactivate A2\n",
		},
		{
			name:     "Deactivate",
			msgType:  MessageDeactivate,
			expected: MessageDeactivate,
			want:     "\tA1-->A2: Test\n\tdeactivate A2\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := msg.SetType(tt.msgType)

			if msg.Type != tt.expected {
				t.Errorf("SetType() = %v, want %v", msg.Type, tt.expected)
			}

			if result != msg {
				t.Error("SetType() should return message for chaining")
			}

			if output := msg.String(""); output != tt.want {
				t.Errorf("String() for %v = %q, want %q", tt.msgType, output, tt.want)
			}
		})
	}
}
