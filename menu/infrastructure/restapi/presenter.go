package restapi

import (
	"github.com/gofiber/fiber/v2"
)

type Presenter[T any] struct {
	c *fiber.Ctx
}

func NewPresenter[T any](c *fiber.Ctx) *Presenter[T] {
	return &Presenter[T]{c}
}

func (cmp *Presenter[T]) Present(response T) {
	err := cmp.c.JSON(response)
	if err != nil {
		return
	}
}
