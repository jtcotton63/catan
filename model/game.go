package model

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func New(initial *InitialConfig) (*Game, error) {
	resources := ResourceCardDeck{
		Brick: 19,
		Ore:   19,
		Sheep: 19,
		Wheat: 19,
		Wood:  19,
	}

	devCards := DevCardDeck{
		Monopoly:     2,
		Monument:     5,
		RoadBuilding: 2,
		Soldier:      14,
		YearOfPlenty: 2,
	}

	g := Game{
		Players:      initial.Players,
		Active:       initial.Players[0],
		ResourceBank: &resources,
		DevCardBank:  &devCards,
	}

	return &g, nil
}

type Game struct {
	Players []*Player
	Active  *Player
	// TODO Board
	ResourceBank *ResourceCardDeck
	DevCardBank  *DevCardDeck
}

func (g *Game) GetPlayer(id uuid.UUID) (*Player, error) {
	for _, v := range g.Players {
		if v.ID == id {
			return v, nil
		}
	}

	return nil, errors.Errorf("Unable to find player %s", id)
}
