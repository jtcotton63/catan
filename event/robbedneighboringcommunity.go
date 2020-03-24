package event

import (
	"github.com/google/uuid"
	"github.com/jtcotton63/catan/eventtype"
	"github.com/jtcotton63/catan/model"
)

func NewRobbedNeighboringCommunity(gameID uuid.UUID, playerID uuid.UUID, robbedPlayerID uuid.UUID, resources *model.ResourceCardDeck) (*RobbedNeighboringCommunity, error) {
	r := RobbedNeighboringCommunity{
		gameID:         gameID,
		playerID:       playerID,
		robbedPlayerID: robbedPlayerID,
		resources:      resources,
	}
	return &r, nil
}

type RobbedNeighboringCommunity struct {
	gameID         uuid.UUID
	playerID       uuid.UUID
	robbedPlayerID uuid.UUID
	resources      *model.ResourceCardDeck
}

func (r *RobbedNeighboringCommunity) GameID() uuid.UUID {
	return r.gameID
}

func (r *RobbedNeighboringCommunity) PlayerID() uuid.UUID {
	return r.playerID
}

func (r *RobbedNeighboringCommunity) Type() eventtype.T {
	return eventtype.RobbedNeighboringCommunity
}

func (r *RobbedNeighboringCommunity) RobbedPlayerID() uuid.UUID {
	return r.robbedPlayerID
}

func (r *RobbedNeighboringCommunity) Resources() *model.ResourceCardDeck {
	return r.resources
}
