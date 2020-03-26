package game

import (
	"testing"

	"github.com/jtcotton63/catan/state"
	"github.com/jtcotton63/catan/statetype"

	"github.com/google/uuid"
	"github.com/jtcotton63/catan/event"
	"github.com/jtcotton63/catan/model"
	"github.com/pkg/errors"
)

// TODO These same tests thru integration tests,
// so this file doesn't have to be bundled with the binary

// An in-memory data store
// used for testing
type mockDataStore struct {
	initials   map[uuid.UUID]*model.InitialConfig
	eventLists map[uuid.UUID][]*event.Applied
}

func (ds *mockDataStore) GetInitialState(gameID uuid.UUID) (*model.InitialConfig, error) {
	return ds.initials[gameID], nil
}

func (ds *mockDataStore) SaveInitialState(initial *model.InitialConfig) (*model.InitialConfig, error) {
	// The initial config is being saved for the first time
	if initial.ID == uuid.Nil {
		initial.ID = uuid.New()
	}

	ds.initials[initial.ID] = initial
	return ds.GetInitialState(initial.ID)
}

func (ds *mockDataStore) GetEvents(gameID uuid.UUID) ([]*event.Applied, error) {
	return ds.eventLists[gameID], nil
}

func (ds *mockDataStore) SaveEvent(e event.E) (*event.Applied, error) {
	id := uuid.New()
	applied, err := event.NewApplied(e, id)
	if err != nil {
		return nil, errors.Wrapf(err, "Unexpected error while trying to instantiate a new applied event from a vanilla event of type %s", e.Type())
	}

	gameID := applied.Event().GameID()
	if ds.eventLists[gameID] == nil {
		ds.eventLists[gameID] = make([]*event.Applied, 0)
	}
	ds.eventLists[gameID] = append(ds.eventLists[gameID], applied)

	return applied, nil
}

type mockPublisher struct{}

func (p *mockPublisher) Publish(e *event.Applied) error {
	// Event intentionally ignored
	return nil
}

func newMockedService() (*Service, error) {
	mockDataStore := mockDataStore{
		initials:   make(map[uuid.UUID]*model.InitialConfig),
		eventLists: make(map[uuid.UUID][]*event.Applied),
	}

	mockPublisher := mockPublisher{}

	mockedService, err := newService(&mockPublisher, &mockDataStore)
	if err != nil {
		return nil, errors.Wrap(err, "An unexpected error occurred while trying to instantiate a new service")
	}
	return mockedService, nil
}

func TestGameplay(t *testing.T) {

	// players := make([]*model.Player, 3, 3)
	// players = append(players, p1, p2, p3)

	// // Set up the model
	// gameModel, err := model.New(players)
	// if err != nil {
	// 	t.Error(errors.Wrap(err, "Unexpected error during setup"))
	// }

	gameSvc, err := newMockedService()
	if err != nil {
		t.Error(errors.Wrap(err, "An unexpected error occurred while instantiating a mocked service instance"))
	}

	initial, err := gameSvc.Create()
	gameID := initial.ID
	if err != nil {
		t.Error(errors.Wrap(err, "Couldn't create game"))
	}

	// Set up the players
	p1, err := model.NewPlayer(uuid.New(), "player1", model.Blue)
	if err != nil {
		t.Error(errors.Wrap(err, "Unexpected error during setup"))
	}
	err = gameSvc.AddPlayer(gameID, p1)
	if err != nil {
		t.Error(errors.Wrap(err, "Couldn't add p1 to the game"))
	}

	p2, err := model.NewPlayer(uuid.New(), "player2", model.Red)
	if err != nil {
		t.Error(errors.Wrap(err, "Unexpected error during setup"))
	}
	err = gameSvc.AddPlayer(gameID, p2)
	if err != nil {
		t.Error(errors.Wrap(err, "Couldn't add p2 to the game"))
	}

	p3, err := model.NewPlayer(uuid.New(), "player3", model.Yellow)
	if err != nil {
		t.Error(errors.Wrap(err, "Unexpected error during setup"))
	}
	err = gameSvc.AddPlayer(gameID, p3)
	if err != nil {
		t.Error(errors.Wrap(err, "Couldn't add p3 to the game"))
	}

	// Start the game
	err = gameSvc.Start(gameID)
	if err != nil {
		t.Error(errors.Wrap(err, "Couldn't start the game"))
	}

	// p1's turn
	p1t1e1, err := event.NewRolledNumber(gameID, p1.ID, 6)
	if err != nil {
		t.Error(errors.Wrap(err, "Couldn't instantiate p1t1e1"))
	}

	err = gameSvc.ApplyToGame(gameID, p1t1e1)
	if err != nil {
		t.Error(errors.Wrap(err, "Couldn't apply p1t1e1"))
	}

	game, err := gameSvc.load(gameID)
	if err != nil {
		t.Error(errors.Wrap(err, "Couldn't get game after p1t1e1"))
	}

	normalGameplayState, ok := game.state.(*state.NormalGameplay)
	if !ok {
		t.Errorf("Expected p1t1e1 to produce a state of type %s but it produced a state of type %s", statetype.NormalGameplay, normalGameplayState.Type())
	}

	p1t1e2, err := event.NewEndedTurn(gameID, p1.ID)
	if err != nil {
		t.Error(errors.Wrap(err, "Couldn't instantiate p1t1e2"))
	}

	err = gameSvc.ApplyToGame(gameID, p1t1e2)
	if err != nil {
		t.Error(errors.Wrap(err, "Couldn't apply p1t1e2"))
	}

	game, err = gameSvc.load(gameID)
	if err != nil {
		t.Error(errors.Wrap(err, "Couldn't get game after p1t1e2"))
	}

	rollingState, ok := game.state.(*state.Rolling)
	if !ok {
		t.Errorf("Expected p1t1e2 to produce a state of type %s but it produced a state of type %s", statetype.Rolling, rollingState.Type())
	}

	if game.model.GetActivePlayer().ID != p2.ID {
		t.Errorf("Expected p1t1e2 to change the active player to %s but the active player is %s", p2.ID, game.model.GetActivePlayer().ID)
	}
}
