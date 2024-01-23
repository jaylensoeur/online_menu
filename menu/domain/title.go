package domain

type Title struct {
	value string
}

func NewTitle(title string) *Title {
	return &Title{
		value: title,
	}
}

func (t *Title) GetValue() string {
	return t.value
}
