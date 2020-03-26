package event

import "github.com/google/uuid"

func NewApplied(e E, id uuid.UUID) (*Applied, error) {
	a := Applied{
		event: e,
		id:    id,
	}
	return &a, nil
}

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
