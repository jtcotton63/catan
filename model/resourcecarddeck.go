package model

import (
	"github.com/pkg/errors"
)

type ResourceCardDeck struct {
	Brick uint
	Ore   uint
	Sheep uint
	Wheat uint
	Wood  uint
}

// AddResourceCardDecks adds the values in r2 to the values in r1.
func AddResourceCardDecks(r1 *ResourceCardDeck, r2 *ResourceCardDeck) (*ResourceCardDeck, error) {
	brick := r1.Brick + r2.Brick
	ore := r1.Ore + r2.Ore
	sheep := r1.Sheep + r2.Sheep
	wheat := r1.Wheat + r2.Wheat
	wood := r1.Wood + r2.Wood

	rtn := ResourceCardDeck{
		Brick: brick,
		Ore:   ore,
		Sheep: sheep,
		Wheat: wheat,
		Wood:  wood,
	}
	return &rtn, nil
}

// SubtractResourceCardDecks subtracts the values in r2 from the values in r1.
func SubtractResourceCardDecks(r1 *ResourceCardDeck, r2 *ResourceCardDeck) (*ResourceCardDeck, error) {
	if r2.Brick > r1.Brick || r2.Ore > r1.Ore || r2.Sheep > r1.Sheep || r2.Wheat > r1.Wheat || r2.Wood > r1.Wood {
		return nil, errors.New("One or resources to be subtracted exceeded the quantity at hand")
	}

	brick := r1.Brick - r2.Brick
	ore := r1.Ore - r2.Ore
	sheep := r1.Sheep - r2.Sheep
	wheat := r1.Wheat - r1.Wheat
	wood := r1.Wood - r1.Wood

	rtn := ResourceCardDeck{
		Brick: brick,
		Ore:   ore,
		Sheep: sheep,
		Wheat: wheat,
		Wood:  wood,
	}
	return &rtn, nil
}
