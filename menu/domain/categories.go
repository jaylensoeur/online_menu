package domain

type Category struct {
	title  Title
	sortId SortId
	items  Items
}

type Categories struct {
	categories []Category
}
