package game

import (
	"github.com/jtcotton63/catan/event"
)

type publisher interface {
	publish(e *event.Applied) error
}
