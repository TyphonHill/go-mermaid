package block

import (
	"fmt"
	"strings"
)

// BlockArrowDirection specifies the direction of a block arrow
type BlockArrowDirection string

const (
	baseBlockArrowShape = `<["%s"]>(%s)`
)

// Available arrow directions for block arrows
const (
	BlockArrowDirectionRight BlockArrowDirection = "right"
	BlockArrowDirectionLeft  BlockArrowDirection = "left"
	BlockArrowDirectionUp    BlockArrowDirection = "up"
	BlockArrowDirectionDown  BlockArrowDirection = "down"
	BlockArrowDirectionX     BlockArrowDirection = "x"
	BlockArrowDirectionY     BlockArrowDirection = "y"
)

// BlockArrowShape formats a block arrow with given text and direction(s)
func BlockArrowShape(text string, directions ...BlockArrowDirection) string {
	strs := make([]string, len(directions))
	for i, d := range directions {
		strs[i] = string(d)
	}
	dirStr := strings.Join(strs, ", ")
	return fmt.Sprintf(baseBlockArrowShape, text, dirStr)
}
