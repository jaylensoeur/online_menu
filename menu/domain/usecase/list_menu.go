package usecase

import (
	"menu/domain"
)

type ListMenu struct {
	menuRepository domain.MenuRepository
}

func NewListMenu(menuRepository domain.MenuRepository) *ListMenu {
	return &ListMenu{
		menuRepository,
	}
}

func (cm *ListMenu) ListAllMenu(listMenuRequest ListMenuRequest, presenter domain.Presenter[ListMenuResponseDto]) {
	menus := cm.menuRepository.FindAllBy(listMenuRequest.Page, listMenuRequest.Limit, listMenuRequest.Sort)

	//mapper domain entity to dto
	var menuDtos []Menu
	for _, menu := range menus {
		menuDtos = append(menuDtos, Menu{
			Uuid:  menu.GetCafeId().GetValue(),
			Title: menu.GetTitle().GetValue(),
		})
	}

	presenter.Present(
		ListMenuResponseDto{
			MetaData: MetaData{
				listMenuRequest.Page,
				listMenuRequest.Sort,
				listMenuRequest.Limit,
			},
			Data: menuDtos,
		},
	)
}
