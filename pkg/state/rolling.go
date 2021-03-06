package state

import (
	"github.com/jtcotton63/catan/pkg/event"
	"github.com/jtcotton63/catan/pkg/eventtype"
	"github.com/jtcotton63/catan/pkg/model"
	"github.com/jtcotton63/catan/pkg/statetype"
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
	e, ok := vanilla.(*event.RolledNumber)
	if !ok {
		return nil, nil, errors.Errorf("Expected an event of type %s but got an event of type %s", eventtype.RolledNumber, vanilla.Type())
	}

	// Make sure the active player is the one doing the rolling
	if e.PlayerID() != model.GetActivePlayer().ID {
		return nil, nil, errors.Errorf("Player %s is not the active player", e.PlayerID())
	}

	roll := e.Roll()
	if roll < 2 || roll > 12 {
		return nil, nil, errors.Errorf("Event has an invalid roll value %d", roll)
	}

	// A 7 was rolled, activate the robber
	if roll == 7 {
		return NewDiscarding(), model, nil
	}

	// A normal number was rolled, continue normal gameplay
	// TODO Distribute resources
	return NewNormalGameplay(), model, nil
}
