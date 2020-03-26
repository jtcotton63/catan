package data

import (
	"github.com/google/uuid"
	"github.com/jtcotton63/catan/event"
	"github.com/jtcotton63/catan/model"
)

type Store interface {
	// Initial state
	GetInitialState(gameID uuid.UUID) (*model.InitialConfig, error)
	SaveInitialState(initial *model.InitialConfig) (*model.InitialConfig, error)

	// Events
	GetEvents(gameID uuid.UUID) ([]*event.Applied, error)
	SaveEvent(e event.E) (*event.Applied, error)
}
