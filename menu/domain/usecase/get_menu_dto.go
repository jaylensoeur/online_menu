package usecase

import (
	"fmt"
)

type GetMenuRequest struct {
	Uuid string
}

type GetMenuResponseDto struct {
	Title string `json:"title,omitempty"`
	Uuid  string `json:"uuid,omitempty"`
	Error string `json:"error,omitempty"`
}

func NewGetMenuErrorResponse(uuid string) GetMenuResponseDto {
	msg := fmt.Sprintf("Can not find menu with uuid %s", uuid)
	return GetMenuResponseDto{
		Error: msg,
	}
}

func NewGetMenuErrorUuidResponse() GetMenuResponseDto {
	return GetMenuResponseDto{
		Error: "Invalid Uuid string requested",
	}
}
