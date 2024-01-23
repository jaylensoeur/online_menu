package domain

import "github.com/google/uuid"

type Uuid struct {
	id string
}

func NewUuid() Uuid {
	return Uuid{
		id: uuid.New().String(),
	}
}

func (u *Uuid) GetId() string {
	return u.id
}
