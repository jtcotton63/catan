package event

import (
	"github.com/google/uuid"
	"github.com/jtcotton63/catan/pkg/eventtype"
)

func NewPlacedRobber(gameID uuid.UUID, playerID uuid.UUID) (*PlacedRobber, error) {
	p := PlacedRobber{
		gameID:   gameID,
		playerID: playerID,
	}
	return &p, nil
}

type PlacedRobber struct {
	gameID   uuid.UUID
	playerID uuid.UUID
}

func (p *PlacedRobber) GameID() uuid.UUID {
	return p.gameID
}

func (p *PlacedRobber) PlayerID() uuid.UUID {
	return p.playerID
}

func (p *PlacedRobber) Type() eventtype.T {
	return eventtype.PlacedRobber
}
