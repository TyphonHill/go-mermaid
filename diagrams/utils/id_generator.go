package utils

import "fmt"

// IDGenerator defines the interface for generating unique IDs
type IDGenerator interface {
	NextID() string
}

// DefaultIDGenerator provides a simple incremental ID generator
type DefaultIDGenerator struct {
	nextID int
}

// NewIDGenerator creates a new DefaultIDGenerator
func NewIDGenerator() *DefaultIDGenerator {
	return &DefaultIDGenerator{nextID: 0}
}

// NextID generates the next unique ID
func (g *DefaultIDGenerator) NextID() string {
	id := g.nextID
	g.nextID++
	return fmt.Sprintf("%d", id)
}
