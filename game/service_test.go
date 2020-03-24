package game_test

// import (
// 	"testing"

// 	"github.com/pkg/errors"

// 	"github.com/google/uuid"
// 	"github.com/jtcotton63/catan/model"
// )

// func TestGameplay(t *testing.T) {

// 	// Set up the players
// 	p1, err := model.NewPlayer(uuid.New(), "player1", model.Blue)
// 	if err != nil {
// 		t.Error(errors.Wrap(err, "Unexpected error during setup"))
// 	}
// 	p2, err := model.NewPlayer(uuid.New(), "player2", model.Red)
// 	if err != nil {
// 		t.Error(errors.Wrap(err, "Unexpected error during setup"))
// 	}
// 	p3, err := model.NewPlayer(uuid.New(), "player3", model.Yellow)
// 	if err != nil {
// 		t.Error(errors.Wrap(err, "Unexpected error during setup"))
// 	}
// 	players := make([]*model.Player, 3, 3)
// 	players = append(players, p1, p2, p3)

// 	// Set up the model
// 	gameModel, err := model.New(players)
// 	if err != nil {
// 		t.Error(errors.Wrap(err, "Unexpected error during setup"))
// 	}

// }

// Add mocked services
// Call Create, get an initial state back (?)
// Add all players
// Call Start
// Add all events
// Add a CurrentState method
