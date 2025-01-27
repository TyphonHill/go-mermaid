package sequencediagram

import (
	"fmt"
	"strings"
)

const (
	baseParticipantString      string = "\tparticipant %s\n"
	baseParticipantAliasString string = "\tparticipant %s as %s\n"
)

type Participant struct {
	Name  string
	Alias string
}

func NewParticipant(name string) (newParticipant *Participant) {
	newParticipant = &Participant{
		Name: name,
	}

	return
}

func (p *Participant) String() string {
	var sb strings.Builder

	if len(p.Alias) > 0 {
		sb.WriteString(fmt.Sprintf(string(baseParticipantAliasString), p.Alias, p.Name))
	} else {
		sb.WriteString(fmt.Sprintf(string(baseParticipantString), p.Name))
	}

	return sb.String()
}
