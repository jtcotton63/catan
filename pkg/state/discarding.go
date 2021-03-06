package state

import (
	"github.com/google/uuid"
	"github.com/jtcotton63/catan/pkg/event"
	"github.com/jtcotton63/catan/pkg/eventtype"
	"github.com/jtcotton63/catan/pkg/model"
	"github.com/jtcotton63/catan/pkg/statetype"
	"github.com/pkg/errors"
)

func NewDiscarding() *Discarding {
	d := Discarding{
		playersWhoHaveDiscarded: make(map[uuid.UUID]bool),
	}
	return &d
}

type Discarding struct {
	playersWhoHaveDiscarded map[uuid.UUID]bool
}

func (d *Discarding) Type() statetype.T {
	return statetype.DiscardingResources
}

func (d *Discarding) Next(gameModel *model.Game, vanilla event.E) (S, *model.Game, error) {
	e, ok := vanilla.(*event.ResourcesDiscarded)
	if !ok {
		return nil, nil, errors.Errorf("Expected an event of type %s but got an event of type %s", eventtype.ResourcesDiscarded, vanilla.Type())
	}

	playerID := e.PlayerID()
	player, err := gameModel.GetPlayer(playerID)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "Player %s doesn't appear to be associated with game %s", playerID, e.GameID())
	}

	if player.Resources.Count() > 7 {
		player.Resources, err = model.SubtractResourceCardDecks(player.Resources, e.Resources())
		if err != nil {
			return nil, nil, errors.Wrapf(err, "Unable to subtract resources from player %s because of an error", playerID)
		}
	}

	// Mark this player as having discarded
	d.playersWhoHaveDiscarded[playerID] = true

	// If all players have discarded, go to the next state
	if len(d.playersWhoHaveDiscarded) == len(gameModel.Players) {
		return NewMovingRobber(), gameModel, nil
	}

	// Not all players have discarded, we need to continue to wait
	return d, gameModel, nil
}
