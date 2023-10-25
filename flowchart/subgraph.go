package flowchart

import (
	"fmt"
	"strings"
)

type subgraphDirection string

const (
	baseSubgraphString          string = "\tsubgraph %d [%s]\n"
	baseSubgraphDirectionString string = "\t\tdirection %s\n"
	baseSubgraphEndString       string = "\tend\n"
	baseSubgraphLinkString      string = "\t%s\n"
	baseSubgraphSubgraphString  string = "\t%s"
)

const (
	SubgraphDirectionNone        subgraphDirection = ""
	SubgraphDirectionTopToBottom subgraphDirection = "TB"
	SubgraphDirectionBottomUp    subgraphDirection = "BT"
	SubgraphDirectionRightLeft   subgraphDirection = "RL"
	SubgraphDirectionLeftRight   subgraphDirection = "LR"
)

type Subgraph struct {
	ID        uint64
	Title     string
	Direction subgraphDirection
	subgraphs []*Subgraph
	links     []*Link
}

func NewSubgraph(id uint64, title string) (newSubgraph *Subgraph) {
	newSubgraph = &Subgraph{
		ID:        id,
		Title:     title,
		Direction: SubgraphDirectionNone,
	}

	return
}

func (s *Subgraph) AddSubgraph(title string) (newSubgraph *Subgraph) {
	newSubgraph = NewSubgraph(NewID(), title)

	s.subgraphs = append(s.subgraphs, newSubgraph)

	return
}

func (s *Subgraph) AddLink(from *Node, to *Node) (newLink *Link) {
	newLink = NewLink(from, to)

	s.links = append(s.links, newLink)

	return
}

func (s *Subgraph) String(curIndentation string) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf(string(curIndentation), fmt.Sprintf(string(baseSubgraphString), s.ID, s.Title)))

	direction := ""
	if s.Direction != SubgraphDirectionNone {
		direction = fmt.Sprintf(string(curIndentation), fmt.Sprintf(string(baseSubgraphDirectionString), string(s.Direction)))
	}

	sb.WriteString(direction)

	for _, subgraph := range s.subgraphs {
		sb.WriteString(subgraph.String(fmt.Sprintf(string(baseSubgraphSubgraphString), string(curIndentation))))
	}

	for _, link := range s.links {
		sb.WriteString(fmt.Sprintf(string(curIndentation), fmt.Sprintf(string(baseSubgraphLinkString), link.String())))
	}

	sb.WriteString(fmt.Sprintf(string(curIndentation), baseSubgraphEndString))

	return sb.String()
}
