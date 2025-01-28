package sequence_diagram

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
