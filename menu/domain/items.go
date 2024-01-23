package domain

type Item struct {
	title       Title
	costPrice   Price
	salePrice   Price
	ingredients Ingredients
}

type Items struct {
	items []Item
}
