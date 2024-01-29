package usecase

type ListMenuRequest struct {
	MetaData
}

type MetaData struct {
	Page      int    `json:"page"`
	Sort      string `json:"sort"`
	Limit     int    `json:"limit"`
	Count     int    `json:"count"`
	PageTotal int    `json:"pageTotal"`
}

type MenuDto struct {
	Title string `json:"title,omitempty"`
	Uuid  string `json:"uuid,omitempty"`
}

type ListMenuResponseDto struct {
	MetaData MetaData  `json:"metaData"`
	Data     []MenuDto `json:"data"`
	Error    string    `json:"error,omitempty"`
}
