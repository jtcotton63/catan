package state

import (
	"github.com/jtcotton63/catan/event"
	"github.com/jtcotton63/catan/model"
	"github.com/jtcotton63/catan/statetype"
)

func NewNormalGameplay() *NormalGameplay {
	n := NormalGameplay{}
	return &n
}

// TODO Only play one development card
type NormalGameplay struct{}

func (n *NormalGameplay) Type() statetype.T {
	return statetype.NormalGameplay
}

// TODO Active player ended their turn
func (n *NormalGameplay) Next(model *model.Game, vanilla event.E) (S, *model.Game, error) {
	return nil, nil, nil
}
