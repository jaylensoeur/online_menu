package create_menu

import (
	"menu/domain"
)

type CreateMenu struct {
	menuRepository domain.MenuRepository
}

func NewCreateMenu(menuRepository domain.MenuRepository) *CreateMenu {
	return &CreateMenu{
		menuRepository,
	}
}

func (cm *CreateMenu) Create(createMenuRequest CreateMenuRequest, createMenuPresenter CreateMenuPresenter) {

	menu := domain.NewMenu(
		domain.NewCafeId(domain.NewUuid()),
		domain.NewTitle(createMenuRequest.Title),
	)

	createMenuPresenter.Present(
		CreateMenuResponse{
			Uuid:  menu.GetCafeId().GetValue(),
			Title: menu.GetTitle().GetValue(),
		},
	)
}
