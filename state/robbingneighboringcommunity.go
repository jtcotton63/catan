package state

import (
	"github.com/google/uuid"
	"github.com/jtcotton63/catan/event"
	"github.com/jtcotton63/catan/eventtype"
	"github.com/jtcotton63/catan/model"
	"github.com/jtcotton63/catan/statetype"
	"github.com/pkg/errors"
)

func NewRobbingNeighboringCommunity() *RobbingNeighboringCommunity {
	r := RobbingNeighboringCommunity{}
	return &r
}

type RobbingNeighboringCommunity struct{}

func (r *RobbingNeighboringCommunity) Type() statetype.T {
	return statetype.RobbingNeighboringCommunity
}

func (r *RobbingNeighboringCommunity) Next(gameModel *model.Game, vanilla event.E) (S, *model.Game, error) {
	e, ok := vanilla.(*event.RobbedNeighboringCommunity)
	if !ok {
		return nil, nil, errors.Errorf("Expected an event of type %s but got an event of type %s", eventtype.RobbedNeighboringCommunity, vanilla.Type())
	}

	activePlayerID := e.PlayerID()
	if activePlayerID != gameModel.GetActivePlayer().ID {
		return nil, nil, errors.Errorf("Player %s cannot rob a neighboring community when it is player %s 's turn", activePlayerID, gameModel.GetActivePlayer().ID)
	}

	// The active player may not have robbed anyone
	robbedPlayerID := e.RobbedPlayerID()
	if robbedPlayerID == uuid.Nil {
		return NewNormalGameplay(), gameModel, nil
	}

	// The active player did rob someone
	// Subtract the cards from the robbed player
	robbedPlayer, err := gameModel.GetPlayer(robbedPlayerID)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "Player %s is not associated with game %s", robbedPlayerID, e.GameID())
	}

	robbedPlayer.Resources, err = model.SubtractResourceCardDecks(robbedPlayer.Resources, e.Resources())
	if err != nil {
		return nil, nil, errors.Wrapf(err, "Unable to subtract resources from player %s because of an error", robbedPlayerID)
	}

	// Add the cards to the robbing player's hand
	activePlayer, err := gameModel.GetPlayer(activePlayerID)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "Player %s is not associated with game %s", activePlayerID, e.GameID())
	}

	activePlayer.Resources, err = model.AddResourceCardDecks(activePlayer.Resources, e.Resources())
	if err != nil {
		return nil, nil, errors.Wrapf(err, "Unable to add resources to player %s 's hand because of an error", activePlayerID)
	}

	// Transition to normal gameplay state
	return NewNormalGameplay(), gameModel, nil
}
