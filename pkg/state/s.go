package state

import (
	"github.com/jtcotton63/catan/pkg/event"
	"github.com/jtcotton63/catan/pkg/model"
	"github.com/jtcotton63/catan/pkg/statetype"
)

type S interface {
	Next(model *model.Game, e event.E) (S, *model.Game, error)
	Type() statetype.T
}
