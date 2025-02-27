package flowchart

import (
	"fmt"
	"strings"

	"github.com/TyphonHill/go-mermaid/diagrams/utils"
	"github.com/TyphonHill/go-mermaid/diagrams/utils/basediagram"
)

const (
	baseSubgraphString          string = basediagram.Indentation + "subgraph %s [%s]\n"
	baseSubgraphDirectionString string = basediagram.Indentation + "direction %s\n"
	baseSubgraphEndString       string = basediagram.Indentation + "end\n"
	baseSubgraphLinkString      string = basediagram.Indentation + "%s"
	baseSubgraphSubgraphString  string = basediagram.Indentation + "%s"
)

// List of possible Subgraph directions.
// Reference: https://mermaid.js.org/syntax/flowchart.html#direction
const (
	SubgraphDirectionNone        subgraphDirection = ""
	SubgraphDirectionTopToBottom subgraphDirection = "TB"
	SubgraphDirectionBottomUp    subgraphDirection = "BT"
	SubgraphDirectionRightLeft   subgraphDirection = "RL"
	SubgraphDirectionLeftRight   subgraphDirection = "LR"
)

type subgraphDirection string

type Subgraph struct {
	ID          string
	Title       string
	Direction   subgraphDirection
	subgraphs   []*Subgraph
	links       []*Link
	idGenerator utils.IDGenerator
}

// NewSubgraph creates a new Subgraph with the given ID and title,
// setting the default direction to none.
func NewSubgraph(id string, title string) (newSubgraph *Subgraph) {
	newSubgraph = &Subgraph{
		ID:        id,
		Title:     title,
		Direction: SubgraphDirectionNone,
	}

	return
}

// AddSubgraph adds a new Subgraph to the current Subgraph and returns the created subgraph.
func (s *Subgraph) AddSubgraph(title string) (newSubgraph *Subgraph) {
	if s.idGenerator == nil {
		s.idGenerator = utils.NewIDGenerator()
	}

	newSubgraph = NewSubgraph(s.idGenerator.NextID(), title)
	newSubgraph.idGenerator = s.idGenerator

	s.subgraphs = append(s.subgraphs, newSubgraph)

	return
}

// AddLink adds a new Link to the Subgraph and returns the created link.
func (s *Subgraph) AddLink(from *Node, to *Node) (newLink *Link) {
	newLink = NewLink(from, to)

	s.links = append(s.links, newLink)

	return
}

// String generates a Mermaid string representation of the Subgraph,
// including its subgraphs, direction, and links with the specified indentation.
func (s *Subgraph) String(curIndentation string) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf(string(curIndentation), fmt.Sprintf(string(baseSubgraphString), s.ID, s.Title)))

	direction := ""
	if s.Direction != SubgraphDirectionNone {
		direction = fmt.Sprintf(string(curIndentation), fmt.Sprintf(string(baseSubgraphDirectionString), string(s.Direction)))
	}

	sb.WriteString(direction)

	for _, subgraph := range s.subgraphs {
		nextIndentation := fmt.Sprintf(string(baseSubgraphSubgraphString), string(curIndentation))
		sb.WriteString(subgraph.String(nextIndentation))
	}

	for _, link := range s.links {
		sb.WriteString(fmt.Sprintf(string(curIndentation), fmt.Sprintf(string(baseSubgraphLinkString), link.String())))
	}

	sb.WriteString(fmt.Sprintf(string(curIndentation), baseSubgraphEndString))

	return sb.String()
}
