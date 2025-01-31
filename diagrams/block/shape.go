package block

import (
	"fmt"
	"strings"
)

// BlockArrowDirection represents the direction of a block arrow
type BlockArrowDirection string

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
	return fmt.Sprintf(`<["%s"]>(%s)`, text, dirStr)
}
