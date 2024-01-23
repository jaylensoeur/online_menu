package domain

type MenuRepository interface {
	Add(menu Menu) error
}
