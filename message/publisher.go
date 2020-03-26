package message

import (
	"github.com/jtcotton63/catan/event"
)

type Publisher interface {
	Publish(e *event.Applied) error
}
