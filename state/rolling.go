package state

import (
	"github.com/jtcotton63/catan/event"
	"github.com/jtcotton63/catan/eventtype"
	"github.com/jtcotton63/catan/model"
	"github.com/jtcotton63/catan/statetype"
	"github.com/pkg/errors"
)

func NewRolling() *Rolling {
	r := Rolling{}
	return &r
}

type Rolling struct{}

func (r *Rolling) Type() statetype.T {
	return statetype.Rolling
}

// TODO Evaluate if someone has won
func (r *Rolling) Next(model *model.Game, vanilla event.E) (S, *model.Game, error) {
	e, ok := vanilla.(*event.RolledNumberEvent)
	if !ok {
		return nil, nil, errors.Errorf("Expected an event of type %s but got an event of type %s", eventtype.RolledNumber, vanilla.Type())
	}

	roll := e.Roll()
	if roll < 2 || roll > 12 {
		return nil, nil, errors.Errorf("Event has an invalid roll value %d", roll)
	}

	// A 7 was rolled, activate the robber
	// TODO The robber
	if roll == 7 {
		return NewDiscarding(), model, nil
	}

	// A normal number was rolled, continue normal gameplay
	// TODO Distribute resources
	return NewNormalGameplay(), model, nil
}
