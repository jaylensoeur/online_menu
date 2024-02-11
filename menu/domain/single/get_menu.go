package single

import (
	"menu/domain"
)

type GetMenu struct {
	menuRepository domain.MenuRepository
}

func NewGetMenu(menuRepository domain.MenuRepository) *GetMenu {
	return &GetMenu{
		menuRepository,
	}
}

func (cm *GetMenu) Retrieve(getMenuRequest GetMenuRequest, presenter domain.Presenter[GetMenuResponseDto]) {
	requestedUuid := domain.NewUuidWithUuid(getMenuRequest.Uuid)
	menu := cm.menuRepository.FindById(requestedUuid)
	if menu != nil {
		presenter.Present(GetMenuResponseDto{Uuid: menu.GetCafeId().GetValue(), Title: menu.GetTitle().GetValue()})
		return
	}
	errorResponse := NewGetMenuErrorUuidResponse()
	if getMenuRequest.Uuid != "" {
		errorResponse = NewGetMenuErrorResponse(getMenuRequest.Uuid)
	}
	presenter.Present(errorResponse)
	return
}
