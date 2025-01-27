package classdiagram

import (
	"fmt"
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

type ClassDiagram struct {
	Title      string
	Direction  classDiagramDirection
	namespaces []*Namespace
	notes      []*Note
	classes    []*Class
	relations  []*Relation
}

func NewClassDiagram() (newClassDiagram *ClassDiagram) {
	newClassDiagram = &ClassDiagram{
		Direction: ClassDiagramDirectionTopToBottom,
	}

	return
}

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
		sb.WriteString(namespace.String())
	}

	for _, class := range cd.classes {
		sb.WriteString(class.String("%s"))
	}

	for _, relation := range cd.relations {
		sb.WriteString(relation.String())
	}

	return sb.String()
}

func (cd *ClassDiagram) AddNamespace(name string) (newNamespace *Namespace) {
	newNamespace = NewNamespace(name)

	cd.namespaces = append(cd.namespaces, newNamespace)

	return
}

func (cd *ClassDiagram) AddNote(text string, class *Class) {
	newNote := NewNote(text, class)

	cd.notes = append(cd.notes, newNote)
}

func (cd *ClassDiagram) AddClass(name string, namespace *Namespace) (newClass *Class) {
	if namespace == nil {
		newClass = NewClass(name)
		cd.classes = append(cd.classes, newClass)
	} else {
		newClass = namespace.AddClass(name)
	}

	return
}

func (cd *ClassDiagram) AddRelation(classA *Class, classB *Class) (newRelation *Relation) {
	newRelation = NewRelation(classA, classB)

	cd.relations = append(cd.relations, newRelation)

	return
}
