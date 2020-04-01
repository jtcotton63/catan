package event

import (
	"github.com/google/uuid"
	"github.com/jtcotton63/catan/pkg/eventtype"
	"github.com/jtcotton63/catan/pkg/model"
)

func NewResourcesDiscarded(gameID uuid.UUID, playerID uuid.UUID, resources *model.ResourceCardDeck) (*ResourcesDiscarded, error) {
	r := ResourcesDiscarded{
		gameID:    gameID,
		playerID:  playerID,
		resources: resources,
	}
	return &r, nil
}

type ResourcesDiscarded struct {
	gameID    uuid.UUID
	playerID  uuid.UUID
	resources *model.ResourceCardDeck
}

func (r *ResourcesDiscarded) GameID() uuid.UUID {
	return r.gameID
}

func (r *ResourcesDiscarded) PlayerID() uuid.UUID {
	return r.playerID
}

func (r *ResourcesDiscarded) Type() eventtype.T {
	return eventtype.ResourcesDiscarded
}

func (r *ResourcesDiscarded) Resources() *model.ResourceCardDeck {
	return r.resources
}
