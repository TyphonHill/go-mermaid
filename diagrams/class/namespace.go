package class

import (
	"fmt"
	"strings"
)

// Namespace constants for formatting the Mermaid syntax representation.
const (
	baseNamespaceStartString string = "\tnamespace %s{\n"
	baseNamespaceEndString   string = "\t}\n"
)

// Namespace represents a grouping of related classes in a class diagram.
// It contains a name and a collection of classes within that namespace.
type Namespace struct {
	Name    string
	classes []*Class
}

// NewNamespace creates a new Namespace with the given name.
// It initializes an empty collection of classes.
func NewNamespace(name string) (newNamespace *Namespace) {
	newNamespace = &Namespace{
		Name: name,
	}

	return
}

// AddClass creates and adds a new class to the namespace.
// It returns the newly created Class, allowing for further configuration.
func (n *Namespace) AddClass(name string) (newClass *Class) {
	newClass = NewClass(name)

	n.classes = append(n.classes, newClass)

	return
}

// String generates the Mermaid syntax representation of the namespace.
// It includes the namespace name and all classes within the namespace.
// If no classes exist, an empty string is returned.
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
