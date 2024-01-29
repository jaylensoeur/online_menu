package inmemory

import (
	"fmt"
	"menu/domain"
)

type InMemoryMenuRepository struct {
	data map[string]*domain.Menu
}

func NewInMemoryRepository() *InMemoryMenuRepository {
	return &InMemoryMenuRepository{
		data: map[string]*domain.Menu{},
	}
}

func (im *InMemoryMenuRepository) Find(title domain.Title) (*domain.Menu, error) {
	return im.data[title.GetValue()], nil
}

func (im *InMemoryMenuRepository) Add(menu *domain.Menu) (*domain.Menu, error) {
	checkMenu := im.data[menu.GetTitle().GetValue()]
	if checkMenu != nil {
		return nil, fmt.Errorf("menu with that title already exist")
	}
	im.data[menu.GetTitle().GetValue()] = menu
	return menu, nil
}

func (im *InMemoryMenuRepository) FindById(uuid domain.Uuid) *domain.Menu {
	for _, menu := range im.data {
		if menu.GetCafeId().GetValue() == uuid.GetValue() {
			return menu
		}
	}
	return nil
}

func (im *InMemoryMenuRepository) FindAllBy(page int, limit int, sort string) []*domain.Menu {
	var menus []*domain.Menu
	for _, menu := range im.data {
		menus = append(menus, menu)
	}
	return menus
}
