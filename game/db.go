package game

import (
	"github.com/google/uuid"
	"github.com/jtcotton63/catan/event"
	"github.com/jtcotton63/catan/model"
)

type db interface {
	// Initial state
	getInitialState(gameID uuid.UUID) (*model.InitialConfig, error)
	saveInitialState(initial *model.InitialConfig) (*model.InitialConfig, error)

	// Events
	getEvents(gameID uuid.UUID) ([]*event.Applied, error)
	saveEvent(event.E) (*event.Applied, error)
}
