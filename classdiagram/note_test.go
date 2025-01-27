package classdiagram

import (
	"reflect"
	"testing"
)

func TestNewNote(t *testing.T) {
	class := NewClass("Test Class")
	type args struct {
		text  string
		class *Class
	}
	tests := []struct {
		name        string
		args        args
		wantNewNote *Note
	}{
		{
			name: "Nominal test with no class",
			args: args{
				text:  "This is a note",
				class: nil,
			},
			wantNewNote: &Note{
				Text:  "This is a note",
				Class: nil,
			},
		},
		{
			name: "Nominal test with class",
			args: args{
				text:  "This is a note for a class",
				class: class,
			},
			wantNewNote: &Note{
				Text:  "This is a note for a class",
				Class: class,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewNote := NewNote(tt.args.text, tt.args.class); !reflect.DeepEqual(gotNewNote, tt.wantNewNote) {
				t.Errorf("NewNote() = %v, want %v", gotNewNote, tt.wantNewNote)
			}
		})
	}
}

func TestNote_String(t *testing.T) {
	class := NewClass("Test Class")

	tests := []struct {
		name string
		note *Note
		want string
	}{
		{
			name: "Nominal test with no class",
			note: NewNote("This is a note", nil),
			want: "\tnote \"This is a note\"\n",
		},
		{
			name: "Nominal test with class",
			note: NewNote("This is a note for a class", class),
			want: "\tnote for Test Class \"This is a note for a class\"\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Note{
				Text:  tt.note.Text,
				Class: tt.note.Class,
			}
			if got := n.String(); got != tt.want {
				t.Errorf("Note.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
