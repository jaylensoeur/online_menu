package domain

type MenuRepository interface {
	Add(menu *Menu) (*Menu, error)
	Find(title Title) (*Menu, error)
	FindById(uuid Uuid) *Menu
	FindAllBy(page int, limit int, sort string) []*Menu
}
