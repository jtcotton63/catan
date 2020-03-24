package game

import (
	"github.com/google/uuid"
	"github.com/jtcotton63/catan/event"
	"github.com/jtcotton63/catan/model"
	"github.com/jtcotton63/catan/state"
)

func new(players []*model.Player) (*game, error) {
	id := uuid.New()

	model, err := model.New(players)
	if err != nil {
		return nil, err
	}

	initialState := state.NewRolling()

	g := game{
		ID:    id,
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
