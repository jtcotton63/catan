package event

import (
	"github.com/google/uuid"
	"github.com/jtcotton63/catan/pkg/eventtype"
)

func NewEndedTurn(gameID uuid.UUID, playerID uuid.UUID) (*EndedTurn, error) {
	e := EndedTurn{
		gameID:   gameID,
		playerID: playerID,
	}
	return &e, nil
}

type EndedTurn struct {
	gameID   uuid.UUID
	playerID uuid.UUID
}

func (e *EndedTurn) GameID() uuid.UUID {
	return e.gameID
}

func (e *EndedTurn) PlayerID() uuid.UUID {
	return e.playerID
}

func (e *EndedTurn) Type() eventtype.T {
	return eventtype.EndedTurn
}
