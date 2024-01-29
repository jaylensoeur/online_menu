package inmemory

import (
	"cmp"
	"fmt"
	"menu/domain"
	"slices"
	"sort"
)

type InMemoryMenuRepository struct {
	data map[string]*domain.Menu
}

func NewInMemoryRepository(data map[string]*domain.Menu) *InMemoryMenuRepository {
	if data != nil {
		return &InMemoryMenuRepository{
			data: data,
		}
	}
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

func (im *InMemoryMenuRepository) FindAllBy(page int, limit int, order string) ([]*domain.Menu, domain.MetaData) {
	var menus []*domain.Menu
	itemTracker := 0
	skip := page * limit
	count := len(im.data)
	keys := make([]string, 0, count)
	pageTotal := count / limit

	if page > pageTotal {
		return nil, domain.NewMetaData(page, order, pageTotal, count, limit)
	}

	for k := range im.data {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	keys = keys[skip:]

	for _, key := range keys {
		if itemTracker >= limit {
			break
		}
		menus = append(menus, im.data[key])
		itemTracker++
	}

	slices.SortFunc[[]*domain.Menu, *domain.Menu](menus, func(a, b *domain.Menu) int {
		if n := cmp.Compare(a.GetTitle().GetValue(), b.GetTitle().GetValue()); n != 0 {
			return n
		}

		return 0
	})

	return menus, domain.NewMetaData(page, order, pageTotal, count, limit)
}
