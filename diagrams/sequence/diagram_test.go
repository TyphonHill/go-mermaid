package sequence

import (
	"reflect"
	"strings"
	"testing"
)

func TestNewDiagram(t *testing.T) {
	tests := []struct {
		name string
		want *Diagram
	}{
		{
			name: "Create new diagram with default settings",
			want: &Diagram{
				Actors:     make([]*Actor, 0),
				Messages:   make([]*Message, 0),
				autonumber: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDiagram()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDiagram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiagram_String(t *testing.T) {
	actor1 := NewActor("user", "User", ActorParticipant)
	actor2 := NewActor("system", "System", ActorActor)

	tests := []struct {
		name     string
		diagram  *Diagram
		setup    func(*Diagram)
		wantStr  string
		contains []string
	}{
		{
			name:    "Empty diagram without fence",
			diagram: NewDiagram(),
			wantStr: "sequenceDiagram\n",
		},
		{
			name:    "Empty diagram with fence",
			diagram: NewDiagram(),
			setup: func(d *Diagram) {
				d.EnableMarkdownFence()
			},
			wantStr: "```mermaid\nsequenceDiagram\n\n```\n",
		},
		{
			name: "Diagram with title and fence",
			setup: func(d *Diagram) {
				d.EnableMarkdownFence()
				d.Title = "Test Sequence"
			},
			diagram: NewDiagram(),
			contains: []string{
				"```mermaid\n",
				"---\ntitle: Test Sequence\n---\n",
				"sequenceDiagram\n",
				"```\n",
			},
		},
		{
			name:    "Diagram with actors and messages with fence",
			diagram: NewDiagram(),
			setup: func(d *Diagram) {
				d.EnableMarkdownFence()
				d.Title = "Interaction Flow"
				d.AddActor("user", "User", ActorParticipant)
				d.AddActor("system", "System", ActorActor)
				d.AddMessage(actor1, actor2, MessageSolid, "Request")
				d.AddMessage(actor2, actor1, MessageResponse, "Response")
			},
			contains: []string{
				"```mermaid\n",
				"participant user as User",
				"actor system as System",
				"user-->system: Request",
				"system->>user: Response",
				"```\n",
			},
		},
		{
			name:    "Diagram with autonumbering and fence",
			diagram: NewDiagram(),
			setup: func(d *Diagram) {
				d.EnableMarkdownFence()
				d.EnableAutoNumber()
				d.AddActor("user", "User", ActorParticipant)
				d.AddActor("system", "System", ActorActor)
				d.AddMessage(actor1, actor2, MessageSolid, "Request")
			},
			contains: []string{
				"```mermaid\n",
				"autonumber",
				"participant user as User",
				"actor system as System",
				"user-->system: Request",
				"```\n",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup(tt.diagram)
			}

			got := tt.diagram.String()

			if tt.wantStr != "" && got != tt.wantStr {
				t.Errorf("Diagram.String() = %v, want %v", got, tt.wantStr)
			}

			for _, want := range tt.contains {
				if !strings.Contains(got, want) {
					t.Errorf("Diagram.String() output missing expected content: %v", want)
				}
			}

			// Verify fence markers appear together or not at all
			hasFenceStart := strings.Contains(got, "```mermaid\n")
			hasFenceEnd := strings.Contains(got, "```\n")
			if hasFenceStart != hasFenceEnd {
				t.Error("Markdown fence markers are not properly paired")
			}

			// Verify fence markers match markdownFence setting
			if tt.diagram.IsMarkdownFenceEnabled() != hasFenceStart {
				t.Error("Markdown fence presence doesn't match markdownFence setting")
			}
		})
	}
}

func TestDiagram_CreateAndDestroyActor(t *testing.T) {
	creator := NewActor("system", "System", ActorParticipant)

	tests := []struct {
		name      string
		diagram   *Diagram
		creator   *Actor
		id        string
		actorName string
		actorType ActorType
	}{
		{
			name:      "Create and destroy actor",
			diagram:   NewDiagram(),
			creator:   creator,
			id:        "temp",
			actorName: "Temporary Actor",
			actorType: ActorParticipant,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Add creator to diagram
			tt.diagram.AddActor(tt.creator.ID, tt.creator.Name, tt.creator.Type)

			// Test CreateActor
			got := tt.diagram.CreateActor(tt.creator, tt.id, tt.actorName, tt.actorType)

			if got.ID != tt.id || got.Name != tt.actorName || got.Type != tt.actorType {
				t.Errorf("CreateActor() = %v, want actor with id=%v, name=%v, type=%v",
					got, tt.id, tt.actorName, tt.actorType)
			}

			// Verify creation message was added
			if len(tt.diagram.Messages) != 1 || tt.diagram.Messages[0].Type != MessageCreate {
				t.Errorf("Creation message not added correctly")
			}

			// Test DestroyActor
			tt.diagram.DestroyActor(got)

			// Verify destruction message was added
			if len(tt.diagram.Messages) != 2 || tt.diagram.Messages[1].Type != MessageDestroy {
				t.Errorf("Destruction message not added correctly")
			}
		})
	}
}

func TestDiagram_AddNote(t *testing.T) {
	actor1 := NewActor("user", "User", ActorParticipant)
	actor2 := NewActor("system", "System", ActorParticipant)

	tests := []struct {
		name     string
		diagram  *Diagram
		position NotePosition
		text     string
		actors   []*Actor
		want     *Note
	}{
		{
			name:     "Add left note",
			diagram:  NewDiagram(),
			position: NoteLeft,
			text:     "Left note",
			actors:   []*Actor{actor1},
			want: &Note{
				Position: NoteLeft,
				Text:     "Left note",
				Actors:   []*Actor{actor1},
			},
		},
		{
			name:     "Add over note for multiple actors",
			diagram:  NewDiagram(),
			position: NoteOver,
			text:     "Over note",
			actors:   []*Actor{actor1, actor2},
			want: &Note{
				Position: NoteOver,
				Text:     "Over note",
				Actors:   []*Actor{actor1, actor2},
			},
		},
		{
			name:     "Add note with no actors",
			diagram:  NewDiagram(),
			position: NoteLeft,
			text:     "Invalid note",
			actors:   []*Actor{},
			want:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Add actors to diagram
			for _, actor := range tt.actors {
				tt.diagram.AddActor(actor.ID, actor.Name, actor.Type)
			}

			got := tt.diagram.AddNote(tt.position, tt.text, tt.actors...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddNote() = %v, want %v", got, tt.want)
			}

			// If note was added successfully, verify it was converted to a message
			if got != nil {
				if len(tt.diagram.Messages) != 1 {
					t.Errorf("Note was not converted to message")
				}
				noteMsg := tt.diagram.Messages[0]
				if noteMsg.Note != got {
					t.Errorf("Message does not contain correct note reference")
				}
			}
		})
	}
}

func TestDiagram_AddMessage(t *testing.T) {
	actor1 := NewActor("user", "User", ActorParticipant)
	actor2 := NewActor("system", "System", ActorParticipant)

	tests := []struct {
		name    string
		diagram *Diagram
		from    *Actor
		to      *Actor
		msgType MessageType
		text    string
		want    *Message
	}{
		{
			name:    "Add solid message",
			diagram: NewDiagram(),
			from:    actor1,
			to:      actor2,
			msgType: MessageSolid,
			text:    "Request",
			want: &Message{
				From:   actor1,
				To:     actor2,
				Type:   MessageSolid,
				Text:   "Request",
				Nested: make([]*Message, 0),
			},
		},
		{
			name:    "Add async message",
			diagram: NewDiagram(),
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
			// Add actors to diagram
			tt.diagram.AddActor(tt.from.ID, tt.from.Name, tt.from.Type)
			tt.diagram.AddActor(tt.to.ID, tt.to.Name, tt.to.Type)

			got := tt.diagram.AddMessage(tt.from, tt.to, tt.msgType, tt.text)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddMessage() = %v, want %v", got, tt.want)
			}

			if len(tt.diagram.Messages) != 1 || !reflect.DeepEqual(tt.diagram.Messages[0], got) {
				t.Errorf("Message not added to diagram correctly")
			}
		})
	}
}

func TestMessage_String(t *testing.T) {
	actor1 := NewActor("user", "User", ActorParticipant)
	actor2 := NewActor("system", "System", ActorParticipant)

	tests := []struct {
		name     string
		message  *Message
		indent   string
		wantText string
	}{
		{
			name: "Creation message with text",
			message: &Message{
				From: actor1,
				To:   actor2,
				Type: MessageCreate,
				Text: "creates",
			},
			indent:   "",
			wantText: "\tcreate user-->system: creates\n",
		},
		{
			name: "Creation message without text",
			message: &Message{
				From: actor1,
				To:   actor2,
				Type: MessageCreate,
				Text: "",
			},
			indent:   "",
			wantText: "", // Should not generate any output for creation without text
		},
		{
			name: "Destruction message",
			message: &Message{
				From: nil,
				To:   actor2,
				Type: MessageDestroy,
			},
			indent:   "",
			wantText: "\tdestroy system\n",
		},
		{
			name: "Regular message",
			message: &Message{
				From: actor1,
				To:   actor2,
				Type: MessageSolid,
				Text: "regular message",
			},
			indent:   "",
			wantText: "\tuser-->system: regular message\n",
		},
		{
			name: "Message with nested messages",
			message: &Message{
				From: actor1,
				To:   actor2,
				Type: MessageSolid,
				Text: "parent",
				Nested: []*Message{
					{
						From: actor2,
						To:   actor1,
						Type: MessageResponse,
						Text: "nested",
					},
				},
			},
			indent:   "",
			wantText: "\tuser-->system: parent\n\t\tsystem->>user: nested\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.message.String(tt.indent)
			if got != tt.wantText {
				t.Errorf("Message.String() = %q, want %q", got, tt.wantText)
			}
		})
	}
}

func TestNote_String(t *testing.T) {
	actor1 := NewActor("user", "User", ActorParticipant)
	actor2 := NewActor("system", "System", ActorParticipant)

	tests := []struct {
		name     string
		note     *Note
		indent   string
		wantText string
	}{
		{
			name: "Left note for single actor",
			note: &Note{
				Position: NoteLeft,
				Text:     "This is a left note",
				Actors:   []*Actor{actor1},
			},
			indent:   "",
			wantText: "\tNote left of user: This is a left note\n",
		},
		{
			name: "Right note for single actor",
			note: &Note{
				Position: NoteRight,
				Text:     "This is a right note",
				Actors:   []*Actor{actor1},
			},
			indent:   "",
			wantText: "\tNote right of user: This is a right note\n",
		},
		{
			name: "Over note for multiple actors",
			note: &Note{
				Position: NoteOver,
				Text:     "This is an over note",
				Actors:   []*Actor{actor1, actor2},
			},
			indent:   "",
			wantText: "\tNote over user,system: This is an over note\n",
		},
		{
			name: "Invalid note with no actors",
			note: &Note{
				Position: NoteLeft,
				Text:     "This note has no actors",
				Actors:   []*Actor{},
			},
			indent:   "",
			wantText: "",
		},
		{
			name: "Note with different indentation",
			note: &Note{
				Position: NoteLeft,
				Text:     "Indented note",
				Actors:   []*Actor{actor1},
			},
			indent:   "\t",
			wantText: "\t\tNote left of user: Indented note\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.note.String(tt.indent)
			if got != tt.wantText {
				t.Errorf("Note.String() = %q, want %q", got, tt.wantText)
			}
		})
	}
}

func TestDiagram_EnableAutoNumber(t *testing.T) {
	tests := []struct {
		name    string
		diagram *Diagram
	}{
		{
			name:    "Enable autonumbering",
			diagram: NewDiagram(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.diagram.EnableAutoNumber()
			if !tt.diagram.autonumber {
				t.Errorf("EnableAutoNumber() did not set autonumber to true")
			}

			// Verify it appears in the output
			output := tt.diagram.String()
			if !strings.Contains(output, "autonumber") {
				t.Errorf("Diagram string does not contain autonumber directive")
			}
		})
	}
}
