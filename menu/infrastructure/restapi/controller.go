package restapi

import (
	"menu/domain/create"
	"menu/domain/list"
	"menu/domain/single"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"menu/domain"
)

type Controller struct {
	createMenu *create.CreateMenu
	getMenu    *single.GetMenu
	listMenu   *list.ListMenu
}

func NewMenuController(
	createMenu *create.CreateMenu,
	getMenu *single.GetMenu,
	listMenu *list.ListMenu,
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
		var menuRequestDto create.CreateMenuRequestDto
		err := c.BodyParser(&menuRequestDto)

		if err != nil {
			mc.createMenu.Create(menuRequestDto, NewPresenter[create.CreateMenuResponseDto](c))
			return nil
		}

		mc.createMenu.Create(menuRequestDto, NewPresenter[create.CreateMenuResponseDto](c))
		return nil
	}
}

func (mc *Controller) Get() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		uuid := c.Params("id")
		newUuid := domain.NewUuidWithUuid(uuid)
		mc.getMenu.Retrieve(single.GetMenuRequest{Uuid: newUuid.GetValue()}, NewPresenter[single.GetMenuResponseDto](c))
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

		mc.listMenu.ListAllMenu(list.ListMenuRequest{
			MetaData: list.MetaData{
				Page:  pageInt,
				Limit: limitInt,
				Sort:  sort,
			},
		}, NewPresenter[list.ListMenuResponseDto](c))
		return nil
	}
}
