package sequencediagram

// ActorType represents the visual representation of an actor in a sequence diagram.
type ActorType string

// Predefined actor types for sequence diagrams.
const (
	ActorParticipant ActorType = "participant"
	ActorActor       ActorType = "actor"
)

// Actor represents an entity participating in a sequence diagram.
type Actor struct {
	ID   string
	Name string
	Type ActorType
}

// NewActor creates a new Actor with the specified properties.
func NewActor(id, name string, actorType ActorType) *Actor {
	return &Actor{
		ID:   id,
		Name: name,
		Type: actorType,
	}
}
