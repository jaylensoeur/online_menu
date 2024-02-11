package create

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

func (cm *CreateMenu) Create(createMenuRequest CreateMenuRequestDto, createMenuPresenter domain.Presenter[CreateMenuResponseDto]) {
	menu, _ := cm.menuRepository.Find(*domain.NewTitle(createMenuRequest.Title))
	if menu != nil {
		createMenuPresenter.Present(NewCreateErrorMenuResponse())
		return
	}

	if menu == nil {
		newMenu, _ := cm.menuRepository.Add(domain.NewMenu(domain.NewCafeId(domain.NewUuid()), domain.NewTitle(createMenuRequest.Title)))
		createMenuPresenter.Present(CreateMenuResponseDto{Uuid: newMenu.GetCafeId().GetValue(), Title: newMenu.GetTitle().GetValue()})
		return
	}

	createMenuPresenter.Present(NewCreateErrorMenuResponse())
}
