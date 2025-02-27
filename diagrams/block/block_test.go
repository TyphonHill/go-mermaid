package block

import (
	"reflect"
	"strings"
	"testing"
)

func TestNewBlock(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		text      string
		wantBlock *Block
	}{
		{
			name: "Create new block with text",
			id:   "0",
			text: "Test Block",
			wantBlock: &Block{
				ID:       "0",
				Text:     "Test Block",
				Shape:    BlockShapeDefault,
				Style:    "",
				Children: make([]*Block, 0),
				Width:    1,
			},
		},
		{
			name: "Create new block without text",
			id:   "1",
			text: "",
			wantBlock: &Block{
				ID:       "1",
				Text:     "",
				Shape:    BlockShapeDefault,
				Style:    "",
				Children: make([]*Block, 0),
				Width:    1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idGenerator = idGenerator.Reset()

			got := NewBlock(tt.id, tt.text)
			if !reflect.DeepEqual(got, tt.wantBlock) {
				t.Errorf("NewBlock() = %v, want %v", got, tt.wantBlock)
			}
		})
	}
}

func TestBlock_SetWidth(t *testing.T) {
	block := NewBlock("0", "Test")
	result := block.SetWidth(2)

	if block.Width != 2 {
		t.Errorf("SetWidth() = %v, want %v", block.Width, 2)
	}

	if result != block {
		t.Error("SetWidth() should return block for chaining")
	}
}

func TestBlock_SetStyle(t *testing.T) {
	block := NewBlock("0", "Test")
	style := "fill:#f9f9f9,stroke:#333"
	result := block.SetStyle(style)

	if block.Style != style {
		t.Errorf("SetStyle() = %v, want %v", block.Style, style)
	}

	if result != block {
		t.Error("SetStyle() should return block for chaining")
	}
}

func TestBlock_String(t *testing.T) {
	tests := []struct {
		name     string
		block    *Block
		setup    func(*Block)
		contains []string
	}{
		{
			name:  "Empty block",
			block: NewBlock("0", ""),
			contains: []string{
				"0",
			},
		},
		{
			name:  "Block with text",
			block: NewBlock("1", "Test"),
			contains: []string{
				"1[\"Test\"]",
			},
		},
		{
			name:  "Block with width",
			block: NewBlock("2", "Test"),
			setup: func(b *Block) {
				b.SetWidth(2)
			},
			contains: []string{
				"2[\"Test\"]",
				":2",
			},
		},
		{
			name:  "Block with style",
			block: NewBlock("3", "Test"),
			setup: func(b *Block) {
				b.SetStyle("fill:#f9f9f9")
			},
			contains: []string{
				"3[\"Test\"]",
				"style 3 fill:#f9f9f9",
			},
		},
		{
			name:  "Parent block with children",
			block: NewBlock("4", "Parent"),
			setup: func(b *Block) {
				b.diagram = &Diagram{}
				b.AddBlock("Child 1")
				b.AddBlock("Child 2")
			},
			contains: []string{
				"block:4:1",
				"0[\"Child 1\"]",
				"1[\"Child 2\"]",
				"end",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idGenerator = idGenerator.Reset()

			if tt.setup != nil {
				tt.setup(tt.block)
			}

			got := tt.block.String()
			for _, want := range tt.contains {
				if !strings.Contains(got, want) {
					t.Errorf("String() missing expected content %q in:\n%s", want, got)
				}
			}
		})
	}
}

func TestBlock_SetShape(t *testing.T) {
	tests := []struct {
		name     string
		shape    blockShape
		text     string
		contains []string
	}{
		{
			name:  "Default shape",
			shape: BlockShapeDefault,
			text:  "Test",
			contains: []string{
				"0[\"Test\"]",
			},
		},
		{
			name:  "Round edges shape",
			shape: BlockShapeRoundEdges,
			text:  "Test",
			contains: []string{
				"0(\"Test\")",
			},
		},
		{
			name:  "Stadium shape",
			shape: BlockShapeStadium,
			text:  "Test",
			contains: []string{
				"0([\"Test\"])",
			},
		},
		{
			name:  "Subroutine shape",
			shape: BlockShapeSubroutine,
			text:  "Test",
			contains: []string{
				"0[[\"Test\"]]",
			},
		},
		{
			name:  "Cylindrical shape",
			shape: BlockShapeCylindrical,
			text:  "Test",
			contains: []string{
				"0[(\"Test\")]",
			},
		},
		{
			name:  "Circle shape",
			shape: BlockShapeCircle,
			text:  "Test",
			contains: []string{
				"0((\"Test\"))",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idGenerator = idGenerator.Reset()

			block := NewBlock("0", tt.text)
			block.SetShape(tt.shape)
			got := block.String()
			for _, want := range tt.contains {
				if !strings.Contains(got, want) {
					t.Errorf("String() missing expected content %q in:\n%s", want, got)
				}
			}
		})
	}
}

func TestBlock_SetArrow(t *testing.T) {
	tests := []struct {
		name      string
		text      string
		width     int
		direction []BlockArrowDirection
		contains  []string
	}{
		{
			name:      "Right arrow",
			text:      "Test",
			direction: []BlockArrowDirection{BlockArrowDirectionRight},
			contains: []string{
				"0<[\"Test\"]>(right)",
			},
		},
		{
			name:      "Left arrow",
			text:      "Test",
			direction: []BlockArrowDirection{BlockArrowDirectionLeft},
			contains: []string{
				"0<[\"Test\"]>(left)",
			},
		},
		{
			name:      "Up arrow",
			text:      "Test",
			direction: []BlockArrowDirection{BlockArrowDirectionUp},
			contains: []string{
				"0<[\"Test\"]>(up)",
			},
		},
		{
			name:      "Down arrow",
			text:      "Test",
			direction: []BlockArrowDirection{BlockArrowDirectionDown},
			contains: []string{
				"0<[\"Test\"]>(down)",
			},
		},
		{
			name:      "X arrow",
			text:      "Test",
			direction: []BlockArrowDirection{BlockArrowDirectionX},
			contains: []string{
				"0<[\"Test\"]>(x)",
			},
		},
		{
			name:      "Y arrow",
			text:      "Test",
			direction: []BlockArrowDirection{BlockArrowDirectionY},
			contains: []string{
				"0<[\"Test\"]>(y)",
			},
		},
		{
			name:      "Multiple directions",
			text:      "Test",
			direction: []BlockArrowDirection{BlockArrowDirectionX, BlockArrowDirectionDown},
			contains: []string{
				"0<[\"Test\"]>(x, down)",
			},
		},
		{
			name:      "Arrow with width",
			text:      "Test",
			width:     2,
			direction: []BlockArrowDirection{BlockArrowDirectionRight},
			contains: []string{
				"0<[\"Test\"]>(right):2",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idGenerator = idGenerator.Reset()

			block := NewBlock("0", tt.text)
			if tt.width > 0 {
				block.SetWidth(tt.width)
			}
			block.SetArrow(tt.direction...)
			got := block.String()
			for _, want := range tt.contains {
				if !strings.Contains(got, want) {
					t.Errorf("String() missing expected content %q in:\n%s", want, got)
				}
			}
		})
	}
}

func TestBlock_SetColumns(t *testing.T) {
	tests := []struct {
		name     string
		block    *Block
		setup    func(*Block)
		contains []string
	}{
		{
			name:  "Block with columns",
			block: NewBlock("0", "Parent"),
			setup: func(b *Block) {
				b.diagram = &Diagram{}
				b.SetColumns(2)
				b.AddBlock("Child 1")
				b.AddBlock("Child 2")
			},
			contains: []string{
				"block:0:1",
				"columns 2",
				"0[\"Child 1\"]",
				"1[\"Child 2\"]",
				"end",
			},
		},
		{
			name:  "Block with columns and width",
			block: NewBlock("0", "Parent"),
			setup: func(b *Block) {
				b.diagram = &Diagram{}
				b.SetColumns(2)
				b.SetWidth(3)
				b.AddBlock("Child 1")
				b.AddBlock("Child 2")
			},
			contains: []string{
				"block:0:3",
				"columns 2",
				"0[\"Child 1\"]",
				"1[\"Child 2\"]",
				"end",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idGenerator = idGenerator.Reset()

			if tt.setup != nil {
				tt.setup(tt.block)
			}

			got := tt.block.String()
			for _, want := range tt.contains {
				if !strings.Contains(got, want) {
					t.Errorf("String() missing expected content %q in:\n%s", want, got)
				}
			}
		})
	}
}

func TestBlock_AddColumn(t *testing.T) {
	tests := []struct {
		name     string
		block    *Block
		setup    func(*Block)
		contains []string
	}{
		{
			name:  "Add column to empty block",
			block: NewBlock("0", "Parent"),
			setup: func(b *Block) {
				b.diagram = &Diagram{}
				b.AddColumn()
				b.AddBlock("Child 1")
				b.AddBlock("Child 2")
			},
			contains: []string{
				"block:0:1",
				"columns 1",
				"0[\"Child 1\"]",
				"1[\"Child 2\"]",
				"end",
			},
		},
		{
			name:  "Add column to block with existing columns",
			block: NewBlock("0", "Parent"),
			setup: func(b *Block) {
				b.diagram = &Diagram{}
				b.SetColumns(2)
				b.AddColumn()
				b.AddBlock("Child 1")
				b.AddBlock("Child 2")
			},
			contains: []string{
				"block:0:1",
				"columns 3",
				"0[\"Child 1\"]",
				"1[\"Child 2\"]",
				"end",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idGenerator = idGenerator.Reset()

			if tt.setup != nil {
				tt.setup(tt.block)
			}

			got := tt.block.String()
			for _, want := range tt.contains {
				if !strings.Contains(got, want) {
					t.Errorf("String() missing expected content %q in:\n%s", want, got)
				}
			}
		})
	}
}

func TestBlock_RemoveColumn(t *testing.T) {
	tests := []struct {
		name     string
		block    *Block
		setup    func(*Block)
		contains []string
	}{
		{
			name:  "Remove column from block with multiple columns",
			block: NewBlock("0", "Parent"),
			setup: func(b *Block) {
				b.diagram = &Diagram{}
				b.SetColumns(3)
				b.RemoveColumn()
				b.AddBlock("Child 1")
				b.AddBlock("Child 2")
			},
			contains: []string{
				"block:0:1",
				"columns 2",
				"0[\"Child 1\"]",
				"1[\"Child 2\"]",
				"end",
			},
		},
		{
			name:  "Remove column from block with one column",
			block: NewBlock("0", "Parent"),
			setup: func(b *Block) {
				b.diagram = &Diagram{}
				b.SetColumns(1)
				b.RemoveColumn()
				b.AddBlock("Child 1")
				b.AddBlock("Child 2")
			},
			contains: []string{
				"block:0:1",
				"0[\"Child 1\"]",
				"1[\"Child 2\"]",
				"end",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idGenerator = idGenerator.Reset()

			if tt.setup != nil {
				tt.setup(tt.block)
			}

			got := tt.block.String()
			for _, want := range tt.contains {
				if !strings.Contains(got, want) {
					t.Errorf("String() missing expected content %q in:\n%s", want, got)
				}
			}
		})
	}
}
