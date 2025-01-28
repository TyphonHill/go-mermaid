package sequence_diagram

type ActorType string

const (
	ActorParticipant ActorType = "participant"
	ActorActor       ActorType = "actor"
)

type Actor struct {
	ID   string
	Name string
	Type ActorType
}

// Creates a new Actor with the specified properties
func NewActor(id, name string, actorType ActorType) *Actor {
	return &Actor{
		ID:   id,
		Name: name,
		Type: actorType,
	}
}
