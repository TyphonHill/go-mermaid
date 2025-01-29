package sequence

import (
	"reflect"
	"testing"
)

func TestNewActor(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		actorName string
		actorType ActorType
		want      *Actor
	}{
		{
			name:      "Create participant actor",
			id:        "user",
			actorName: "User",
			actorType: ActorParticipant,
			want: &Actor{
				ID:   "user",
				Name: "User",
				Type: ActorParticipant,
			},
		},
		{
			name:      "Create stick figure actor",
			id:        "admin",
			actorName: "Administrator",
			actorType: ActorActor,
			want: &Actor{
				ID:   "admin",
				Name: "Administrator",
				Type: ActorActor,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewActor(tt.id, tt.actorName, tt.actorType)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewActor() = %v, want %v", got, tt.want)
			}
		})
	}
}
