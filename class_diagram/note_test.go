package classdiagram

import (
	"strings"
	"testing"
)

func TestNewNote(t *testing.T) {
	tests := []struct {
		name  string
		text  string
		class *Class
		want  *Note
	}{
		{
			name:  "Create note without class",
			text:  "Test Note",
			class: nil,
			want: &Note{
				Text:  "Test Note",
				Class: nil,
			},
		},
		{
			name:  "Create note with class",
			text:  "Class Note",
			class: NewClass("TestClass"),
			want: &Note{
				Text:  "Class Note",
				Class: NewClass("TestClass"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewNote(tt.text, tt.class)

			if got.Text != tt.text {
				t.Errorf("NewNote() Text = %v, want %v", got.Text, tt.text)
			}

			// Compare class references
			if (got.Class == nil) != (tt.class == nil) {
				t.Errorf("NewNote() Class = %v, want %v", got.Class, tt.class)
			}

			if got.Class != nil && tt.class != nil && got.Class.Name != tt.class.Name {
				t.Errorf("NewNote() Class Name = %v, want %v", got.Class.Name, tt.class.Name)
			}
		})
	}
}

func TestNote_String(t *testing.T) {
	tests := []struct {
		name     string
		note     *Note
		contains []string
	}{
		{
			name: "General diagram note",
			note: NewNote("This is a general note", nil),
			contains: []string{
				"note \"This is a general note\"",
			},
		},
		{
			name: "Note for a specific class",
			note: func() *Note {
				class := NewClass("TestClass")
				return NewNote("Note about TestClass", class)
			}(),
			contains: []string{
				"note for TestClass \"Note about TestClass\"",
			},
		},
		{
			name: "Note with special characters",
			note: NewNote("Note with \"quotes\" and special chars", nil),
			contains: []string{
				`note "Note with "quotes" and special chars"`,
			},
		},
		{
			name: "Multiline note",
			note: NewNote("First line\nSecond line", nil),
			contains: []string{
				"note \"First line\nSecond line\"",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := tt.note.String()

			for _, expectedContent := range tt.contains {
				if !strings.Contains(output, expectedContent) {
					t.Errorf("String() output missing expected content: %q, got %v", expectedContent, output)
				}
			}
		})
	}
}
