package event

import (
	"github.com/google/uuid"
	"github.com/jtcotton63/catan/eventtype"
	"github.com/pkg/errors"
)

func NewRolledNumber(gameID uuid.UUID, playerID uuid.UUID, roll uint) (*RolledNumber, error) {
	if roll < 2 || roll > 12 {
		return nil, errors.Errorf("Invalid roll value %d", roll)
	}

	r := RolledNumber{
		gameID:   gameID,
		playerID: playerID,
		roll:     roll,
	}

	return &r, nil
}

type RolledNumber struct {
	gameID   uuid.UUID
	playerID uuid.UUID
	roll     uint
}

func (r *RolledNumber) GameID() uuid.UUID {
	return r.gameID
}

func (r *RolledNumber) PlayerID() uuid.UUID {
	return r.playerID
}

func (r *RolledNumber) Type() eventtype.T {
	return eventtype.RolledNumber
}

func (r *RolledNumber) Roll() uint {
	return r.roll
}
