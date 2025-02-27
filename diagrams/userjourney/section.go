package userjourney

import (
	"fmt"
	"strings"

	"github.com/TyphonHill/go-mermaid/diagrams/utils/basediagram"
)

// Base string formats for user journey sections and tasks
const (
	baseSectionTitle   string = basediagram.Indentation + "section %s\n"
	baseTaskWithPartic string = basediagram.Indentation + basediagram.Indentation + "%s: %d: %s\n"
	baseTaskNoPartic   string = basediagram.Indentation + basediagram.Indentation + "%s: %d\n"
)

// Section represents a section in the user journey
type Section struct {
	Title string
	Tasks []*Task
}

// Task represents a task in a section
type Task struct {
	Title        string
	Score        int      // Score must be between 1-5
	Participants []string // Optional list of participants
}

// NewSection creates a new section
func NewSection(title string) *Section {
	return &Section{
		Title: title,
		Tasks: make([]*Task, 0),
	}
}

// AddTask adds a new task to the section
func (s *Section) AddTask(title string, score int, participants ...string) *Task {
	if score < 1 {
		score = 1
	}
	if score > 5 {
		score = 5
	}

	task := &Task{
		Title:        title,
		Score:        score,
		Participants: participants,
	}
	s.Tasks = append(s.Tasks, task)
	return task
}

// String generates the Mermaid syntax for the section
func (s *Section) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf(baseSectionTitle, s.Title))
	for _, task := range s.Tasks {
		if len(task.Participants) > 0 {
			sb.WriteString(fmt.Sprintf(baseTaskWithPartic,
				task.Title,
				task.Score,
				strings.Join(task.Participants, ",")))
		} else {
			sb.WriteString(fmt.Sprintf(baseTaskNoPartic, task.Title, task.Score))
		}
	}

	return sb.String()
}
