package sequence

import (
	"reflect"
	"strings"
	"testing"

	"github.com/TyphonHill/go-mermaid/diagrams/utils/basediagram"
)

func TestNewDiagram(t *testing.T) {
	tests := []struct {
		name string
		want *Diagram
	}{
		{
			name: "Create new diagram with default settings",
			want: &Diagram{
				BaseDiagram: basediagram.NewBaseDiagram(NewSequenceConfigurationProperties()),
				Actors:      make([]*Actor, 0),
				Messages:    make([]*Message, 0),
				autonumber:  false,
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
			name:    "Diagram with actors and messages",
			diagram: NewDiagram(),
			setup: func(d *Diagram) {
				d.SetTitle("Interaction Flow")
				d.AddActor("user", "User", ActorParticipant)
				d.AddActor("system", "System", ActorActor)
				d.AddMessage(actor1, actor2, MessageSolid, "Request")
				d.AddMessage(actor2, actor1, MessageResponse, "Response")
			},
			contains: []string{
				"participant user as User",
				"actor system as System",
				"user-->system: Request",
				"system->>user: Response",
			},
		},
		{
			name:    "Diagram with autonumbering",
			diagram: NewDiagram(),
			setup: func(d *Diagram) {
				d.EnableAutoNumber()
				d.AddActor("user", "User", ActorParticipant)
				d.AddActor("system", "System", ActorActor)
				d.AddMessage(actor1, actor2, MessageSolid, "Request")
			},
			contains: []string{
				"autonumber",
				"participant user as User",
				"actor system as System",
				"user-->system: Request",
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

func TestDiagram_AddNote(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(*Diagram) *Note
		contains []string
	}{
		{
			name: "Add left note for single actor",
			setup: func(d *Diagram) *Note {
				actor := d.AddActor("A", "Actor A", ActorParticipant)
				return d.AddNote(NoteLeft, "Left note", actor)
			},
			contains: []string{
				"participant A as Actor A",
				"Note left of A: Left note",
			},
		},
		{
			name: "Add right note for single actor",
			setup: func(d *Diagram) *Note {
				actor := d.AddActor("B", "Actor B", ActorParticipant)
				return d.AddNote(NoteRight, "Right note", actor)
			},
			contains: []string{
				"participant B as Actor B",
				"Note right of B: Right note",
			},
		},
		{
			name: "Add over note for single actor",
			setup: func(d *Diagram) *Note {
				actor := d.AddActor("C", "Actor C", ActorParticipant)
				return d.AddNote(NoteOver, "Over note", actor)
			},
			contains: []string{
				"participant C as Actor C",
				"Note over C: Over note",
			},
		},
		{
			name: "Add over note for multiple actors",
			setup: func(d *Diagram) *Note {
				actor1 := d.AddActor("A", "Actor A", ActorParticipant)
				actor2 := d.AddActor("B", "Actor B", ActorParticipant)
				return d.AddNote(NoteOver, "Over both note", actor1, actor2)
			},
			contains: []string{
				"participant A as Actor A",
				"participant B as Actor B",
				"Note over A,B: Over both note",
			},
		},
		{
			name: "Add note between messages",
			setup: func(d *Diagram) *Note {
				actor1 := d.AddActor("A", "Actor A", ActorParticipant)
				actor2 := d.AddActor("B", "Actor B", ActorParticipant)
				d.AddMessage(actor1, actor2, MessageSolid, "First message")
				note := d.AddNote(NoteOver, "Note between", actor1, actor2)
				d.AddMessage(actor2, actor1, MessageSolid, "Second message")
				return note
			},
			contains: []string{
				"participant A as Actor A",
				"participant B as Actor B",
				"A-->B: First message",
				"Note over A,B: Note between",
				"B-->A: Second message",
			},
		},
		{
			name: "Add note without actors",
			setup: func(d *Diagram) *Note {
				return d.AddNote(NoteOver, "Empty note")
			},
			contains: []string{
				"sequenceDiagram",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diagram := NewDiagram()
			note := tt.setup(diagram)

			// Verify note was added to messages
			if len(diagram.Messages) == 0 && note != nil {
				t.Error("Note was not added to diagram messages")
			}

			// Verify diagram string output
			got := diagram.String()
			for _, want := range tt.contains {
				if !strings.Contains(got, want) {
					t.Errorf("String() missing expected content %q in:\n%s", want, got)
				}
			}
		})
	}
}
