package userjourney

import (
	"reflect"
	"strings"
	"testing"
)

func TestNewSection(t *testing.T) {
	tests := []struct {
		name  string
		title string
		want  *Section
	}{
		{
			name:  "Create new section",
			title: "Test Section",
			want: &Section{
				Title: "Test Section",
				Tasks: make([]*Task, 0),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSection(tt.title)
			if got.Title != tt.want.Title {
				t.Errorf("NewSection().Title = %v, want %v", got.Title, tt.want.Title)
			}
			if len(got.Tasks) != 0 {
				t.Errorf("NewSection().Tasks length = %v, want 0", len(got.Tasks))
			}
		})
	}
}

func TestSection_String(t *testing.T) {
	tests := []struct {
		name     string
		section  *Section
		setup    func(*Section)
		contains []string
	}{
		{
			name:    "Empty section",
			section: NewSection("Empty Section"),
			contains: []string{
				"section Empty Section",
			},
		},
		{
			name:    "Section with single task without participants",
			section: NewSection("Basic Section"),
			setup: func(s *Section) {
				s.AddTask("Task 1", 3)
			},
			contains: []string{
				"section Basic Section",
				"Task 1: 3",
			},
		},
		{
			name:    "Section with single task and participant",
			section: NewSection("Single Participant"),
			setup: func(s *Section) {
				s.AddTask("Task 1", 4, "User")
			},
			contains: []string{
				"section Single Participant",
				"Task 1: 4: User",
			},
		},
		{
			name:    "Section with task and multiple participants",
			section: NewSection("Multiple Participants"),
			setup: func(s *Section) {
				s.AddTask("Task 1", 5, "User", "Admin", "System")
			},
			contains: []string{
				"section Multiple Participants",
				"Task 1: 5: User,Admin,System",
			},
		},
		{
			name:    "Section with multiple tasks and participants",
			section: NewSection("Complex Section"),
			setup: func(s *Section) {
				s.AddTask("Task 1", 2, "User")
				s.AddTask("Task 2", 4, "User", "Admin")
				s.AddTask("Task 3", 3)
			},
			contains: []string{
				"section Complex Section",
				"Task 1: 2: User",
				"Task 2: 4: User,Admin",
				"Task 3: 3",
			},
		},
		{
			name:    "Section with score clamping",
			section: NewSection("Score Limits"),
			setup: func(s *Section) {
				s.AddTask("Low Score", -1, "User")
				s.AddTask("High Score", 10, "User")
			},
			contains: []string{
				"section Score Limits",
				"Low Score: 1: User",
				"High Score: 5: User",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup(tt.section)
			}

			got := tt.section.String()
			for _, want := range tt.contains {
				if !strings.Contains(got, want) {
					t.Errorf("String() missing expected content %q in:\n%s", want, got)
				}
			}
		})
	}
}

func TestSection_AddTask(t *testing.T) {
	tests := []struct {
		name         string
		title        string
		score        int
		participants []string
		want         *Task
	}{
		{
			name:  "Add task without participants",
			title: "Task 1",
			score: 3,
			want: &Task{
				Title: "Task 1",
				Score: 3,
			},
		},
		{
			name:         "Add task with single participant",
			title:        "Task 2",
			score:        4,
			participants: []string{"User"},
			want: &Task{
				Title:        "Task 2",
				Score:        4,
				Participants: []string{"User"},
			},
		},
		{
			name:         "Add task with multiple participants",
			title:        "Task 3",
			score:        5,
			participants: []string{"User", "Admin", "System"},
			want: &Task{
				Title:        "Task 3",
				Score:        5,
				Participants: []string{"User", "Admin", "System"},
			},
		},
		{
			name:         "Add task with score below minimum",
			title:        "Task 4",
			score:        -1,
			participants: []string{"User"},
			want: &Task{
				Title:        "Task 4",
				Score:        1, // Should clamp to 1
				Participants: []string{"User"},
			},
		},
		{
			name:         "Add task with score above maximum",
			title:        "Task 5",
			score:        10,
			participants: []string{"User"},
			want: &Task{
				Title:        "Task 5",
				Score:        5, // Should clamp to 5
				Participants: []string{"User"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			section := NewSection("Test Section")
			got := section.AddTask(tt.title, tt.score, tt.participants...)

			if got.Title != tt.want.Title {
				t.Errorf("AddTask().Title = %v, want %v", got.Title, tt.want.Title)
			}
			if got.Score != tt.want.Score {
				t.Errorf("AddTask().Score = %v, want %v", got.Score, tt.want.Score)
			}
			if !reflect.DeepEqual(got.Participants, tt.want.Participants) {
				t.Errorf("AddTask().Participants = %v, want %v", got.Participants, tt.want.Participants)
			}

			if len(section.Tasks) != 1 || section.Tasks[0] != got {
				t.Error("Task was not added to section correctly")
			}
		})
	}
}
