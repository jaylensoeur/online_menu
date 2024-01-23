package mongo

import "menu/domain"

type MenuMongoRepository struct {
}

func NewMenuMongoRepository() *MenuMongoRepository {
	return &MenuMongoRepository{}
}

func (mmr *MenuMongoRepository) Add(menu domain.Menu) error {
	return nil
}
