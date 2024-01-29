package domain

type Presenter[T any] interface {
	Present(response T)
}
