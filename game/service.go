package game

import (
	"github.com/google/uuid"
	"github.com/jtcotton63/catan/event"
	"github.com/pkg/errors"
)

// TODO New function
type Service struct {
	db        db
	publisher publisher
}

// TODO Test this once the db and publisher can be mocked
func (s *Service) load(id uuid.UUID) (*game, error) {
	game, err := s.db.getInitialState(id)
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to load initial state for game %s", id)
	}

	events, err := s.db.getEvents(id)
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to load events for game %s", id)
	}

	for _, e := range events {
		nextState, nextModel, err := game.next(e.Event())
		if err != nil {
			return nil, errors.Wrapf(err, "Unable to load game %s; event %s could not be applied", id, e.ID())
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
