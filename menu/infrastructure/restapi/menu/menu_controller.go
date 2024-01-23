package menu

import (
	"encoding/json"
	"menu/application/create_menu"
	"menu/infrastructure/restapi/request"
	"net/http"
)

type menuRequestDto struct {
	Title string `json:"title"`
}

type Controller struct {
	createMenu *create_menu.CreateMenu
}

func NewMenuController(createMenu *create_menu.CreateMenu) *Controller {
	return &Controller{
		createMenu: createMenu,
	}
}

func (mc *Controller) Add(methods []string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if request.MethodMatched(methods, r.Method) {
			var menuRequestDto menuRequestDto
			err := json.NewDecoder(r.Body).Decode(&menuRequestDto)
			if err != nil {
				mc.createMenu.Create(create_menu.CreateMenuRequest{}, NewCreateMenuErrorPresenter(w, r))
				return
			}

			mc.createMenu.Create(
				create_menu.CreateMenuRequest{
					Title: menuRequestDto.Title,
				},
				NewCreateMenuPresenter(w),
			)
		}
	}
}
