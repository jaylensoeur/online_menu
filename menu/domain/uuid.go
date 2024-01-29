package domain

import (
	"github.com/google/uuid"
)

type Uuid struct {
	id string
}

func NewUuid() Uuid {
	return Uuid{
		id: uuid.New().String(),
	}
}

func NewUuidWithUuid(id string) Uuid {
	defer func() {
		recover()
	}()

	uuidParsed := uuid.MustParse(id)
	return Uuid{uuidParsed.String()}

}

func (u *Uuid) GetValue() string {
	return u.id
}
