package state_test

import (
	"testing"

	"github.com/jtcotton63/catan/statetype"

	"github.com/pkg/errors"

	"github.com/jtcotton63/catan/model"
	"github.com/jtcotton63/catan/state"

	"github.com/jtcotton63/catan/event"

	"github.com/google/uuid"
)

func TestRollingStateOnlyAcceptsRolledEvents(t *testing.T) {
	e, err := event.NewAcceptedTrade(
		uuid.New(),
		uuid.New(),
	)
	if err != nil {
		t.Error(errors.Wrap(err, "An unexpected error occurred while instantiating a new accepted trade event"))
	}

	s := state.NewRolling()
	_, _, err = s.Next(nil, e)
	if err == nil {
		t.Error("Expected an error but didn't get one")
	}

	msg := err.Error()
	if msg != "Expected an event of type RolledNumber but got an event of type AcceptedTrade" {
		t.Error(errors.Wrap(err, "Expected an error but got the wrong one"))
	}
}

func TestRollingStateTransitionToNormal(t *testing.T) {
	e, err := event.NewRolledNumber(
		uuid.New(),
		uuid.New(),
		2,
	)
	if err != nil {
		t.Error(err)
	}

	p1 := model.Player{
		ID: uuid.New(),
	}

	p2 := model.Player{
		ID: uuid.New(),
	}

	players := make([]*model.Player, 2, 2)
	players = append(players, &p1, &p2)

	gameModel, err := model.New(players)

	s := state.NewRolling()
	nextState, nextModel, err := s.Next(gameModel, e)

	// Verify the correct state has been received
	_, ok := nextState.(*state.NormalGameplay)
	if !ok {
		t.Errorf("Expected state %s but got %s", statetype.NormalGameplay, nextState.Type())
	}

	// TODO Verify that the correct resources have been distributed
	if nextModel == nil {
		t.Error("Unexpected error, nextModel is null")
	}
}

func TestRollingStateTransitionToDiscarding(t *testing.T) {
	e, err := event.NewRolledNumber(
		uuid.New(),
		uuid.New(),
		7,
	)
	if err != nil {
		t.Error(err)
	}

	p1, err := model.NewPlayer(uuid.New(), "player1", model.Blue)
	if err != nil {
		t.Error(err)
	}
	p2, err := model.NewPlayer(uuid.New(), "player2", model.Red)
	if err != nil {
		t.Error(err)
	}
	players := make([]*model.Player, 2, 2)
	players = append(players, p1, p2)

	gameModel, err := model.New(players)
	if err != nil {
		t.Error(err)
	}

	s := state.NewRolling()
	nextState, nextModel, err := s.Next(gameModel, e)

	// Verify the correct state has been received
	_, ok := nextState.(*state.Discarding)
	if !ok {
		t.Errorf("Expected state %s but got %s", statetype.DiscardingResources, nextState.Type())
	}

	// Verify that the model hasn't changed
	if nextModel == nil {
		t.Error("Unexpected error, nextModel is null")
	}
}
