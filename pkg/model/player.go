package model

import (
	"github.com/google/uuid"
)

func NewPlayer(id uuid.UUID, username string, color Color) (*Player, error) {
	p := Player{
		ID:       id,
		Username: username,
		Color:    color,
		Buildings: &BuildingBank{
			Cities:      4,
			Settlements: 5,
			Roads:       15,
		},
		Resources: &ResourceCardDeck{
			Brick: 0,
			Ore:   0,
			Sheep: 0,
			Wheat: 0,
			Wood:  0,
		},
		NewDevCards: &DevCardDeck{
			Monopoly:     0,
			Monument:     0,
			RoadBuilding: 0,
			Soldier:      0,
			YearOfPlenty: 0,
		},
		UnplayedDevCards: &DevCardDeck{
			Monopoly:     0,
			Monument:     0,
			RoadBuilding: 0,
			Soldier:      0,
			YearOfPlenty: 0,
		},
		PlayedDevCards: &DevCardDeck{
			Monopoly:     0,
			Monument:     0,
			RoadBuilding: 0,
			Soldier:      0,
			YearOfPlenty: 0,
		},
	}

	return &p, nil
}

type Player struct {
	ID               uuid.UUID
	Username         string
	Color            Color
	Buildings        *BuildingBank
	Resources        *ResourceCardDeck
	NewDevCards      *DevCardDeck
	UnplayedDevCards *DevCardDeck
	PlayedDevCards   *DevCardDeck
}

type BuildingBank struct {
	Cities      uint
	Settlements uint
	Roads       uint
}
