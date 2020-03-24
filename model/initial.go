package model

import "github.com/google/uuid"

func NewInitialConfig() (*InitialConfig, error) {
	players := make([]*Player, 0)
	i := InitialConfig{
		Players: players,
		Started: false,
	}
	return &i, nil
}

// InitialConfig describes the initial configuration
// used to create the game.
type InitialConfig struct {
	ID      uuid.UUID
	Players []*Player
	Started bool
}
