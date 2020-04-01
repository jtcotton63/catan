package game

import (
	"github.com/google/uuid"
	"github.com/jtcotton63/catan/pkg/event"
	"github.com/jtcotton63/catan/pkg/model"
	"github.com/jtcotton63/catan/pkg/state"
)

func newGame(initial *model.InitialConfig) (*game, error) {
	model, err := model.New(initial)
	if err != nil {
		return nil, err
	}

	initialState := state.NewRolling()

	g := game{
		ID:    initial.ID,
		model: model,
		state: initialState,
	}

	return &g, nil
}

type game struct {
	ID    uuid.UUID
	model *model.Game
	state state.S
}

func (g *game) next(e event.E) (state.S, *model.Game, error) {
	return g.state.Next(g.model, e)
}
