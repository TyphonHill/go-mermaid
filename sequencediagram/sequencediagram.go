package sequencediagram

import (
	"fmt"
	"strings"
)

const (
	baseTitleString           string = "---\ntitle: %s\n---\n\n"
	baseSequenecDiagramString string = "sequenceDiagram\n"
)

type SequenceDiagram struct {
	Title        string
	participants []*Participant
	actors       []*Actor
}

func NewSequenceDiagram() (newSequenceDiagram *SequenceDiagram) {
	newSequenceDiagram = &SequenceDiagram{}

	return
}

func (sd *SequenceDiagram) String() string {
	var sb strings.Builder

	if len(sd.Title) > 0 {
		sb.WriteString(fmt.Sprintf(string(baseTitleString), string(sd.Title)))
	}

	sb.WriteString(baseSequenecDiagramString)

	for _, participant := range sd.participants {
		sb.WriteString(participant.String())
	}

	for _, actor := range sd.actors {
		sb.WriteString(actor.String())
	}

	return sb.String()
}

func (sd *SequenceDiagram) AddParticipant(name string) (newParticipant *Participant) {
	newParticipant = NewParticipant(name)

	sd.participants = append(sd.participants, newParticipant)

	return
}
