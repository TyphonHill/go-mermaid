package block

import (
	"fmt"
	"strings"
)

type blockShape string

// List of possible Block shapes.
// Reference: https://mermaid.js.org/syntax/flowchart.html#node-shapes
const (
	BlockShapeDefault       blockShape = `["%s"]`
	BlockShapeRoundEdges    blockShape = `("%s")`
	BlockShapeStadium       blockShape = `(["%s"])`
	BlockShapeSubroutine    blockShape = `[["%s"]]`
	BlockShapeCylindrical   blockShape = `[("%s")]`
	BlockShapeCircle        blockShape = `(("%s"))`
	BlockShapeAsymmetric    blockShape = `>"%s"]`
	BlockShapeRhombus       blockShape = `{"%s"}`
	BlockShapeHexagon       blockShape = `{{"%s"}}`
	BlockShapeParallelogram blockShape = `[/"%s"/]`
	BlockShapeTrapezoid     blockShape = `[/"%s"\]`
	BlockShapeTrapezoidAlt  blockShape = `[\"%s"/]`
	BlockShapeDoubleCircle  blockShape = `((("%s")))`
)

// Block represents a block in the diagram
type Block struct {
	ID       string
	Text     string
	Style    string
	Shape    blockShape
	Children []*Block
	IsSpace  bool
	Width    int // Number of columns the block spans
	diagram  *Diagram
	// New fields for block arrows
	isArrow   bool
	direction []BlockArrowDirection
}

// NewBlock creates a new block
func NewBlock(id, text string) *Block {
	return &Block{
		ID:       id,
		Text:     text,
		Style:    "",
		Shape:    BlockShapeDefault,
		Children: make([]*Block, 0),
		Width:    1, // Default to 1 column
	}
}

// SetWidth sets the number of columns the block spans
func (b *Block) SetWidth(width int) *Block {
	b.Width = width
	return b
}

// SetStyle sets the block's CSS style properties
func (b *Block) SetStyle(style string) *Block {
	b.Style = style
	return b
}

// SetShape sets the block's shape
func (b *Block) SetShape(shape blockShape) *Block {
	b.Shape = shape
	return b
}

// AddBlock creates and adds a nested block
func (b *Block) AddBlock(text string) *Block {
	// Get the diagram reference to generate ID
	diagram := b.diagram
	block := NewBlock(diagram.idGenerator.NextID(), text)
	b.Children = append(b.Children, block)
	return block
}

// SetArrow sets the block as an arrow with the given directions
func (b *Block) SetArrow(directions ...BlockArrowDirection) *Block {
	b.isArrow = true
	b.direction = directions
	return b
}

// String generates the Mermaid syntax for the block
func (b *Block) String() string {
	var sb strings.Builder

	if b.IsSpace {
		sb.WriteString("\tspace\n")
		return sb.String()
	}

	if len(b.Children) > 0 {
		// Parent block with children
		sb.WriteString(fmt.Sprintf("\tblock:%s\n", b.ID))
		for _, child := range b.Children {
			if child.Text != "" {
				if child.isArrow {
					sb.WriteString(fmt.Sprintf("\t\t%s%s\n", child.ID, BlockArrowShape(child.Text, child.direction...)))
				} else {
					sb.WriteString(fmt.Sprintf("\t\t%s%s\n", child.ID, fmt.Sprintf(string(child.Shape), child.Text)))
				}
			} else {
				sb.WriteString(fmt.Sprintf("\t\t%s\n", child.ID))
			}
		}
		sb.WriteString("\tend\n")
	} else {
		// Single block
		if b.Text != "" {
			if b.isArrow {
				if b.Width > 1 {
					sb.WriteString(fmt.Sprintf("\t%s%s:%d\n", b.ID, BlockArrowShape(b.Text, b.direction...), b.Width))
				} else {
					sb.WriteString(fmt.Sprintf("\t%s%s\n", b.ID, BlockArrowShape(b.Text, b.direction...)))
				}
			} else {
				if b.Width > 1 {
					sb.WriteString(fmt.Sprintf("\t%s%s:%d\n", b.ID, fmt.Sprintf(string(b.Shape), b.Text), b.Width))
				} else {
					sb.WriteString(fmt.Sprintf("\t%s%s\n", b.ID, fmt.Sprintf(string(b.Shape), b.Text)))
				}
			}
		} else {
			if b.Width > 1 {
				sb.WriteString(fmt.Sprintf("\t%s:%d\n", b.ID, b.Width))
			} else {
				sb.WriteString(fmt.Sprintf("\t%s\n", b.ID))
			}
		}
	}

	// Add style if present
	if b.Style != "" {
		sb.WriteString(fmt.Sprintf("\tstyle %s %s\n", b.ID, b.Style))
	}

	return sb.String()
}
