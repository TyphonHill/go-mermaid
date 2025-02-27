package block

import (
	"fmt"
	"strings"

	"github.com/TyphonHill/go-mermaid/diagrams/utils/basediagram"
)

// Mermaid block syntax templates
const (
	tplSpace        = basediagram.Indentation + "space\n"
	tplBlockStart   = basediagram.Indentation + "block:%s:%d\n"
	tplBlockSimple  = basediagram.Indentation + "block:%s\n"
	tplColumns      = basediagram.Indentation + "columns %d\n"
	tplChildBlock   = basediagram.Indentation + "%s%s\n"
	tplChildSimple  = basediagram.Indentation + "%s\n"
	tplBlockEnd     = basediagram.Indentation + "end\n"
	tplBlockWidth   = basediagram.Indentation + "%s%s:%d\n"
	tplBlockNoWidth = basediagram.Indentation + "%s%s\n"
	tplBlockID      = basediagram.Indentation + "%s\n"
	tplStyle        = basediagram.Indentation + "style %s %s\n"
)

// BlockShape defines the visual appearance of a block
type blockShape string

// Available block shapes for use in diagrams
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

// Block represents a node in a block diagram
type Block struct {
	ID        string
	Text      string
	Style     string
	Shape     blockShape
	Children  []*Block
	IsSpace   bool
	Width     int
	diagram   *Diagram
	isArrow   bool
	direction []BlockArrowDirection
	columns   int
}

// NewBlock creates a block with the given ID and text
func NewBlock(id, text string) *Block {
	return &Block{
		ID:       id,
		Text:     text,
		Shape:    BlockShapeDefault,
		Children: make([]*Block, 0),
		Width:    1,
	}
}

// SetWidth sets the number of columns this block spans
func (b *Block) SetWidth(width int) *Block {
	b.Width = width
	return b
}

// SetStyle sets CSS style properties for this block
func (b *Block) SetStyle(style string) *Block {
	b.Style = style
	return b
}

// SetShape sets the visual shape of this block
func (b *Block) SetShape(shape blockShape) *Block {
	b.Shape = shape
	return b
}

// AddBlock creates a nested block with the given text
func (b *Block) AddBlock(text string) *Block {
	block := NewBlock(idGenerator.NextID(), text)
	b.Children = append(b.Children, block)
	return block
}

// SetArrow configures this block as an arrow with the given directions
func (b *Block) SetArrow(directions ...BlockArrowDirection) *Block {
	b.isArrow = true
	b.direction = directions
	return b
}

// SetColumns sets the number of columns for this block's children
func (b *Block) SetColumns(count int) *Block {
	b.columns = count
	return b
}

// AddColumn increases the number of columns for this block's children by one
func (b *Block) AddColumn() *Block {
	b.columns += 1
	return b
}

// RemoveColumn decreases the number of columns for this block's children by one
func (b *Block) RemoveColumn() *Block {
	b.columns -= 1
	return b
}

// String returns the Mermaid syntax representation of this block
func (b *Block) String() string {
	var sb strings.Builder

	if b.IsSpace {
		sb.WriteString(tplSpace)
		return sb.String()
	}

	if len(b.Children) > 0 {
		// Parent block with children
		if b.Width > 0 {
			sb.WriteString(fmt.Sprintf(tplBlockStart, b.ID, b.Width))
		} else {
			sb.WriteString(fmt.Sprintf(tplBlockSimple, b.ID))
		}
		if b.columns > 0 {
			sb.WriteString(fmt.Sprintf(tplColumns, b.columns))
		}
		for _, child := range b.Children {
			if child.Text != "" {
				if child.isArrow {
					sb.WriteString(fmt.Sprintf(tplChildBlock, child.ID, BlockArrowShape(child.Text, child.direction...)))
				} else {
					sb.WriteString(fmt.Sprintf(tplChildBlock, child.ID, fmt.Sprintf(string(child.Shape), child.Text)))
				}
			} else {
				sb.WriteString(fmt.Sprintf(tplChildSimple, child.ID))
			}
		}
		sb.WriteString(tplBlockEnd)
	} else {
		// Single block
		if b.Text != "" {
			if b.isArrow {
				if b.Width > 1 {
					sb.WriteString(fmt.Sprintf(tplBlockWidth, b.ID, BlockArrowShape(b.Text, b.direction...), b.Width))
				} else {
					sb.WriteString(fmt.Sprintf(tplBlockNoWidth, b.ID, BlockArrowShape(b.Text, b.direction...)))
				}
			} else {
				if b.Width > 1 {
					sb.WriteString(fmt.Sprintf(tplBlockWidth, b.ID, fmt.Sprintf(string(b.Shape), b.Text), b.Width))
				} else {
					sb.WriteString(fmt.Sprintf(tplBlockNoWidth, b.ID, fmt.Sprintf(string(b.Shape), b.Text)))
				}
			}
		} else {
			if b.Width > 1 {
				sb.WriteString(fmt.Sprintf(tplBlockWidth, b.ID, "", b.Width))
			} else {
				sb.WriteString(fmt.Sprintf(tplBlockID, b.ID))
			}
		}
	}

	if b.Style != "" {
		sb.WriteString(fmt.Sprintf(tplStyle, b.ID, b.Style))
	}

	return sb.String()
}
