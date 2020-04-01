package message

import (
	"github.com/jtcotton63/catan/pkg/event"
)

type Publisher interface {
	Publish(e *event.Applied) error
}
