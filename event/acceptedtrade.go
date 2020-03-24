package event

import (
	"github.com/google/uuid"
	"github.com/jtcotton63/catan/eventtype"
)

func NewAcceptedTrade(gameID uuid.UUID, playerID uuid.UUID) (*AcceptedTrade, error) {
	a := AcceptedTrade{
		gameID:   gameID,
		playerID: playerID,
	}
	return &a, nil
}

type AcceptedTrade struct {
	gameID   uuid.UUID
	playerID uuid.UUID
}

func (a *AcceptedTrade) GameID() uuid.UUID {
	return a.gameID
}

func (a *AcceptedTrade) PlayerID() uuid.UUID {
	return a.playerID
}

func (a *AcceptedTrade) Type() eventtype.T {
	return eventtype.AcceptedTrade
}
