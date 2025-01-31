// Package class provides functionality for creating Mermaid class diagrams
package class

import (
	"fmt"
	"strings"

	"github.com/TyphonHill/go-mermaid/diagrams/utils"
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
	utils.BaseDiagram
	Direction  classDiagramDirection
	namespaces []*Namespace
	notes      []*Note
	classes    []*Class
	relations  []*Relation
}

// SetDirection sets the diagram direction and returns the diagram for chaining
func (cd *ClassDiagram) SetDirection(direction classDiagramDirection) *ClassDiagram {
	cd.Direction = direction
	return cd
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

	if len(cd.Title) > 0 {
		sb.WriteString(fmt.Sprintf(string(baseClassDiagramTitleString), cd.Title))
	}

	sb.WriteString(baseClassDiagramString)

	sb.WriteString(fmt.Sprintf(string(baseClassDiagramDirectionString), string(cd.Direction)))

	for _, note := range cd.notes {
		sb.WriteString(note.String())
	}

	for _, namespace := range cd.namespaces {
		sb.WriteString(namespace.String(""))
	}

	for _, class := range cd.classes {
		sb.WriteString(class.String("%s"))
	}

	for _, relation := range cd.relations {
		sb.WriteString(relation.String())
	}

	return cd.WrapWithFence(sb.String())
}

// RenderToFile saves the diagram to a file at the specified path.
func (cd *ClassDiagram) RenderToFile(path string) error {
	return utils.RenderToFile(path, cd.String(), cd.IsMarkdownFenceEnabled())
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
	newClass = NewClass(name)
	if namespace == nil {
		cd.classes = append(cd.classes, newClass)
	} else {
		namespace.AddClass(newClass)
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
