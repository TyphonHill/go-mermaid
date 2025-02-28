package basediagram

import (
	"fmt"
	"strings"
)

const Indentation = "    "

const (
	baseDiagramSeparator = "---\n"
	baseDiagramTitle     = "title: %s\n"
)

type BaseDiagram[T DiagramProperties] struct {
	Title  string
	Config T
	MarkdownFencer
}

func NewBaseDiagram[T DiagramProperties](config T) BaseDiagram[T] {
	return BaseDiagram[T]{
		Title:          "",
		Config:         config,
		MarkdownFencer: NewMarkdownFencer(),
	}
}

func (d *BaseDiagram[T]) SetTitle(title string) *BaseDiagram[T] {
	d.Title = title
	return d
}

func (d *BaseDiagram[T]) String(content string) string {
	var sb strings.Builder

	sb.WriteString(baseDiagramSeparator)

	if d.Title != "" {
		sb.WriteString(fmt.Sprintf(baseDiagramTitle, d.Title))
	}

	sb.WriteString(d.Config.String())

	sb.WriteString(baseDiagramSeparator)

	sb.WriteString(content)
	return d.MarkdownFencer.WrapWithFence(sb.String())
}
