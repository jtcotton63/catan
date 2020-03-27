package model_test

import (
	"testing"

	"github.com/jtcotton63/catan/model"
)

func TestResourceCardDeckDefaultInitialization(t *testing.T) {
	r := &model.ResourceCardDeck{}
	cnt := r.Count()
	if cnt != 0 {
		t.Fatalf("Expected a count of 0 to indicate correct default initialization of the resource card deck, but got %d", cnt)
	}
}
