package list

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

func menuEntityToMenuDtoMapper(menus []*domain.Menu) []MenuDto {
	//TODO: mapper domain entity to dto
	var menuDtos []MenuDto = nil
	for _, menu := range menus {
		menuDtos = append(menuDtos, MenuDto{
			Uuid:  menu.GetCafeId().GetValue(),
			Title: menu.GetTitle().GetValue(),
		})
	}
	return menuDtos
}

func (cm *ListMenu) ListAllMenu(listMenuRequest ListMenuRequest, presenter domain.Presenter[ListMenuResponseDto]) {
	menus, metaData := cm.menuRepository.FindAllBy(listMenuRequest.Page, listMenuRequest.Limit, listMenuRequest.Sort)
	presenter.Present(
		ListMenuResponseDto{
			MetaData: MetaData{
				PageTotal: metaData.PageTotal,
				Sort:      metaData.Sort,
				Limit:     metaData.Limit,
				Count:     metaData.Count,
				Page:      metaData.Page,
			},
			Data: menuEntityToMenuDtoMapper(menus),
		},
	)
}
