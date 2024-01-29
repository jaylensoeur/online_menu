package domain

type CafeId struct {
	uuid Uuid
}

func NewCafeId(uuid Uuid) *CafeId {
	return &CafeId{
		uuid: uuid,
	}
}

func (c *CafeId) GetValue() string {
	return c.uuid.GetValue()
}
