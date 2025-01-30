package userjourney

import (
	"strings"
	"testing"
)

func TestNewSection(t *testing.T) {
	title := "Test Section"
	section := NewSection(title)

	if section.Title != title {
		t.Errorf("NewSection().Title = %v, want %v", section.Title, title)
	}
	if len(section.Tasks) != 0 {
		t.Error("NewSection() should create empty tasks slice")
	}
}

func TestSection_AddTask(t *testing.T) {
	tests := []struct {
		name      string
		title     string
		score     int
		wantScore int
	}{
		{
			name:      "Normal score",
			title:     "Task 1",
			score:     3,
			wantScore: 3,
		},
		{
			name:      "Score too low",
			title:     "Task 2",
			score:     0,
			wantScore: 1,
		},
		{
			name:      "Score too high",
			title:     "Task 3",
			score:     6,
			wantScore: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			section := NewSection("Test Section")
			task := section.AddTask(tt.title, tt.score)

			if len(section.Tasks) != 1 {
				t.Error("AddTask() should add task to section")
			}
			if task.Title != tt.title {
				t.Errorf("AddTask().Title = %v, want %v", task.Title, tt.title)
			}
			if task.Score != tt.wantScore {
				t.Errorf("AddTask().Score = %v, want %v", task.Score, tt.wantScore)
			}
		})
	}
}

func TestSection_String(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() *Section
		contains []string
	}{
		{
			name: "Empty section",
			setup: func() *Section {
				return NewSection("Empty")
			},
			contains: []string{
				"section Empty",
			},
		},
		{
			name: "Section with one task",
			setup: func() *Section {
				s := NewSection("Shopping")
				s.AddTask("Buy milk", 3)
				return s
			},
			contains: []string{
				"section Shopping",
				"Buy milk: 3",
			},
		},
		{
			name: "Section with multiple tasks",
			setup: func() *Section {
				s := NewSection("Website")
				s.AddTask("Visit homepage", 5)
				s.AddTask("Search product", 2)
				s.AddTask("Add to cart", 4)
				return s
			},
			contains: []string{
				"section Website",
				"Visit homepage: 5",
				"Search product: 2",
				"Add to cart: 4",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			section := tt.setup()
			result := section.String()

			for _, want := range tt.contains {
				if !strings.Contains(result, want) {
					t.Errorf("String() missing expected content %q in:\n%s", want, result)
				}
			}
		})
	}
}
