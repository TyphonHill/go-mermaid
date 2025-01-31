package block

import (
	"fmt"
	"strings"

	"github.com/TyphonHill/go-mermaid/diagrams/utils"
)

// Diagram represents a Mermaid block diagram
type Diagram struct {
	utils.BaseDiagram
	Blocks      []*Block
	Links       []*Link
	Columns     int
	idGenerator utils.IDGenerator
}

// NewDiagram creates a new block diagram
func NewDiagram() *Diagram {
	return &Diagram{
		Blocks:      make([]*Block, 0),
		Links:       make([]*Link, 0),
		Columns:     0,
		idGenerator: utils.NewIDGenerator(),
	}
}

// SetColumns sets the number of columns in the diagram
func (d *Diagram) SetColumns(count int) *Diagram {
	d.Columns = count
	return d
}

// AddBlock creates and adds a new block to the diagram
func (d *Diagram) AddBlock(text string) *Block {
	block := NewBlock(d.idGenerator.NextID(), text)
	block.diagram = d
	d.Blocks = append(d.Blocks, block)
	return block
}

// AddSpace adds a space block of one column width
func (d *Diagram) AddSpace() {
	d.Blocks = append(d.Blocks, &Block{IsSpace: true})
}

// AddSpaceWithWidth adds a space block with specified column width
func (d *Diagram) AddSpaceWithWidth(width int) {
	d.Blocks = append(d.Blocks, &Block{IsSpace: true, Width: width})
}

// AddLink creates a link between two blocks
func (d *Diagram) AddLink(from, to *Block) *Link {
	link := NewLink(from, to)
	d.Links = append(d.Links, link)
	return link
}

// String generates the Mermaid syntax for the block diagram
func (d *Diagram) String() string {
	var sb strings.Builder

	sb.WriteString("block-beta\n")

	// Add columns if defined
	if d.Columns > 0 {
		sb.WriteString(fmt.Sprintf("\tcolumns %d\n", d.Columns))
	}

	// Add blocks
	for _, block := range d.Blocks {
		sb.WriteString(block.String())
		if !block.IsSpace && block.Style != "" {
			sb.WriteString(fmt.Sprintf("\tstyle %s %s\n", block.ID, block.Style))
		}
	}

	// Add links
	for _, link := range d.Links {
		sb.WriteString(link.String())
	}

	return d.WrapWithFence(sb.String())
}

// RenderToFile saves the diagram to a file at the specified path
func (d *Diagram) RenderToFile(path string) error {
	return utils.RenderToFile(path, d.String(), d.IsMarkdownFenceEnabled())
}
