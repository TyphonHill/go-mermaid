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

// Namespace represents a container for grouping related classes
type Namespace struct {
	Name     string
	Classes  []*Class
	Children []*Namespace
}

// NewNamespace creates a namespace with the given name
func NewNamespace(name string) *Namespace {
	return &Namespace{
		Name:     name,
		Classes:  make([]*Class, 0),
		Children: make([]*Namespace, 0),
	}
}

// AddClass adds a class to this namespace
func (n *Namespace) AddClass(class *Class) {
	n.Classes = append(n.Classes, class)
}

// AddNamespace creates and adds a nested namespace
func (n *Namespace) AddNamespace(name string) *Namespace {
	namespace := NewNamespace(name)
	n.Children = append(n.Children, namespace)
	return namespace
}

// String returns the Mermaid syntax representation of this namespace
func (n *Namespace) String(curIndentation string) string {
	var sb strings.Builder

	if len(n.Classes) > 0 {
		sb.WriteString(fmt.Sprintf(string(baseNamespaceStartString), n.Name))

		for _, class := range n.Classes {
			sb.WriteString(class.String(curIndentation + "\t%s"))
		}

		sb.WriteString(baseNamespaceEndString)
	}

	return sb.String()
}
