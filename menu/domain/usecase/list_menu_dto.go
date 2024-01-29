package usecase

type ListMenuRequest struct {
	MetaData
}

type MetaData struct {
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
	Limit int    `json:"limit"`
}

type Menu struct {
	Title string `json:"title,omitempty"`
	Uuid  string `json:"uuid,omitempty"`
}

type ListMenuResponseDto struct {
	MetaData MetaData `json:"metaData"`
	Data     []Menu   `json:"data"`
	Error    string   `json:"error,omitempty"`
}
