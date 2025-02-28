package basediagram

import (
	"fmt"
	"strings"
)

// DiagramProperty represents a configurable diagram property
type DiagramProperty interface {
	Format() string
	Value() interface{}
}

// BaseProperty provides common functionality for diagram properties
type BaseProperty struct {
	Name string
	Val  interface{}
}

func (p *BaseProperty) Format() string {
	return fmt.Sprintf(Indentation+Indentation+"%s: %v\n", p.Name, p.Val)
}

// Add this method to BaseProperty
func (p *BaseProperty) Value() interface{} {
	return p.Val
}

// Common property types
type BoolProperty struct {
	BaseProperty
}

type IntProperty struct {
	BaseProperty
}

type FloatProperty struct {
	BaseProperty
}

type StringProperty struct {
	BaseProperty
}

type StringArrayProperty struct {
	BaseProperty
}

func (p *StringArrayProperty) Format() string {
	vals := p.Val.([]string)
	quotedVals := make([]string, len(vals))
	for i, v := range vals {
		quotedVals[i] = fmt.Sprintf("%q", v)
	}
	return fmt.Sprintf(Indentation+Indentation+"%s: [%s]\n", p.Name, strings.Join(quotedVals, ", "))
}

type DiagramProperties interface {
	String() string
}
