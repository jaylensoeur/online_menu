package domain

type Menu struct {
	cafeId     *CafeId
	categories Categories
	title      *Title
}

func NewMenu(cafeId *CafeId, title *Title) *Menu {
	return &Menu{
		cafeId: cafeId,
		title:  title,
	}
}

func (m *Menu) GetCafeId() *CafeId {
	return m.cafeId
}

func (m *Menu) GetTitle() *Title {
	return m.title
}
