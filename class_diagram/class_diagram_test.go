package class_diagram

import (
	"reflect"
	"testing"
)

func TestNewClassDiagram(t *testing.T) {
	tests := []struct {
		name                string
		wantNewClassDiagram *ClassDiagram
	}{
		{
			name: "Nominal test",
			wantNewClassDiagram: &ClassDiagram{
				Direction: ClassDiagramDirectionTopToBottom,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewClassDiagram := NewClassDiagram(); !reflect.DeepEqual(gotNewClassDiagram, tt.wantNewClassDiagram) {
				t.Errorf("NewClassDiagram() = %v, want %v", gotNewClassDiagram, tt.wantNewClassDiagram)
			}
		})
	}
}

func TestClassDiagram_String(t *testing.T) {
	diagram := NewClassDiagram()
	diagram.Title = "TestClassDiagram"
	diagram.Direction = ClassDiagramDirectionLeftRight
	diagram.AddNote("TestNote", nil)

	cl1 := diagram.AddClass("Class1", nil)

	ns := diagram.AddNamespace("Namespace1")
	cl2 := diagram.AddClass("Class2", ns)
	cl3 := diagram.AddClass("Class3", ns)
	diagram.AddRelation(cl1, cl2)
	diagram.AddRelation(cl2, cl3)

	tests := []struct {
		name         string
		classDiagram *ClassDiagram
		want         string
	}{
		{
			name:         "Nominal test",
			classDiagram: NewClassDiagram(),
			want: `classDiagram
	direction TB
`,
		},
		{
			name:         "Test with complex diagram",
			classDiagram: diagram,
			want: `---
title: TestClassDiagram
---

classDiagram
	direction LR
	note "TestNote"
	namespace Namespace1{
		class Class2{
		}
		class Class3{
		}
	}
	class Class1{
	}
	Class1 -- Class2
	Class2 -- Class3
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cd := &ClassDiagram{
				Title:      tt.classDiagram.Title,
				Direction:  tt.classDiagram.Direction,
				namespaces: tt.classDiagram.namespaces,
				notes:      tt.classDiagram.notes,
				classes:    tt.classDiagram.classes,
				relations:  tt.classDiagram.relations,
			}
			if got := cd.String(); got != tt.want {
				t.Errorf("ClassDiagram.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
