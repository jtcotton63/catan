package game

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jtcotton63/catan/pkg/event"
	"github.com/jtcotton63/catan/pkg/model"
	"github.com/jtcotton63/catan/pkg/state"
	"github.com/jtcotton63/catan/pkg/statetype"
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

func applyEventAndVerifyGameState(gameSvc *Service, gameID uuid.UUID, vanilla event.E, expectedStateType statetype.T) (state.S, *model.Game, error) {
	err := gameSvc.ApplyToGame(gameID, vanilla)
	if err != nil {
		return nil, nil, err
	}

	game, err := gameSvc.load(gameID)
	if err != nil {
		return nil, nil, err
	}

	// Make sure the correct state was produced
	actualStateType := game.state.Type()
	if actualStateType != expectedStateType {
		return nil, nil, errors.Errorf("Expected event to produce a state of type %s but it produced a state of type %s", expectedStateType, actualStateType)
	}

	// TODO Make sure the correct object is behind the state type

	return game.state, game.model, nil
}

func TestGameplay(t *testing.T) {

	gameSvc, err := newMockedService()
	if err != nil {
		t.Fatal(errors.Wrap(err, "An unexpected error occurred while instantiating a mocked service instance"))
	}

	initial, err := gameSvc.Create()
	if err != nil {
		t.Fatal(errors.Wrap(err, "Couldn't create game"))
	}
	gameID := initial.ID

	// Set up the players
	p1, err := model.NewPlayer(uuid.New(), "player1", model.Blue)
	if err != nil {
		t.Fatal(errors.Wrap(err, "Unexpected error during setup"))
	}
	err = gameSvc.AddPlayer(gameID, p1)
	if err != nil {
		t.Fatal(errors.Wrap(err, "Couldn't add p1 to the game"))
	}

	p2, err := model.NewPlayer(uuid.New(), "player2", model.Red)
	if err != nil {
		t.Fatal(errors.Wrap(err, "Unexpected error during setup"))
	}
	err = gameSvc.AddPlayer(gameID, p2)
	if err != nil {
		t.Fatal(errors.Wrap(err, "Couldn't add p2 to the game"))
	}

	p3, err := model.NewPlayer(uuid.New(), "player3", model.Yellow)
	if err != nil {
		t.Fatal(errors.Wrap(err, "Unexpected error during setup"))
	}
	err = gameSvc.AddPlayer(gameID, p3)
	if err != nil {
		t.Fatal(errors.Wrap(err, "Couldn't add p3 to the game"))
	}

	// Start the game
	err = gameSvc.Start(gameID)
	if err != nil {
		t.Fatal(errors.Wrap(err, "Couldn't start the game"))
	}

	// p1's turn
	p1t1e1, err := event.NewRolledNumber(gameID, p1.ID, 6)
	if err != nil {
		t.Fatal(errors.Wrap(err, "Couldn't instantiate p1t1e1"))
	}

	_, _, err = applyEventAndVerifyGameState(gameSvc, gameID, p1t1e1, statetype.NormalGameplay)
	if err != nil {
		t.Fatal(errors.Wrap(err, "p1t1e1 failed"))
	}

	p1t1e2, err := event.NewEndedTurn(gameID, p1.ID)
	if err != nil {
		t.Fatal(errors.Wrap(err, "Couldn't instantiate p1t1e2"))
	}

	_, gameModel, err := applyEventAndVerifyGameState(gameSvc, gameID, p1t1e2, statetype.Rolling)
	if err != nil {
		t.Fatal(errors.Wrap(err, "p1t1e2 failed"))
	}

	if gameModel.GetActivePlayer().ID != p2.ID {
		t.Fatalf("Expected p1t1e2 to change the active player to %s but the active player is %s", p2.ID, gameModel.GetActivePlayer().ID)
	}

	// p2's turn
	p2t1e1, err := event.NewRolledNumber(gameID, p2.ID, 4)
	if err != nil {
		t.Fatal(errors.Wrap(err, "Couldn't instantiate p2t1e1"))
	}

	_, _, err = applyEventAndVerifyGameState(gameSvc, gameID, p2t1e1, statetype.NormalGameplay)
	if err != nil {
		t.Fatal(errors.Wrap(err, "p2t1e1 failed"))
	}

	p2t1e2, err := event.NewEndedTurn(gameID, p2.ID)
	if err != nil {
		t.Fatal(errors.Wrap(err, "Couldn't instantiate p2t1e2"))
	}

	_, gameModel, err = applyEventAndVerifyGameState(gameSvc, gameID, p2t1e2, statetype.Rolling)
	if err != nil {
		t.Fatal(errors.Wrap(err, "p2t1e2 failed"))
	}

	if gameModel.GetActivePlayer().ID != p3.ID {
		t.Fatalf("Expected p2t1e2 to change the active player to %s but the active player is %s", p3.ID, gameModel.GetActivePlayer().ID)
	}

	// p3's turn
	p3t1e1, err := event.NewRolledNumber(gameID, p3.ID, 8)
	if err != nil {
		t.Fatal(errors.Wrap(err, "Couldn't instantiate p3t1e1"))
	}

	_, _, err = applyEventAndVerifyGameState(gameSvc, gameID, p3t1e1, statetype.NormalGameplay)
	if err != nil {
		t.Fatal(errors.Wrap(err, "p3t1e1 failed"))
	}

	p3t1e2, err := event.NewEndedTurn(gameID, p3.ID)
	if err != nil {
		t.Fatal(errors.Wrap(err, "Couldn't instantiate p3t1e2"))
	}

	_, gameModel, err = applyEventAndVerifyGameState(gameSvc, gameID, p3t1e2, statetype.Rolling)
	if err != nil {
		t.Fatal(errors.Wrap(err, "p3t1e2 failed"))
	}

	if gameModel.GetActivePlayer().ID != p1.ID {
		t.Fatalf("Expected p3t1e2 to change the active player to %s but the active player is %s", p1.ID, gameModel.GetActivePlayer().ID)
	}

	// p1's turn again
	// Activating the robber
	p1t2e1, err := event.NewRolledNumber(gameID, p1.ID, 7)
	if err != nil {
		t.Fatal(errors.Wrap(err, "Couldn't instantiate p1t2e1"))
	}

	_, _, err = applyEventAndVerifyGameState(gameSvc, gameID, p1t2e1, statetype.DiscardingResources)
	if err != nil {
		t.Fatal(errors.Wrap(err, "p1t2e1 failed"))
	}

	p3Discarding, err := event.NewResourcesDiscarded(gameID, p3.ID, &model.ResourceCardDeck{
		Brick: 0,
		Ore:   0,
		Sheep: 0,
		Wheat: 0,
		Wood:  0,
	})
	if err != nil {
		t.Fatal(errors.Wrap(err, "Couldn't instantiate p3Discarding"))
	}

	_, _, err = applyEventAndVerifyGameState(gameSvc, gameID, p3Discarding, statetype.DiscardingResources)
	if err != nil {
		t.Fatal(errors.Wrap(err, "p3Dsicarding failed"))
	}

	p1Discarding, err := event.NewResourcesDiscarded(gameID, p1.ID, &model.ResourceCardDeck{
		Brick: 0,
		Ore:   0,
		Sheep: 0,
		Wheat: 0,
		Wood:  0,
	})
	if err != nil {
		t.Fatal(errors.Wrap(err, "Couldn't instantiate p1Discarding"))
	}

	_, _, err = applyEventAndVerifyGameState(gameSvc, gameID, p1Discarding, statetype.DiscardingResources)
	if err != nil {
		t.Fatal(errors.Wrap(err, "p1Discarding failed"))
	}

	p2Discarding, err := event.NewResourcesDiscarded(gameID, p2.ID, &model.ResourceCardDeck{
		Brick: 0,
		Ore:   0,
		Sheep: 0,
		Wheat: 0,
		Wood:  0,
	})
	if err != nil {
		t.Fatal(errors.Wrap(err, "p2Discarding failed"))
	}

	_, _, err = applyEventAndVerifyGameState(gameSvc, gameID, p2Discarding, statetype.MovingRobber)
	if err != nil {
		t.Fatal(errors.Wrap(err, "Couldn't apply p2Discarding"))
	}

	p1PlacingRobber, err := event.NewPlacedRobber(gameID, p1.ID)
	if err != nil {
		t.Fatal(errors.Wrap(err, "Couldn't instantiate p1PlacingRobber"))
	}

	_, _, err = applyEventAndVerifyGameState(gameSvc, gameID, p1PlacingRobber, statetype.RobbingNeighboringCommunity)
	if err != nil {
		t.Fatal(errors.Wrap(err, "Couldn't apply p1PlacingRobber"))
	}

	p1Robbing, err := event.NewRobbedNeighboringCommunity(gameID, p1.ID, p2.ID, &model.ResourceCardDeck{
		Brick: 0,
		Ore:   0,
		Sheep: 0,
		Wheat: 0,
		Wood:  0,
	})
	if err != nil {
		t.Fatal(errors.Wrap(err, "Couldn't instantiate p1Robbing"))
	}

	_, _, err = applyEventAndVerifyGameState(gameSvc, gameID, p1Robbing, statetype.NormalGameplay)
	if err != nil {
		t.Fatal(errors.Wrap(err, "Couldn't apply p1Robbing"))
	}
}
