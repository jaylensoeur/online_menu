package create

type CreateMenuRequestDto struct {
	Title string `json:"title"`
}

type CreateMenuResponseDto struct {
	Title string `json:"title,omitempty"`
	Uuid  string `json:"uuid,omitempty"`
	Error string `json:"error,omitempty"`
}

func NewCreateErrorMenuResponse() CreateMenuResponseDto {
	return CreateMenuResponseDto{
		Error: "Can not create menu, already exist",
	}
}
