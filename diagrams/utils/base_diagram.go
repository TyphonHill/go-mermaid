package utils

// BaseDiagram provides common functionality for all diagram types
type BaseDiagram struct {
	Title string
	MarkdownFencer
}

// NewBaseDiagram creates a new BaseDiagram with default settings
func NewBaseDiagram() BaseDiagram {
	return BaseDiagram{}
}

// SetTitle sets the diagram title and returns the diagram for chaining
func (d *BaseDiagram) SetTitle(title string) *BaseDiagram {
	d.Title = title
	return d
}

// GetTitle returns the diagram title
func (d *BaseDiagram) GetTitle() string {
	return d.Title
}
