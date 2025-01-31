package block

import (
	"reflect"
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
		name    string
		block   *Block
		setup   func(*Block)
		wantStr string
	}{
		{
			name:    "Empty block",
			block:   NewBlock("0", ""),
			wantStr: "\t0\n",
		},
		{
			name:    "Block with text",
			block:   NewBlock("1", "Test"),
			wantStr: "\t1[\"Test\"]\n",
		},
		{
			name:  "Block with width",
			block: NewBlock("2", "Test"),
			setup: func(b *Block) {
				b.SetWidth(2)
			},
			wantStr: "\t2[\"Test\"]:2\n",
		},
		{
			name:  "Block with style",
			block: NewBlock("3", "Test"),
			setup: func(b *Block) {
				b.SetStyle("fill:#f9f9f9")
			},
			wantStr: "\t3[\"Test\"]\n\tstyle 3 fill:#f9f9f9\n",
		},
		{
			name:  "Parent block with children",
			block: NewBlock("4", "Parent"),
			setup: func(b *Block) {
				b.diagram = &Diagram{}
				b.AddBlock("Child 1")
				b.AddBlock("Child 2")
			},
			wantStr: "\tblock:4:1\n\t\t0[\"Child 1\"]\n\t\t1[\"Child 2\"]\n\tend\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idGenerator = idGenerator.Reset()

			if tt.setup != nil {
				tt.setup(tt.block)
			}

			got := tt.block.String()
			if got != tt.wantStr {
				t.Errorf("Block.String() = %v, want %v", got, tt.wantStr)
			}
		})
	}
}

func TestBlock_SetShape(t *testing.T) {
	tests := []struct {
		name    string
		shape   blockShape
		text    string
		wantStr string
	}{
		{
			name:    "Default shape",
			shape:   BlockShapeDefault,
			text:    "Test",
			wantStr: "\t0[\"Test\"]\n",
		},
		{
			name:    "Round edges shape",
			shape:   BlockShapeRoundEdges,
			text:    "Test",
			wantStr: "\t0(\"Test\")\n",
		},
		{
			name:    "Stadium shape",
			shape:   BlockShapeStadium,
			text:    "Test",
			wantStr: "\t0([\"Test\"])\n",
		},
		{
			name:    "Subroutine shape",
			shape:   BlockShapeSubroutine,
			text:    "Test",
			wantStr: "\t0[[\"Test\"]]\n",
		},
		{
			name:    "Cylindrical shape",
			shape:   BlockShapeCylindrical,
			text:    "Test",
			wantStr: "\t0[(\"Test\")]\n",
		},
		{
			name:    "Circle shape",
			shape:   BlockShapeCircle,
			text:    "Test",
			wantStr: "\t0((\"Test\"))\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idGenerator = idGenerator.Reset()

			block := NewBlock("0", tt.text)
			block.SetShape(tt.shape)
			got := block.String()
			if got != tt.wantStr {
				t.Errorf("Block.String() with shape %v = %v, want %v", tt.shape, got, tt.wantStr)
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
		wantStr   string
	}{
		{
			name:      "Right arrow",
			text:      "Test",
			direction: []BlockArrowDirection{BlockArrowDirectionRight},
			wantStr:   "\t0<[\"Test\"]>(right)\n",
		},
		{
			name:      "Left arrow",
			text:      "Test",
			direction: []BlockArrowDirection{BlockArrowDirectionLeft},
			wantStr:   "\t0<[\"Test\"]>(left)\n",
		},
		{
			name:      "Up arrow",
			text:      "Test",
			direction: []BlockArrowDirection{BlockArrowDirectionUp},
			wantStr:   "\t0<[\"Test\"]>(up)\n",
		},
		{
			name:      "Down arrow",
			text:      "Test",
			direction: []BlockArrowDirection{BlockArrowDirectionDown},
			wantStr:   "\t0<[\"Test\"]>(down)\n",
		},
		{
			name:      "X arrow",
			text:      "Test",
			direction: []BlockArrowDirection{BlockArrowDirectionX},
			wantStr:   "\t0<[\"Test\"]>(x)\n",
		},
		{
			name:      "Y arrow",
			text:      "Test",
			direction: []BlockArrowDirection{BlockArrowDirectionY},
			wantStr:   "\t0<[\"Test\"]>(y)\n",
		},
		{
			name:      "Multiple directions",
			text:      "Test",
			direction: []BlockArrowDirection{BlockArrowDirectionX, BlockArrowDirectionDown},
			wantStr:   "\t0<[\"Test\"]>(x, down)\n",
		},
		{
			name:      "Arrow with width",
			text:      "Test",
			width:     2,
			direction: []BlockArrowDirection{BlockArrowDirectionRight},
			wantStr:   "\t0<[\"Test\"]>(right):2\n",
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
			if got != tt.wantStr {
				t.Errorf("Block.String() with arrow = %v, want %v", got, tt.wantStr)
			}
		})
	}
}

func TestBlock_SetColumns(t *testing.T) {
	tests := []struct {
		name    string
		block   *Block
		setup   func(*Block)
		wantStr string
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
			wantStr: "\tblock:0:1\n\t\tcolumns 2\n\t\t0[\"Child 1\"]\n\t\t1[\"Child 2\"]\n\tend\n",
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
			wantStr: "\tblock:0:3\n\t\tcolumns 2\n\t\t0[\"Child 1\"]\n\t\t1[\"Child 2\"]\n\tend\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idGenerator = idGenerator.Reset()

			if tt.setup != nil {
				tt.setup(tt.block)
			}

			got := tt.block.String()
			if got != tt.wantStr {
				t.Errorf("Block.String() = %v, want %v", got, tt.wantStr)
			}
		})
	}
}

func TestBlock_AddColumn(t *testing.T) {
	tests := []struct {
		name    string
		block   *Block
		setup   func(*Block)
		wantStr string
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
			wantStr: "\tblock:0:1\n\t\tcolumns 1\n\t\t0[\"Child 1\"]\n\t\t1[\"Child 2\"]\n\tend\n",
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
			wantStr: "\tblock:0:1\n\t\tcolumns 3\n\t\t0[\"Child 1\"]\n\t\t1[\"Child 2\"]\n\tend\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idGenerator = idGenerator.Reset()

			if tt.setup != nil {
				tt.setup(tt.block)
			}

			got := tt.block.String()
			if got != tt.wantStr {
				t.Errorf("Block.String() = %v, want %v", got, tt.wantStr)
			}
		})
	}
}

func TestBlock_RemoveColumn(t *testing.T) {
	tests := []struct {
		name    string
		block   *Block
		setup   func(*Block)
		wantStr string
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
			wantStr: "\tblock:0:1\n\t\tcolumns 2\n\t\t0[\"Child 1\"]\n\t\t1[\"Child 2\"]\n\tend\n",
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
			wantStr: "\tblock:0:1\n\t\t0[\"Child 1\"]\n\t\t1[\"Child 2\"]\n\tend\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idGenerator = idGenerator.Reset()

			if tt.setup != nil {
				tt.setup(tt.block)
			}

			got := tt.block.String()
			if got != tt.wantStr {
				t.Errorf("Block.String() = %v, want %v", got, tt.wantStr)
			}
		})
	}
}
