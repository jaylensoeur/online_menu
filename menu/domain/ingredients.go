package domain

type Ingredient struct {
	title       Title
	description Description
}

type Ingredients struct {
	ingredients []Ingredient
}
