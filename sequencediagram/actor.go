package sequencediagram

import (
	"fmt"
	"strings"
)

const (
	baseActorString string = "\tactor %s\n"
)

type Actor struct {
	Name  string
	Alias string
}

func (a *Actor) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf(string(baseActorString), a.Name))

	return sb.String()
}
