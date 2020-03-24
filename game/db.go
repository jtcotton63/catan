package game

import (
	"github.com/google/uuid"
	"github.com/jtcotton63/catan/event"
)

type db interface {
	getInitialState(gameID uuid.UUID) (*game, error)
	getEvents(gameID uuid.UUID) ([]*event.Applied, error)
	saveEvent(event.E) (*event.Applied, error)
}
