package event

import "github.com/google/uuid"

type Applied struct {
	event E
	id    uuid.UUID
}

func (a *Applied) Event() E {
	return a.event
}

func (a *Applied) ID() uuid.UUID {
	return a.id
}
