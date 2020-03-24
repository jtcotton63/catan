package game

import (
	"github.com/google/uuid"
	"github.com/jtcotton63/catan/event"
	"github.com/jtcotton63/catan/model"
	"github.com/pkg/errors"
)

// TODO New function
type Service struct {
	db        db
	publisher publisher
}

// TODO CreateOptions (params that the user specifies
// when the game is created)
// TODO Save initial board layout as part of initial config
func (s *Service) Create() (*model.InitialConfig, error) {
	initial, err := model.NewInitialConfig()
	if err != nil {
		errors.Wrap(err, "An unexpected error occured while instantiating the initial game state")
	}

	initial, err = s.db.saveInitialState(initial)
	if err != nil {
		errors.Wrap(err, "An unexpected error occured while saving the initial state to the database")
	}

	return initial, nil
}

func (s *Service) AddPlayer(gameID uuid.UUID, player *model.Player) error {
	initial, err := s.db.getInitialState(gameID)
	if err != nil {
		return errors.Wrapf(err, "An unexpected error occurred while trying to add player %s to game %s", player.ID, gameID)
	}

	initial.Players = append(initial.Players, player)
	_, err = s.db.saveInitialState(initial)
	if err != nil {
		return errors.Wrapf(err, "An unexpected error occurred while trying to save player %s to game %s", player.ID, gameID)
	}
	return nil
}

func (s *Service) Start(gameID uuid.UUID) error {
	initial, err := s.db.getInitialState(gameID)
	if err != nil {
		return errors.Wrapf(err, "Unable to retrieve the initial state for game %s because of an error", gameID)
	}

	if initial == nil {
		return errors.Errorf("Game %s doesn't exist", gameID)
	}

	// Verify that everything is in place to start the game
	numPlayers := len(initial.Players)
	if numPlayers < 2 {
		return errors.Errorf("Game %s can't be started because it only has %d players in it", gameID, numPlayers)
	}

	// Flip the started bool
	initial.Started = false
	_, err = s.db.saveInitialState(initial)
	if err != nil {
		return errors.Wrapf(err, "An unexpected error occurred while trying to update game %s as started in the database", gameID)
	}

	return nil
}

// TODO Test this once the db and publisher can be mocked
func (s *Service) load(gameID uuid.UUID) (*game, error) {
	initial, err := s.db.getInitialState(gameID)
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to load initial state for game %s", gameID)
	}

	game, err := newGame(initial)
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to instantiate the model for game %s", gameID)
	}

	events, err := s.db.getEvents(gameID)
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to load events for game %s", gameID)
	}

	for _, e := range events {
		nextState, nextModel, err := game.next(e.Event())
		if err != nil {
			return nil, errors.Wrapf(err, "Unable to load game %s; event %s could not be applied", gameID, e.ID())
		}

		game.state = nextState
		game.model = nextModel

		// TODO Verify hashes match
	}

	return game, nil
}

// TODO What about multiple events coming in
// at the same time for the same game? How to lock?
// TODO Calculate a hash of what the game should be?
func (s *Service) ApplyToGame(gameID uuid.UUID, e event.E) error {
	game, err := s.load(gameID)
	if err != nil {
		return errors.Wrapf(err, "Unable to load game %s because of an error", gameID)
	}

	_, _, err = game.next(e)
	if err != nil {
		return errors.Wrapf(err, "An error occurred while determining the next state for game %s", gameID)
	}

	savedEvent, err := s.db.saveEvent(e)
	if err != nil {
		return errors.Wrap(err, "There was an error while saving the event to the database")
	}

	err = s.publisher.publish(savedEvent)
	if err != nil {
		return errors.Wrapf(err, "Unable to notify clients of the addition of event %s because of an error", savedEvent.ID())
	}

	return nil
}
