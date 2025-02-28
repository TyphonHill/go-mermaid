// Package block provides functionality for creating Mermaid block diagrams
package block

import (
	"fmt"
	"strings"

	"github.com/TyphonHill/go-mermaid/diagrams/utils"
	"github.com/TyphonHill/go-mermaid/diagrams/utils/basediagram"
)

// Global ID generator for the package
var idGenerator = utils.NewIDGenerator()

// Mermaid diagram syntax templates
const (
	baseDiagramType = "block-beta\n"
	tplDiagramCols  = basediagram.Indentation + "columns %d\n"
)

// Diagram represents a Mermaid block diagram
type Diagram struct {
	basediagram.BaseDiagram[BlockConfigurationProperties]
	Blocks  []*Block
	Links   []*Link
	Columns int
}

// NewDiagram creates a new block diagram
func NewDiagram() *Diagram {
	return &Diagram{
		BaseDiagram: basediagram.NewBaseDiagram(NewBlockConfigurationProperties()),
		Blocks:      make([]*Block, 0),
		Links:       make([]*Link, 0),
		Columns:     0,
	}
}

// SetColumns sets the number of columns in the diagram
func (d *Diagram) SetColumns(count int) *Diagram {
	d.Columns = count
	return d
}

// AddColumn increases the number of columns by one
func (d *Diagram) AddColumn() *Diagram {
	d.Columns += 1
	return d
}

// RemoveColumn decreases the number of columns by one
func (d *Diagram) RemoveColumn() *Diagram {
	d.Columns -= 1
	return d
}

// AddBlock creates and adds a new block to the diagram
func (d *Diagram) AddBlock(text string) *Block {
	block := NewBlock(idGenerator.NextID(), text)
	block.diagram = d
	d.Blocks = append(d.Blocks, block)
	return block
}

// AddSpace adds a space block of one column width
func (d *Diagram) AddSpace() {
	d.Blocks = append(d.Blocks, &Block{IsSpace: true})
}

// AddSpaceWithWidth adds a space block with specified width
func (d *Diagram) AddSpaceWithWidth(width int) {
	d.Blocks = append(d.Blocks, &Block{IsSpace: true, Width: width})
}

// AddLink creates a link between two blocks
func (d *Diagram) AddLink(from, to *Block) *Link {
	link := NewLink(from, to)
	d.Links = append(d.Links, link)
	return link
}

// String returns the Mermaid syntax representation of this diagram
func (d *Diagram) String() string {
	var sb strings.Builder

	sb.WriteString(baseDiagramType)

	if d.Columns > 0 {
		sb.WriteString(fmt.Sprintf(tplDiagramCols, d.Columns))
	}

	for _, block := range d.Blocks {
		sb.WriteString(block.String())
	}

	for _, link := range d.Links {
		sb.WriteString(link.String())
	}

	return d.BaseDiagram.String(sb.String())
}

// RenderToFile saves the diagram to a file at the specified path
func (d *Diagram) RenderToFile(path string) error {
	return utils.RenderToFile(path, d.String())
}
