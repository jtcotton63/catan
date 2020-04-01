package event

import (
	"github.com/google/uuid"
	"github.com/jtcotton63/catan/pkg/eventtype"
)

type E interface {
	GameID() uuid.UUID
	PlayerID() uuid.UUID
	Type() eventtype.T
}
