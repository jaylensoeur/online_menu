package menu

import (
	"encoding/json"
	"menu/application/create_menu"
	"net/http"
)

type CreateMenuPresenter struct {
	write http.ResponseWriter
}

type CreateMenuErrorPresenter struct {
	write   http.ResponseWriter
	request *http.Request
}

func NewCreateMenuErrorPresenter(write http.ResponseWriter, request *http.Request) *CreateMenuErrorPresenter {
	return &CreateMenuErrorPresenter{
		write,
		request,
	}
}

func (cmep *CreateMenuErrorPresenter) Present(response create_menu.CreateMenuResponse) {
	type menuErrorResponseDto struct {
		Error string `json:"error"`
	}

	cmep.write.Header().Set("Content-Type", "application/json")
	cmep.write.WriteHeader(http.StatusBadRequest)

	_ = json.NewEncoder(cmep.write).Encode(
		&menuErrorResponseDto{
			Error: "malformed data request",
		})
}

func NewCreateMenuPresenter(write http.ResponseWriter) *CreateMenuPresenter {
	return &CreateMenuPresenter{
		write: write,
	}
}

func (cmp *CreateMenuPresenter) Present(response create_menu.CreateMenuResponse) {
	cmp.write.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(cmp.write).Encode(response)
}
