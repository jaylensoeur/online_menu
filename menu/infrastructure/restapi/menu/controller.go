package menu

import (
	"strconv"

	"menu/domain"
	"menu/domain/usecase"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	createMenu *usecase.CreateMenu
	getMenu    *usecase.GetMenu
	listMenu   *usecase.ListMenu
}

func NewMenuController(
	createMenu *usecase.CreateMenu,
	getMenu *usecase.GetMenu,
	listMenu *usecase.ListMenu,
) *Controller {
	return &Controller{
		createMenu: createMenu,
		getMenu:    getMenu,
		listMenu:   listMenu,
	}
}

func (mc *Controller) Add() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		var menuRequestDto usecase.CreateMenuRequestDto
		err := c.BodyParser(&menuRequestDto)

		if err != nil {
			mc.createMenu.Create(menuRequestDto, NewPresenter[usecase.CreateMenuResponseDto](c))
			return nil
		}

		mc.createMenu.Create(menuRequestDto, NewPresenter[usecase.CreateMenuResponseDto](c))
		return nil
	}
}

func (mc *Controller) Get() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		uuid := c.Params("id")
		newUuid := domain.NewUuidWithUuid(uuid)
		mc.getMenu.Retrieve(usecase.GetMenuRequest{Uuid: newUuid.GetValue()}, NewPresenter[usecase.GetMenuResponseDto](c))
		return nil
	}
}

func (mc *Controller) List() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		page := c.Query("page", "1")
		limit := c.Query("limit", "2")
		sort := c.Query("sort", "ASC")

		pageInt, err := strconv.Atoi(page)
		if err != nil {
			return err
		}

		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			return err
		}

		mc.listMenu.ListAllMenu(usecase.ListMenuRequest{
			MetaData: usecase.MetaData{
				Page:  pageInt,
				Limit: limitInt,
				Sort:  sort,
			},
		}, NewPresenter[usecase.ListMenuResponseDto](c))
		return nil
	}
}
