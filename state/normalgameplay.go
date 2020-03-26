package state

import (
	"github.com/jtcotton63/catan/event"
	"github.com/jtcotton63/catan/eventtype"
	"github.com/jtcotton63/catan/model"
	"github.com/jtcotton63/catan/statetype"
	"github.com/pkg/errors"
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

// TODO All the other transitions
func (n *NormalGameplay) Next(game *model.Game, vanilla event.E) (S, *model.Game, error) {
	if vanilla.Type() == eventtype.EndedTurn {
		e, ok := vanilla.(*event.EndedTurn)
		if !ok {
			return nil, nil, errors.Errorf("Expected an event of type %s but got an event of type %s", eventtype.EndedTurn, vanilla.Type())
		}

		if e.PlayerID() != game.GetActivePlayer().ID {
			return nil, nil, errors.New("Only the active player can end their turn")
		}

		game.IncrementActivePlayer()
		return NewRolling(), game, nil
	}

	return nil, nil, nil
}
