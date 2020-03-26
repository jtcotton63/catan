package state

import (
	"github.com/jtcotton63/catan/event"
	"github.com/jtcotton63/catan/eventtype"
	"github.com/jtcotton63/catan/model"
	"github.com/jtcotton63/catan/statetype"
	"github.com/pkg/errors"
)

func NewMovingRobber() *MovingRobber {
	m := MovingRobber{}
	return &m
}

type MovingRobber struct{}

func (m *MovingRobber) Type() statetype.T {
	return statetype.MovingRobber
}

func (m *MovingRobber) Next(gameModel *model.Game, vanilla event.E) (S, *model.Game, error) {
	e, ok := vanilla.(*event.PlacedRobber)
	if !ok {
		return nil, nil, errors.Errorf("Expected an event of type %s but got an event of type %s", eventtype.PlacedRobber, vanilla.Type())
	}

	playerID := e.PlayerID()
	if playerID != gameModel.GetActivePlayer().ID {
		return nil, nil, errors.Errorf("Player %s cannot move the robber when it is player %s 's turn", playerID, gameModel.GetActivePlayer().ID)
	}

	// TODO Place the robber on the new hex

	// TODO RobbingNeighboringCommunity
	return nil, nil, nil
}
