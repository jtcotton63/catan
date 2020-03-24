package event

import (
	"github.com/google/uuid"
	"github.com/jtcotton63/catan/eventtype"
	"github.com/pkg/errors"
)

func NewRolledNumber(gameID uuid.UUID, playerID uuid.UUID, roll uint) (*RolledNumberEvent, error) {
	if roll < 2 || roll > 12 {
		return nil, errors.Errorf("Invalid roll value %d", roll)
	}

	r := RolledNumberEvent{
		gameID:   gameID,
		playerID: playerID,
		roll:     roll,
	}

	return &r, nil
}

type RolledNumberEvent struct {
	gameID   uuid.UUID
	playerID uuid.UUID
	roll     uint
}

func (r *RolledNumberEvent) GameID() uuid.UUID {
	return r.gameID
}

func (r *RolledNumberEvent) PlayerID() uuid.UUID {
	return r.playerID
}

func (r *RolledNumberEvent) Type() eventtype.T {
	return eventtype.RolledNumber
}

func (r *RolledNumberEvent) Roll() uint {
	return r.roll
}
