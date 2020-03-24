package state

import (
	"github.com/jtcotton63/catan/event"
	"github.com/jtcotton63/catan/model"
	"github.com/jtcotton63/catan/statetype"
)

type S interface {
	Next(model *model.Game, e event.E) (S, *model.Game, error)
	Type() statetype.T
}
