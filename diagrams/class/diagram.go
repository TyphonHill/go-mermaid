package class

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type classDiagramDirection string

// List of possible Class Diagram directions.
// Reference: https://mermaid.js.org/syntax/classDiagram.html#setting-the-direction-of-the-diagram
const (
	ClassDiagramDirectionTopToBottom classDiagramDirection = "TB"
	ClassDiagramDirectionBottomUp    classDiagramDirection = "BT"
	ClassDiagramDirectionRightLeft   classDiagramDirection = "RL"
	ClassDiagramDirectionLeftRight   classDiagramDirection = "LR"
)

const (
	baseClassDiagramTitleString     string = "---\ntitle: %s\n---\n\n"
	baseClassDiagramString          string = "classDiagram\n"
	baseClassDiagramDirectionString string = "\tdirection %s\n"
)

// ClassDiagram represents a Mermaid class diagram with various diagram components
// such as classes, namespaces, relations, and notes.
type ClassDiagram struct {
	Title         string
	Direction     classDiagramDirection
	namespaces    []*Namespace
	notes         []*Note
	classes       []*Class
	relations     []*Relation
	markdownFence bool
}

// EnableMarkdownFence enables markdown code fencing for the diagram output,
// which adds ```mermaid markers when converting the diagram to a string.
func (cd *ClassDiagram) EnableMarkdownFence() {
	cd.markdownFence = true
}

// DisableMarkdownFence disables markdown code fencing for the diagram output,
// removing the ```mermaid markers when converting the diagram to a string.
func (cd *ClassDiagram) DisableMarkdownFence() {
	cd.markdownFence = false
}

// NewClassDiagram creates and returns a new ClassDiagram with default settings.
// The default direction is set to top-to-bottom.
func NewClassDiagram() (newClassDiagram *ClassDiagram) {
	newClassDiagram = &ClassDiagram{
		Direction: ClassDiagramDirectionTopToBottom,
	}

	return
}

// String generates the Mermaid syntax representation of the class diagram.
// It includes title, direction, notes, namespaces, classes, and relations.
func (cd *ClassDiagram) String() string {
	var sb strings.Builder

	if cd.markdownFence {
		sb.WriteString("```mermaid\n")
	}

	if len(cd.Title) > 0 {
		sb.WriteString(fmt.Sprintf(string(baseClassDiagramTitleString), cd.Title))
	}

	sb.WriteString(baseClassDiagramString)

	sb.WriteString(fmt.Sprintf(string(baseClassDiagramDirectionString), string(cd.Direction)))

	for _, note := range cd.notes {
		sb.WriteString(note.String())
	}

	for _, namespace := range cd.namespaces {
		sb.WriteString(namespace.String())
	}

	for _, class := range cd.classes {
		sb.WriteString(class.String("%s"))
	}

	for _, relation := range cd.relations {
		sb.WriteString(relation.String())
	}

	if cd.markdownFence {
		sb.WriteString("```\n")
	}

	return sb.String()
}

// RenderToFile saves the diagram to a file at the specified path.
// If the file extension is .md, markdown fencing is automatically enabled.
// It creates the directory if it does not exist and writes the diagram content.
func (cd *ClassDiagram) RenderToFile(path string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	originalFenceState := cd.markdownFence
	if strings.ToLower(filepath.Ext(path)) == ".md" {
		cd.EnableMarkdownFence()
	}

	content := cd.String()

	cd.markdownFence = originalFenceState

	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// AddNamespace creates and adds a new namespace to the class diagram.
// It returns the newly created Namespace.
func (cd *ClassDiagram) AddNamespace(name string) (newNamespace *Namespace) {
	newNamespace = NewNamespace(name)

	cd.namespaces = append(cd.namespaces, newNamespace)

	return
}

// AddNote creates and adds a new note to the class diagram.
// The note can be associated with a specific class or be a general diagram note.
func (cd *ClassDiagram) AddNote(text string, class *Class) {
	newNote := NewNote(text, class)

	cd.notes = append(cd.notes, newNote)
}

// AddClass creates and adds a new class to the class diagram.
// If namespace is nil, the class is added directly to the diagram.
// Returns the newly created Class.
func (cd *ClassDiagram) AddClass(name string, namespace *Namespace) (newClass *Class) {
	if namespace == nil {
		newClass = NewClass(name)
		cd.classes = append(cd.classes, newClass)
	} else {
		newClass = namespace.AddClass(name)
	}

	return
}

// AddRelation creates and adds a new relation between two classes.
// Returns the newly created Relation.
func (cd *ClassDiagram) AddRelation(classA *Class, classB *Class) (newRelation *Relation) {
	newRelation = NewRelation(classA, classB)

	cd.relations = append(cd.relations, newRelation)

	return
}
