package classdiagram

import (
	"fmt"
	"strings"
)

const (
	baseNamespaceStartString string = "\tnamespace %s{\n"
	baseNamespaceEndString   string = "\t}\n"
)

type Namespace struct {
	Name    string
	classes []*Class
}

func NewNamespace(name string) (newNamespace *Namespace) {
	newNamespace = &Namespace{
		Name: name,
	}

	return
}

func (n *Namespace) AddClass(name string) (newClass *Class) {
	newClass = NewClass(name)

	n.classes = append(n.classes, newClass)

	return
}

func (n *Namespace) String() string {
	var sb strings.Builder

	if len(n.classes) > 0 {
		sb.WriteString(fmt.Sprintf(string(baseNamespaceStartString), n.Name))

		for _, class := range n.classes {
			sb.WriteString(class.String("\t%s"))
		}

		sb.WriteString(baseNamespaceEndString)
	}

	return sb.String()
}
